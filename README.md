# bashRPC
Simple HTTP server that executes configured commands remotely.


### Installation

Grab the latest binary and put it in somewhere in your path.  Check the [releases](https://github.com/binarymason/bashRPC/releases) page. Alternatively, you can download a package.

Example for Debian systems:
```bash
# Note: this link may be out of date.  Be sure to check releases page to get latest version
wget https://github.com/binarymason/bashRPC/releases/download/v19.07.16-1002/bashrpc-v19.07.16-1002.deb
sudo apt install bashrpc-v19.07.16-1002.deb

```

### Example Usage


1) Create a config file.  If you are using bashRPC as a system service, the config is located at `/etc/bashrpc/bashrpc.yml`

```yml
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

```bash
bashrpc -c /path/to/config
```

If you installed bashRPC with your package manager, you can alternatively start bashRPC as a system service:

Example for systemd:
```bash
sudo systemctl daemon-reload
sudo systemctl enable bashrpc
sudo systemctl start bashrpc

```

3) ping server

```bash
curl -H "Authorization: supersecret" localhost:8675/uptime
```


