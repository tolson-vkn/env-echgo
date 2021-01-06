# env-echgo

A small echo server in docker!

This is a go port of a project I use quite freqently for testing services and/or demoing kube scheduling and load balancing.
This project was previously done in [Python](https://github.com/tolson-vkn/env-echo). It's been pulled about 3k times in the last year, and the python version is 214MB.
Yikes. As of writing ech**go** is 10.4MB woot!

An improvement over the python version is that any envar in the container specified with `ECHGO_*` adds to the payload.

# Usage

There are no command line args. Maybe someday.

* Server runs at 8080
* Add key/values to output by adding an envar with `ECHGO_` prefix
  * e.g. `ECHGO_DEPLOY_COLOR=Orange` shows as `{"deploy_color": "Orange"}`

For Docker: 

```
docker run --rm -i -t -e ECHGO_FOO=bar -p 8080:8080 timmyolson/env-echgo
```

For Kube (see `kubectl explain pod.spec.containers.env.valueFrom`):

```
env:
- name: ECHGO_POD_NAME
  valueFrom:
    fieldRef:
      fieldPath: metadata.name
```

## Demo

```
$ kubectl apply -f k8s/
$ curl -s 10.5.1.151:8080 | jq
{
  "message": "Hello from Ecgho Server",
  "node_name": "alyx-01",
  "pod_ip": "192.168.139.135",
  "pod_name": "env-echgo-bbb58fb4-zqwm7"
}
$ curl -s 10.5.1.151:8080 | jq
{
  "message": "Hello from Ecgho Server",
  "node_name": "alyx-02",
  "pod_ip": "192.168.152.91",
  "pod_name": "env-echgo-bbb58fb4-4hdwz"
}
$ curl -s 10.5.1.151:8080 | jq
{
  "message": "Hello from Ecgho Server",
  "node_name": "alyx-01",
  "pod_ip": "192.168.139.134",
  "pod_name": "env-echgo-bbb58fb4-qhc28"
}
$ curl -s 10.5.1.151:8080 | jq
{
  "message": "Hello from Ecgho Server",
  "node_name": "alyx-01",
  "pod_ip": "192.168.139.135",
  "pod_name": "env-echgo-bbb58fb4-zqwm7"
}
```
