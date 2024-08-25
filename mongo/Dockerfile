FROM mongo:latest

# Copy the initialization script into the appropriate directory
COPY init.sh /docker-entrypoint-initdb.d/init.sh

# Set executable permissions for the script
RUN chmod +x /docker-entrypoint-initdb.d/init.sh
