# Introduction to FluentBit 

We always start with the [official documentation](https://docs.fluentbit.io/manual)


## Running fluentbit

```
# starts fluentbit with default config
docker run -it cr.fluentbit.io/fluent/fluent-bit:4.0

# see HELP text 
docker run -it cr.fluentbit.io/fluent/fluent-bit:4.0 -h
```

## Configure fluentbit

We've created a new [fluent-bit.yaml](./fluent-bit.yaml) with the example content from the documentation. </br>

We can update our `docker` command to mount the config and tell `fluent-bit` to use the new configuration file we created 

```shell
docker run -it -p 2020:2020 \
-v ./monitoring/fluentbit/fluent-bit.yaml:/fluent-bit/etc/fluent-bit.yaml \
cr.fluentbit.io/fluent/fluent-bit:4.0 --config=/fluent-bit/etc/fluent-bit.yaml
```

## Inputs

The example config file has a `random` input. </br>
Let's create two inputs, one reading log files and another accepting HTTP inputs 

```yaml
pipeline:
  inputs:
  - name: tail
    path: /var/lib/docker/containers/*/*.log
    multiline.parser: docker

  - name: http
    listen: 0.0.0.0
    port: 8888
```

We'll need to add a volume mount so our fluent-bit container has access to the container logs, so let's mount `/var/lib/docker/` 

We also need port `8888` to match our input plugin port

```shell
docker run -it \
-p 2020:2020 \
-p 8888:8888 \
-v ./monitoring/fluentbit/fluent-bit.yaml:/fluent-bit/etc/fluent-bit.yaml \
-v /var/lib/docker:/var/lib/docker \
cr.fluentbit.io/fluent/fluent-bit:4.0 --config=/fluent-bit/etc/fluent-bit.yaml

```

Test our HTTP input plugin using `curl` 

```shell
curl -d '{"namespace":"product","message":"hello from product"}' -X POST -H "content-type: application/json" http://localhost:8888/product.log
```

## Outputs

In the example config file we had a `stdout` output plugin that writes all data from the input plugin to the terminal. </br>
Let's create another output plugin that writes data to file so we can inspect it:

```yaml
- name: file
  match: '*'
  template: '{log}'
  path: /fluent-bit/.data
```

Let's add a volume mount so we can see the output data. </br>
Let's mount a local `.data` folder to `/fluent-bit/.data` 

```shell
docker run -it \
-p 2020:2020 \
-p 8888:8888 \
-v ./monitoring/fluentbit/fluent-bit.yaml:/fluent-bit/etc/fluent-bit.yaml \
-v /var/lib/docker:/var/lib/docker \
-v ./monitoring/fluentbit/.data:/fluent-bit/.data \
cr.fluentbit.io/fluent/fluent-bit:4.0 --config=/fluent-bit/etc/fluent-bit.yaml
```

## Concepts

There are key [concepts](https://docs.fluentbit.io/manual/concepts/key-concepts#events-or-records) to learn about fluent-bit, now that we have events being ingested. </br>

### Filter plugin

This plugin only includes requests from "product" namespace
```yaml
filters:
- name: grep
  match: '*'
  regex: namespace product
```

This plugin excludes all requests from "search" namespace

```yaml
filters:
- name: grep
  match: '*'
  exclude: namespace search
```

Test our HTTP input plugin using `curl` to send different events to our fluent-bit service:

```shell

#request allowed
curl -d '{"namespace":"product","message":"hello from product"}' -X POST -H "content-type: application/json" http://localhost:8888/product.log

# request dropped\filtered
curl -d '{"namespace":"search","message":"hello from search"}' -X POST -H "content-type: application/json" http://localhost:8888/search.log

```