#!/bin/bash

case "$1" in
    start)
	echo "Starting PostgreSQL via service..."
	sudo service postgresql start
	;;
    stop)
	echo "Stopping PostgreSQL via service..."
	sudo service postgresql stop
	;;
    *)
	echo "Usage: $0 {start|stop}"
esac
