#!/bin/bash

# Define the local backup path and Google Cloud Storage path
LOCAL_BACKUP_PATH="/docker-entrypoint-initdb.d/backup.archive"
REMOTE_BACKUP_PATH="butlerai:/backups/backup.archive"

# Download the backup from Google Cloud Storage
rclone copy $REMOTE_BACKUP_PATH $LOCAL_BACKUP_PATH
echo "Backup downloaded from Google Cloud Storage."

# Restore the backup to MongoDB
mongorestore --archive=$LOCAL_BACKUP_PATH --gzip --drop
echo "Data restore completed."
