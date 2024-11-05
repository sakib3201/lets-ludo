FROM node:20-alpine

WORKDIR /app

COPY package.json package.json

RUN npm install

COPY . .

EXPOSE 6969

CMD [ "npm", "run", "dev" ]
