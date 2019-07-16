# bashRPC
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
  - path: /uptime
    cmd: uptime

  - path: /tail/systemd
    cmd: grep systemd /var/log/syslog | tail -n 50

  - path: /deploy
    cmd: |
      cd /srv/webapp
      git pull
      ./script/start-app


```


2) start server
```
remote -c /path/to/config
```

3) ping server
```
curl -H "Authorization: supersecret" localhost:8675/uptime
```


