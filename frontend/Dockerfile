# Use the latest stable Node.js image
FROM node:18-bullseye-slim

# Set environment to development by default
ARG NODE_ENV=development
ENV NODE_ENV=$NODE_ENV

# Default ports for Expo
ARG PORT=19006
ENV PORT=$PORT
EXPOSE 19006 19001 19002

# Install global packages (latest npm)
RUN npm install -g npm@latest expo-cli@latest

# Create and set the working directory
RUN mkdir /opt/react_native_app
WORKDIR /opt/react_native_app

# Copy package.json and package-lock.json first to leverage Docker cache
COPY package.json package-lock.json ./

# Install Expo package and other dependencies using legacy peer deps
RUN npm install --legacy-peer-deps

# Install specific versions for compatibility
RUN npm install expo@~51.0.31
RUN npm install react-dom@18.2.0

# Copy the rest of the project files
COPY . .

# Ensure the entrypoint uses npm to start the web version of the Expo app
ENTRYPOINT ["npm", "run"]
CMD ["web"]
