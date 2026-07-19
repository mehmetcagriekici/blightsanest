# Deploying BlightSanest to Kubernetes

These manifests target a **local minikube cluster** (docker driver) — there's
no registry push involved, images are built straight into the cluster's
container runtime. If you're deploying to a real cluster later, replace the
image-loading step with `docker push` to your registry and update the
`image:` fields + `imagePullPolicy` in each Deployment/Job accordingly.

## Architecture notes (read before you deploy)

- **`server`, `client`, `search` are interactive REPLs**, not HTTP/TCP
  services — they read commands from stdin (see `README.md` at the repo
  root). They're deployed as Deployments with `stdin: true, tty: true` so you
  can drive them with `kubectl attach` / `kubectl exec -it`, not as anything
  you'd put behind a Service.
- `client` starts at **0 replicas** — it's meant to be run "from separate
  terminals" per the app's own design, one pod per concurrent user session.
  Scale it up and exec into the specific pod you want.
- `semantic-api` must stay at **1 replica**. It keeps computed embeddings in
  a single in-process Python object with no shared backing store; a second
  pod would build its own independent cache and give inconsistent answers.
- Postgres and RabbitMQ are stateful (PVC-backed): Postgres via a
  StatefulSet, RabbitMQ via a Deployment + PVC (`strategy: Recreate` so the
  old pod releases the volume before the new one claims it).

## One-time setup

```bash
cp k8s/secret.env.example k8s/secret.env
# edit k8s/secret.env: set real passwords and your COIN_GECKO_KEY
```

`k8s/secret.env` is gitignored — never commit it. It feeds a Kustomize
`secretGenerator`, so `kubectl apply -k k8s/` regenerates the Secret from it
automatically.

## Build images into minikube

```bash
minikube start --driver=docker   # if not already running

eval $(minikube docker-env)      # point your shell's docker CLI at minikube's daemon
# (if your user isn't in the docker group, use: sg docker -c "eval $(minikube docker-env) && ...")

docker build -t blightsanest-go:local .
docker build -t blightsanest-semantic:local ./blightsanest_semantic_search
```

If you'd rather not switch `docker-env`, build normally on your host and load
the images in instead:

```bash
docker build -t blightsanest-go:local .
docker build -t blightsanest-semantic:local ./blightsanest_semantic_search
minikube image load blightsanest-go:local
minikube image load blightsanest-semantic:local
```

Both Deployments/Jobs reference these images with `imagePullPolicy: Never`,
so the cluster never tries to pull them from a registry.

## Deploy

```bash
kubectl apply -k k8s/
```

This creates the `blightsanest` namespace, config/secrets, Postgres,
RabbitMQ, the semantic-search API, runs the `blightsanest-migrate` Job
(applies the goose migrations in `sql/schema/` against Postgres), and starts
`server`/`search` (and `client` at 0 replicas).

Check status:

```bash
kubectl -n blightsanest get pods
kubectl -n blightsanest get job blightsanest-migrate   # should show Complete
```

## Using the app

The server and search REPLs are already running — attach to drive them:

```bash
kubectl -n blightsanest attach -it deploy/server
kubectl -n blightsanest attach -it deploy/search
```

(Ctrl-P Ctrl-Q detaches without killing the process; `quit` inside the REPL
exits it, which will make the pod restart since the container's main process
exited.)

For a client session, scale up and exec into a specific pod:

```bash
kubectl -n blightsanest scale deploy/client --replicas=1
kubectl -n blightsanest get pods -l app=client
kubectl -n blightsanest exec -it <client-pod-name> -- ./client
```

Scale `client` up further (`--replicas=N`) for N concurrent sessions, one
pod per user.

## Re-running migrations

The migrate Job runs once. If you add new files under `sql/schema/`, rebuild
the `blightsanest-go` image (the migrations are embedded into the `migrate`
binary via `go:embed`) and re-run the job:

```bash
kubectl -n blightsanest delete job blightsanest-migrate
kubectl apply -k k8s/
```

## Known limitations / things not covered here

- No Ingress/external access — nothing in this app is meant to be reached by
  a browser; `semantic-api` is only ever called cluster-internally by
  `search`. If you need to reach RabbitMQ's management UI (port 15672) from
  outside the cluster, use `kubectl -n blightsanest port-forward
  deploy/rabbitmq 15672:15672`.
- No NetworkPolicies, no TLS between services, no HorizontalPodAutoscaler —
  none of these are meaningful for a single-user local deployment; add them
  if you take this to a shared/production cluster.
- `k8s/secret.env` ships default/placeholder credentials in the `.example`
  file. Don't reuse `changeme` anywhere but a local throwaway cluster.
