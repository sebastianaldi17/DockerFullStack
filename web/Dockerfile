FROM node:16.2.0

RUN mkdir -p /web
COPY . /web

WORKDIR /web
RUN npm install

EXPOSE 3000

CMD npm run start