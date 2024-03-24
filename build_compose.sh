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
echo "Building Docker Gym Avail MS..."
docker build -t gd-gym-avail-ms gd-gym-avail-ms/.

echo " "
echo "Building Docker User MS..."
docker build -t gd-users-ms gd-users-ms

echo " "
echo "Building Docker Kong Gateway..."
docker build -t gd-kong-gateway gd-kong-gateway/.


echo ""
echo "Building Docker Classes MS..."
docker build -t gd-classes-ms gd-classes-ms/docker/.

echo ""
echo "Building Docker Booking MS..."
docker build -t gd-booking-ms gd-booking-ms/.

echo ""
echo "Building Docker CreateBooking CMS..."
docker build -t gd-createbooking-cms gd-createbooking-cms/.

echo " "
echo "Running Docker Compose..."
docker-compose up -d
