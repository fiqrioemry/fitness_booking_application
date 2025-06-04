#!/bin/sh

host="$1"
shift
cmd="$@"

echo "Waiting for $host..."

while ! nc -z $(echo "$host" | cut -d: -f1) $(echo "$host" | cut -d: -f2); do
  sleep 2
done

echo "✅ $host is up. Starting app..."
exec $cmd
