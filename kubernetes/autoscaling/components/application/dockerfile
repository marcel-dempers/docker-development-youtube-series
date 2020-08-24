FROM golang:1.14-alpine as build

RUN apk add --no-cache git curl

WORKDIR /src 

COPY app.go /src 

RUN go build app.go

FROM alpine as runtime

COPY --from=build /src/app /app/app

CMD [ "/app/app" ]