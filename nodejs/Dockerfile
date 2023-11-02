FROM node:14.15.5-alpine
COPY package.json package-lock.json ./
RUN apk --no-cache add build-base
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN npm install
COPY . .

CMD [ "node", "app.js" ]