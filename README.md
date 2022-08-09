# Xserver exporter
Prometheus exporter for xserver.a-real firewall. It collects Xserver statistics and exports them via HTTP for use by Prometheus.

# Usage

### Binary release
You can download the latest release on the [release page](https://github.com/antikuz/xserver_exporter/releases).

## Docker container
Docker images are push to [docker hub](https://hub.docker.com/r/antikuz/xserver-exporter).

### Installing as Windows Service

1. Download binary
2. Install [nssm](https://nssm.cc/)

| Action  | Command                                                         |
| ------- | --------------------------------------------------------------- |
| install | nssm install xserver_exporter C:\xserver_exporter.exe \[args\]  |
| remove  | nssm remove xserver_exporter confirm                            |

# Build

### Build Binary
```shell
go build
```

### Build Docker Image
```shell
docker build . -t xserver-exporter
```

### 

# Configuration
The image is setup to take parameters from command flags, environment variables or config file:

Accept flags:
```bash
-u, --url string           Xserver configuration file path.
-l, --login string         User account to authenticate.
-p, --passwd string        User account password.
-i, --insecure             Allow insecure server connections when using SSL
    --log-level string     the maximum level of messages that should be logged. (possible values: debug, info, warn, error) (default "info")
-c, --config-file string   xserver configuration file path.
-h, --help                 Show help.
```

The available environment variables are:
* `URL` Xserver address, example https://127.0.0.1:81
* `LOGIN` User to access xserver
* `PASSWD` Password to access xserver
* `INSECURE` Ignore server certificate verification, default false
* `LOGLEVEL` Sets the logging level, default info

When using a configuration file:
* `url` Xserver address, example https://127.0.0.1:81
* `login` User to access xserver
* `passwd` Password to access xserver
* `insecure` Ignore server certificate verification, default false
* `logLevel` Sets the logging level, default info

# Dashboard
Grafana dashboard https://grafana.com/grafana/dashboards/16525
