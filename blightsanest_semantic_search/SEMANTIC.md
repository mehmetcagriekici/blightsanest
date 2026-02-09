# BlightSanest - Semantic Search Engine

Semantic Search Engine is built upon Sentence Transformers model "all-MiniLM-L6-v2" and implemented using Python.It interacts with the main application via a FastAPI server using two endpoints, one for creating the embeddings, and the other for the search.

```
@app.post("/embeddings")
@app.post("/search")
```

You need to trigger embeddings before committing any search, as instructed in the search CLI help.
