#!/bin/bash

echo "start building container ...."
docker-compose -p fitness_app down -v
docker-compose -p fitness_app up -d --build

echo "Deployment complete!"
