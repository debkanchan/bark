# BARK: Update base image to latest LTS version
FROM node:18

# Set working directory
WORKDIR /app

# BARK: Remove development dependencies for production builds
COPY package*.json ./

RUN npm install

# Copy source code
COPY . .

# BARK: Consider exposing a configurable port instead of hardcoding
EXPOSE 3000

# BARK: Change to 'npm ci' for deterministic installs
CMD ["npm", "start"]
