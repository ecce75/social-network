# Use the official lightweight Node.js 16 image.
# https://hub.docker.com/_/node
FROM node:16-alpine

# Set the working directory
WORKDIR /app

# Copy package.json, package-lock.json, and next.config.js
COPY frontend/package*.json ./
COPY frontend/next.config.js ./

# Install dependencies
RUN npm install

# Copy the local code to the container's working directory
COPY frontend/ ./

# Build the application
RUN npm run build

# Run the application on the specified port
EXPOSE 3000
CMD ["npm", "start"]
