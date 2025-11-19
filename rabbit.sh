#!/bin/bash

start_or_run() {
    docker inspect blightsanest_rabbitmq > /dev/null 2>&1

    if [ $? -eq 0 ]; then
       echo "Starting BlightSanest RabbitMQ container..."
       docker start blightsanest_rabbitmq
    else
	echo "BlightSanest RabbitMQ container not found creating a new one..."
        docker run -d --name blightsanest_rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management
    fi
}

case "$1" in
    start)
	start_or_run
	;;
    stop)
	echo "Stopping BlightSanest RabbitMQ container..."
	docker stop blightsanest_rabbitmq
	;;
    logs)
	echo "Fetching logs for BlightSanest RabbitMQ container..."
	docker logs -f blightsanest_rabbitmq 
	;;
    *)
	echo "Usage: $0 {start|stop|logs}"
esac
