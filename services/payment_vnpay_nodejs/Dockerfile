FROM oven/bun:1.3.8


WORKDIR /app

COPY package*.json ./
RUN npm install

COPY . .

EXPOSE 8080
CMD [ "npm", "run", "start" ]