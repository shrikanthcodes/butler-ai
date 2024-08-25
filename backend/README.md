To set up Docker for your entire project, including frontend, backend, and MongoDB, along with dependency management, follow these revised instructions.

### Step 1: Uninstall System-Wide MongoDB

Since you want to containerize everything, start by uninstalling the system-wide MongoDB:

1. **Stop MongoDB Service:**
   ```bash
   sudo systemctl stop mongod
   ```

2. **Uninstall MongoDB Packages:**
   ```bash
   sudo apt-get purge mongodb-org*
   ```

3. **Remove MongoDB Data Directories:**
   ```bash
   sudo rm -r /var/log/mongodb
   sudo rm -r /var/lib/mongodb
   ```

4. **Clean Up:**
   ```bash
   sudo apt-get autoremove
   sudo apt-get autoclean
   ```

### Step 2: Install Docker

Install Docker to manage your project dependencies:

1. **Install necessary packages:**
   ```bash
   sudo apt-get update
   sudo apt-get install \
     ca-certificates \
     curl \
     gnupg \
     lsb-release
   ```

2. **Add Docker’s official GPG key:**
   ```bash
   curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
   ```

3. **Set up the stable repository:**
   ```bash
   echo \
     "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
     $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
   ```

4. **Install Docker Engine:**
   ```bash
   sudo apt-get update
   sudo apt-get install docker-ce docker-ce-cli containerd.io
   ```

5. **Verify the installation:**
   ```bash
   sudo docker run hello-world
   ```

### Step 3: Set Up Docker for the Entire Project

Now that Docker is installed, you can containerize your entire project, including frontend, backend, and MongoDB.

#### 1. **Dockerize Backend (Go Project)**

1. **Create a `Dockerfile` for the backend:**
   In your backend directory (`/Documents/GitHub/butlerAI/backend`), create a `Dockerfile`:

   ```Dockerfile
   # Use the official Golang image as a build stage
   FROM golang:1.20-alpine AS builder

   # Set the working directory inside the container
   WORKDIR /app

   # Copy the go.mod and go.sum files and download dependencies
   COPY go.mod go.sum ./
   RUN go mod download

   # Copy the rest of the application source code
   COPY . .

   # Build the Go application
   RUN go build -o main ./cmd/butler

   # Use a minimal image to run the application
   FROM alpine:latest
   WORKDIR /root/
   COPY --from=builder /app/main .
   
   # Expose the port that your application runs on
   EXPOSE 8080
   
   # Run the Go app
   CMD ["./main"]
   ```

2. **Build the Docker image:**
   Navigate to the backend directory and build the image:

   ```bash
   cd /Documents/GitHub/butlerAI/backend
   sudo docker build -t my-go-backend .
   ```

#### 2. **Dockerize Frontend (React Project)**

1. **Create a `Dockerfile` for the frontend:**
   In your frontend directory (`/Documents/GitHub/butlerAI/frontend`), create a `Dockerfile`:

   ```Dockerfile
   # Use an official Node.js runtime as the base image
   FROM node:18-alpine

   # Set the working directory inside the container
   WORKDIR /app

   # Copy package.json and install dependencies
   COPY package.json yarn.lock ./
   RUN yarn install

   # Copy the rest of the application source code
   COPY . .

   # Build the React application
   RUN yarn build

   # Serve the React application using a simple HTTP server
   RUN yarn global add serve
   CMD ["serve", "-s", "build"]

   # Expose the port that the React app runs on
   EXPOSE 5000
   ```

2. **Build the Docker image:**
   Navigate to the frontend directory and build the image:

   ```bash
   cd /Documents/GitHub/butlerAI/frontend
   sudo docker build -t my-react-frontend .
   ```

#### 3. **Set Up MongoDB with Docker**

1. **Run the MongoDB Docker container:**
   You can start MongoDB in a Docker container as follows:

   ```bash
   sudo docker run -dp 27017:27017 -v local-mongo:/data/db --name local-mongo --restart=always mongo:latest
   ```

   This sets up MongoDB with persistent storage using a Docker volume.

#### 4. **Create a `docker-compose.yml` File**

To simplify running everything together, create a `docker-compose.yml` file in the root of your project directory (`/Documents/GitHub/butlerAI`):

```yaml
version: '3.8'

services:
  backend:
    build: ./backend
    ports:
      - "8080:8080"
    depends_on:
      - mongo

  frontend:
    build: ./frontend
    ports:
      - "5000:5000"

  mongo:
    image: mongo
    ports:
      - "27017:27017"
    volumes:
      - mongo-data:/data/db

volumes:
  mongo-data:
```

This configuration sets up three services:
- **backend**: Your Go backend.
- **frontend**: Your React frontend.
- **mongo**: MongoDB, with persistent storage.

#### 5. **Run Everything Together**

Navigate to your project root directory (`/Documents/GitHub/butlerAI`) and run:

```bash
sudo docker-compose up --build
```

This will build and start all your services (frontend, backend, and MongoDB) together. The frontend will be accessible on port 5000, the backend on port 8080, and MongoDB on port 27017.

### Step 4: Sharing MongoDB Data Across Environments

To ensure both you and your partner have the same MongoDB data:

1. **Export MongoDB data:**
   ```bash
   sudo docker exec local-mongo mongodump --out /data/dump
   sudo docker cp local-mongo:/data/dump ./local-dump
   ```

2. **Share and import the data:**
   Your partner can import the data into their MongoDB container using:

   ```bash
   sudo docker cp ./local-dump local-mongo:/data/dump
   sudo docker exec local-mongo mongorestore /data/dump
   ```

This setup ensures your project, including all dependencies, is containerized and can be easily shared across different environments.

Let me know if you need further assistance!
## How to run `backend`

- cd backend
- go build ./cmd/butler
- ./butler

## How to initialize go project

- cd butlerai
- go mod init backend