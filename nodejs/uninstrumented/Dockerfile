# Use an official Node.js runtime as a parent image
FROM node:18

# Set the working directory
WORKDIR /app

# Copy package.json and install dependencies
COPY package.json package-lock.json ./
RUN npm install --omit=dev

# Copy the application files
COPY . .

# Expose the application's port
EXPOSE 3000

# Start the application
CMD ["node", "server.js"]
