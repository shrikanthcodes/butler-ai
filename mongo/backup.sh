#!/bin/bash

# Define the local backup path and Google Cloud Storage path
LOCAL_BACKUP_PATH="/docker-entrypoint-initdb.d/backup.archive"
REMOTE_BACKUP_PATH="butlerai:/backups/backup.archive"

# Create a MongoDB dump
mongodump --archive=$LOCAL_BACKUP_PATH --gzip
echo "Local backup completed."

# Upload the backup to Google Cloud Storage
rclone copy $LOCAL_BACKUP_PATH $REMOTE_BACKUP_PATH
echo "Backup uploaded to Google Cloud Storage."
