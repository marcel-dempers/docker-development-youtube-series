# Introduction to Fluentd

<a href="https://youtu.be/Gp0-7oVOtPw" title="fluentd-intro"><img src="https://i.ytimg.com/vi/Gp0-7oVOtPw/hqdefault.jpg" width="20%" alt="fluentd-intro" /></a> 

## Collecting logs from files

Reading logs from a file we need an application that writes logs to a file. <br/>
Lets start one:

```
cd monitoring\logging\fluentd\introduction\

docker-compose up -d file-myapp

```

To collect the logs, lets start fluentd

```
docker-compose up -d fluentd
```

## Collecting logs over HTTP (incoming)

```
cd monitoring\logging\fluentd\introduction\

docker-compose up -d http-myapp

```