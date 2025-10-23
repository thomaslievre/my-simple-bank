#!/bin/sh

set -e

echo "Running migrations..."
/app/migrate

echo "start the app"
exec "$@"