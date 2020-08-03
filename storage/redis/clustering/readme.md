
## Replication

Documentation [here](https://redis.io/topics/replication)

### redis-0 Configuration

```
protected-mode no
port 6379
masterauth a-very-complex-password-here
requirepass a-very-complex-password-here
```
### redis-1 Configuration

```
protected-mode no
port 6379
slaveof redis-0 6379
masterauth a-very-complex-password-here
requirepass a-very-complex-password-here
```
### redis-2 Configuration

```
protected-mode no
port 6379
slaveof redis-0 6379
masterauth a-very-complex-password-here
requirepass a-very-complex-password-here
```



```
docker network create redis

cd .\storage\redis\clustering\

#redis-0
docker run -d --rm --name redis-0 `
    --net redis `
    -v ${PWD}/redis-0:/etc/redis/ `
    redis:6.0-alpine redis-server /etc/redis/redis.conf

#redis-1
docker run -d --rm --name redis-1 `
    --net redis `
    -v ${PWD}/redis-1:/etc/redis/ `
    redis:6.0-alpine redis-server /etc/redis/redis.conf


#redis-2
docker run -d --rm --name redis-2 `
    --net redis `
    -v ${PWD}/redis-2:/etc/redis/ `
    redis:6.0-alpine redis-server /etc/redis/redis.conf

```

## Running Sentinels

Documentation [here](https://redis.io/topics/sentinel)

Starting Redis in sentinel mode

```
cd .\storage\redis\clustering\

docker run -d --rm --name sentinel-0 --net redis `
    -v ${PWD}/sentinel-0:/etc/redis/ `
    redis:6.0-alpine `
    redis-sentinel /etc/redis/sentinel.conf

docker run -d --rm --name sentinel-1 --net redis `
    -v ${PWD}/sentinel-1:/etc/redis/ `
    redis:6.0-alpine `
    redis-sentinel /etc/redis/sentinel.conf

docker run -d --rm --name sentinel-2 --net redis `
    -v ${PWD}/sentinel-2:/etc/redis/ `
    redis:6.0-alpine `
    redis-sentinel /etc/redis/sentinel.conf
```
