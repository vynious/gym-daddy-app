#!/bin/bash

# Build Docker images
echo "Building Docker images..."

echo " "

echo "Building Docker Queue CMS..."
docker build -t gd-join-queue-cms gym-daddy-joinqueue-cms/.

echo " "

echo "Building Docker Notification MS..."
docker build -t gd-notification-ms gym-daddy-notification-ms/.

echo " "

echo "Building Docker Queue MS..."
docker build -t gd-queue-ms gym-daddy-queue-ms/.

echo " "

echo "Building Docker Telemessenger MS..."
docker build -t gd-telemessager-ms gym-daddy-telemessenger-ms/.

echo " "

echo "Building Docker User MS..."
docker build -t gd-users-ms gd-users-ms        

echo " "

echo "Building Docker Gym Avail MS..."
docker build -t gd-gym-avail-ms gd-gym-avail-ms 

echo " "
# Add more docker build commands for additional images if needed

echo "Building Docker Log MS..."
docker build -t gd-log-ms gym-daddy-log-ms/.


echo " "

echo "Building Docker Kong Gateway..."
docker build -t gd-kong-gateway gym-daddy-kong-gateway/.

echo " "

# Run Docker Compose
echo "Running Docker Compose..."
docker-compose up -d
