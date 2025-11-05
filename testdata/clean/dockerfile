FROM node:18

# Set working directory
WORKDIR /app

COPY package*.json ./

RUN npm install

# Copy source code
COPY . .

EXPOSE 3000

CMD ["npm", "start"]
