# Use the official lightweight Node.js 18 image for pi
FROM arm64v8/node:18-alpine

# Set the working directory
WORKDIR /app

# Copy package.json, package-lock.json, and next.config.js
COPY frontend/package*.json ./
COPY frontend/next.config.js ./

# Install dependencies including 'sharp' for image optimization
RUN npm install && npm install sharp

# Copy the local code to the container's working directory
COPY frontend/ ./

# Build the application
RUN npm run build

# Run the application on the specified port
EXPOSE 3000
CMD ["npm", "start"]
