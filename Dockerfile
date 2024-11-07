# Use a Node.js base image
FROM node:18-alpine

# Set the working directory
WORKDIR /.

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

#non root user
USER 1000
# Expose the port the app will listen on
EXPOSE 80

# Start the app
CMD ["node", "index.js"]