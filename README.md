# bashRPC
Simple HTTP server that executes configured commands remotely.

### Why use bashRPC instead of chef/ansible/saltstack/etc?

Use bashRPC when you don't want to give complete super user privileges. That prevents situations like:

```
salt "*" cmd.run "rm -rf /"
# - or -
ansible -i production all -a "rm -rf /"
```

Instead, you can configure an endpoint that does only a select few super user tasks, such as restarting a system service, etc.

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
$ curl -k -H "Authorization: supersecret" https://localhost:8675/uptime
```

### Security

There are few security measures implemented in bashRPC:

* No HTTP traffic.  HTTPS is required.
* User can specify their own SSL certificate, if desired.
* Restricted to whitelist of IP addresses.
* `Authorization` header is required for authentication on every request.
* No parameterized inputs. Every command must be pre-configured in `bashrpc.yml`.

### Output

bashRPC returns plain text responses, very similar if you were to be executing a command over SSH. This makes it easy to save responses to a variable, check for status code, etc. Both STDOUT and STDERR are combined in the output.

```
$ curl -k -H "Authorization: supersecret" https://localhost:8675/uptime
14:31:29 up 1 day,  1:16,  2 users,  load average: 1.77, 1.47, 1.43
```

If you care about whether or not your command fails, you can check the response.  Using `curl`, for example, you can exit non-zero if a command fails using the `--fail` argument:

```
$ curl -k -H --fail "Authorization: supersecret" https://localhost:8675/iwillfail
iwillfail: command not found

$ echo "$?"
1
```
