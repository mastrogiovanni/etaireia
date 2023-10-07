FROM node:20.5.0
WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .
RUN npm run build
CMD [ "npm", "run", "dev" ]

