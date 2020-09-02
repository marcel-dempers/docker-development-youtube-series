FROM node:12.4.0-alpine as dev

RUN mkdir /work/
WORKDIR /work/

COPY ./src/package.json /work/package.json
RUN npm install

COPY ./src/ /work/

CMD node .
