# ButlerAI

## Overview
ButlerAI is a full-stack application utilizing Docker to streamline development and deployment. The project includes a Go-based backend, a React Native frontend powered by Expo, and MongoDB as the database. The setup is designed for ease of deployment and local development using Docker.

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
│   ├── Dockerfile
│   ├── init.sh
│   └── ...
├── docker-compose.yml
└── README.md
```

## Requirements
- **Docker**: Ensure Docker is installed on your system. [Install Docker](https://docs.docker.com/get-docker/)
- **Docker Compose**: Install Docker Compose. [Install Docker Compose](https://docs.docker.com/compose/install/)
- **Node.js and npm**: Required for running the frontend locally. [Install Node.js](https://nodejs.org/)

## Useful Commands

### 1.Check active docker containers
```bash
sudo docker ps
``` 

### 2.Stop docker container
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

### 3. Build and Run the Entire Project

To build and run the frontend, backend, and MongoDB together:

```bash
sudo docker-compose up --build
```

This will start all services:
- **Backend** on port `8080`
- **Frontend (Expo)** on ports `19006`, `19001`, and `19002`
- **MongoDB** on port `27017` (or `27018` if changed)

### 4. Stopping the Services
To stop the services:

```bash
sudo docker-compose down
```

## MongoDB Initialization and Data Management

### Automate MongoDB Import and Export

The MongoDB service has been configured to automatically import data upon startup and export data upon shutdown.

1. **Data Import on Startup:**
   - If a backup file (`/data/backup.archive`) exists inside the MongoDB container, it will automatically be restored when the container starts.

2. **Data Export on Shutdown:**
   - When the MongoDB container is stopped, the current data will be exported to `/data/backup.archive` to preserve the state for future use.

### Script Details:
- **Initialization Script (`init.sh`):**
  The `init.sh` script located in the `mongo/` directory handles the data import/export process.

```bash
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
```

### Dockerfile for MongoDB
The custom `Dockerfile` for MongoDB is located in the `mongo/` directory. It ensures the `init.sh` script is copied and executed during the container lifecycle.

```Dockerfile
FROM mongo:latest

# Copy the initialization script
COPY init.sh /docker-entrypoint-initdb.d/init.sh

# Set executable permissions for the script
RUN chmod +x /docker-entrypoint-initdb.d/init.sh
```

### Build and Run the MongoDB Service Separately

To build and run the MongoDB service separately for development:

1. **Build the MongoDB Image:**
   ```bash
   cd mongo
   sudo docker build -t mongo-dev .
   ```

2. **Run the MongoDB Container:**
   ```bash
   sudo docker run -p 27017:27017 mongo-dev
   ```

This will allow you to work with MongoDB independently of the other services.

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



