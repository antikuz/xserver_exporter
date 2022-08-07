# Xserver exporter
Prometheus exporter for xserver.a-real firewall. It collects Xserver statistics and exports them via HTTP for use by Prometheus.

# Installation


# Configuration
The image is setup to take parameters from environment variables or config.yaml:

The available environment variables are:

* `URL` Xserver address, example https://127.0.0.1:81
* `LOGIN` User to access xserver
* `PASSWD` Password to access xserver
* `INSECURE` Ignore server certificate verification, defaul false
* `LOGLEVEL` Sets the logging level, default info

When using a configuration file `config.yaml`:
* `url` Xserver address, example https://127.0.0.1:81
* `login` User to access xserver
* `passwd` Password to access xserver
* `insecure` Ignore server certificate verification, defaul false
* `logLevel` Sets the logging level, default info

# Dashboard
Grafana dashboard https://grafana.com/grafana/dashboards/16525
