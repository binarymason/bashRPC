# go-remote
Simple HTTP server that executes configured commands remotely.


### Example Usage


1) create a config file

```
---

port: 8675
secret: supersecret

whitelisted_clients:
  - 127.0.0.1

routes:
  - path: /foo
    cmd: echo foo

  - path: /bar
    cmd: echo bar

  - path: /uptime
    cmd: uptime

```


2) start server
```
remote -c /path/to/config
```

3) ping server
```
curl -H "Authorization: supersecret" localhost:8675/uptime
```


