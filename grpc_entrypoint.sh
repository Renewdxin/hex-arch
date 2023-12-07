#!/bin/sh

set -e
COMMAND=$@

echo 'Waiting for database to be ready...'
maxTries = 10
while ["$maxTries" -gt 0] && ! mysql -h "$MYSQL_HOST" -u "$MYSQL_USER" -p"$MYSQL_PASSWORD" -e "SHOW TABLES"; do
  maxTries=$(($maxTries-1))
  sleep 3
done
echo

if [ "$maxTries" -le 0 ]; then
  echo >&2 'Database is not ready after 30 seconds, aborting.'
  exit 1
fi

exec $COMMAND