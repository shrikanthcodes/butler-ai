#!/bin/bash

# Import data if the backup exists
if [ -f /data/backup.archive ]; then
  mongorestore --archive=/data/backup.archive
  echo "Data import completed."
fi

# Trap the stop signal to export data on shutdown
trap "mongodump --archive=/data/backup.archive; echo 'Data export completed.'" SIGTERM

# Start MongoDB
exec mongod --bind_ip 0.0.0.0
