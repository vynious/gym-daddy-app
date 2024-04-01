#!/bin/sh
# migrate.sh
# Run migrations
set -e

echo "running migrations........"

cd /app/db/migrations

go run . init && go run . up