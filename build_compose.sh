#!/bin/bash

# Build Docker images
echo "Building Docker images..."

echo " "

echo "Building Docker Queue CMS..."
docker build -t gd-join-queue-cms gd-joinqueue-cms/.

echo " "

echo "Building Docker Notification MS..."
docker build -t gd-notification-ms gd-notification-ms/.

echo " "

echo "Building Docker Queue MS..."
docker build -t gd-queue-ms gd-queue-ms/.

echo " "

echo "Building Docker Telemessenger MS..."
docker build -t gd-telemessenger-ms gd-telemessenger-ms/.

echo " "

# Add more docker build commands for additional images if needed
#
#echo "Building Docker Log MS..."
#docker build -t gd-log-ms gd-log-ms/.


echo " "

echo "Building Docker Kong Gateway..."
docker build -t gd-kong-gateway gd-kong-gateway/.


echo " "

# Run Docker Compose
echo "Running Docker Compose..."
docker-compose up -d
