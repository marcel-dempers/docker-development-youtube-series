FROM node:12.4.0-alpine as debug

WORKDIR /work/

COPY ./src/package.json /work/package.json
RUN npm install
RUN npm install -g nodemon

COPY ./src/ /work/src/

ENTRYPOINT [ "nodemon","--inspect=0.0.0.0","./src/server.js" ]

FROM node:12.4.0-alpine as prod

WORKDIR /work/
COPY ./src/package.json /work/package.json
RUN npm install
COPY ./src/ /work/

CMD node .