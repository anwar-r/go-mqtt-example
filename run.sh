#!/bin/sh

echo "Starting Publisher..."
cd ./publisher && go run publisher.go &

echo "Starting Consumer..."
cd ./consumer && go run consumer.go
