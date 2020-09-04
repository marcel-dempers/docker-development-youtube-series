

```
docker run -it -v ${PWD}:/work -w /work nginx bash
```

```
docker run -it -v ${PWD}/nginx.conf:/etc/nginx/nginx.conf -p 80:80 nginx
```

