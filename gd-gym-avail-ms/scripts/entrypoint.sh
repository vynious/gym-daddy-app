#!/bin/bash

# Wait for the database to be ready
until timeout 1 bash -c "echo > /dev/tcp/user-db/5432"; do
  echo "Waiting for database..."
  sleep 5
done

chmod +x /app/scripts/migrate.sh

# Run database migrations
./scripts/migrate.sh

# Start the main application
./main
