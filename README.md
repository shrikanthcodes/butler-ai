Here is the updated README reflecting the changes made with MongoDB Dockerfile and `init.sh` being removed and replaced by `backup.sh`, `restore.sh`, and the usage of `rclone` for cloud backups:

---

# ButlerAI

## Overview
ButlerAI is a full-stack application utilizing Docker to streamline development and deployment. The project includes a Go-based backend, a React Native frontend powered by Expo, and MongoDB as the database. The setup is designed for ease of deployment and local development using Docker, with automated backups to Google Drive using `rclone`.

## Project Structure
```bash
ButlerAI/
├── backend/
│   ├── cmd/
│   ├── internal/
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── ...
├── frontend/
│   ├── src/
│   ├── Dockerfile
│   ├── package.json
│   ├── package-lock.json
│   └── ...
├── mongo/
│   ├── backup.sh
│   ├── restore.sh
│   └── ...
├── docker-compose.yml
└── README.md
```

## Requirements
- **Docker**: Ensure Docker is installed on your system. [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: Install Docker Compose. [Install Docker Compose](https://docs.docker.com/compose/install/)
- **Node.js and npm**: Required for running the frontend locally. [Install Node.js](https://nodejs.org/)
- **rclone**: Required for interacting with Google Drive for backup and restore operations. [Install rclone](https://rclone.org/install/)

## Useful Commands

### 1. Check active docker containers
```bash
sudo docker ps
``` 

### 2. Stop docker container
```bash
sudo docker stop <container-id/container-name>
``` 

## Installation Steps

### 1. Clone the Repository
```bash
git clone https://github.com/your-repo/butlerAI.git
cd butlerAI
```

### 2. Set Up Docker
Ensure Docker and Docker Compose are installed on your system. Follow the official guides if needed.

### 3. Configure `rclone`
Set up `rclone` to interact with Google Drive by following the steps outlined in the [rclone Google Drive guide](https://rclone.org/drive/). Ensure you create a remote named `butlerai`. Request access to project app directory. If interested, reach out to `shrikanthsubramanian@gmail.com`.

### 4. Build and Run the Entire Project

To build and run the frontend, backend, and MongoDB together:

```bash
sudo docker-compose up --build
```

This will start all services:
- **Backend** on port `8080`
- **Frontend (Expo)** on ports `19006`, `19001`, and `19002`
- **MongoDB** on port `27017`

### 5. Stopping the Services
To stop the services:

```bash
sudo docker-compose down
```

## MongoDB Data Management

### Backup and Restore Scripts

The MongoDB service data can be managed using the provided `backup.sh` and `restore.sh` scripts, which automate the backup and restore process using `rclone` to sync with Google Drive.

- **Backup Script (`backup.sh`)**: This script creates a local backup of the MongoDB database and uploads it to Google Drive.

```bash
#!/bin/bash

# Define the local backup path and Google Drive path
LOCAL_BACKUP_PATH="/docker-entrypoint-initdb.d/backup.archive"
REMOTE_BACKUP_PATH="butlerai:/backups/backup.archive"

# Create a MongoDB dump
mongodump --archive=$LOCAL_BACKUP_PATH --gzip
echo "Local backup completed."

# Upload the backup to Google Drive
rclone copy $LOCAL_BACKUP_PATH $REMOTE_BACKUP_PATH
echo "Backup uploaded to Google Drive."
```

- **Restore Script (`restore.sh`)**: This script downloads the latest backup from Google Drive and restores it to the MongoDB instance.

```bash
#!/bin/bash

# Define the local backup path and Google Drive path
LOCAL_BACKUP_PATH="/docker-entrypoint-initdb.d/backup.archive"
REMOTE_BACKUP_PATH="butlerai:/backups/backup.archive"

# Download the backup from Google Drive
rclone copy $REMOTE_BACKUP_PATH $LOCAL_BACKUP_PATH
echo "Backup downloaded from Google Drive."

# Restore the backup to MongoDB
mongorestore --archive=$LOCAL_BACKUP_PATH --gzip --drop
echo "Data restore completed."
```

### Running the Scripts

- **Backup MongoDB Data:**
  ```bash
  bash mongo/backup.sh
  ```
- **Restore MongoDB Data:**
  ```bash
  bash mongo/restore.sh
  ```

These scripts will ensure your data is safely stored in Google Drive and can be restored as needed.

## Development Instructions

### Backend Development
To build and run the backend separately for development:

1. **Build the Backend:**
   ```bash
   cd backend
   sudo docker build -t backend-dev .
   ```

2. **Run the Backend:**
   ```bash
   sudo docker run -p 8080:8080 backend-dev
   ```

This exposes the backend on port `8080`.

### Frontend Development
To build and run the frontend separately for development:

1. **Build the Frontend:**
   ```bash
   cd frontend
   sudo docker build -t frontend-dev .
   ```

2. **Run the Frontend:**
   ```bash
   sudo docker run -p 19006:19006 -p 19001:19001 -p 19002:19002 frontend-dev
   ```

This exposes the frontend on the necessary Expo ports.

## Known Issues and Resolutions

- **MongoDB Port Conflict**: If you encounter the error `Bind for 0.0.0.0:27017 failed: port is already allocated`, ensure no other MongoDB service or container is running on port `27017`. Alternatively, update the port mapping in `docker-compose.yml` to use a different port.

- **Dependency Conflicts**: Ensure that `expo` and `react-dom` are installed with the correct versions as described. Add these specific installs in the Dockerfile to maintain consistency.

- **Expo CLI Warning**: If you see a warning about the legacy Expo CLI, ensure you are using the latest Expo CLI by running `npm install -g expo-cli@latest`.

## Future Improvements
- Implement CI/CD pipelines for automatic deployment.
- Add unit and integration tests for both frontend and backend.
- Optimize Docker images to reduce build time and image size.

## Conclusion
This README provides the necessary steps to get ButlerAI up and running. By following the instructions, you should have a working development environment ready for further improvements.
