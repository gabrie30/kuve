# kuve

> Useful commands for working with kubernetes

# install

```
go get github.com/gabrie30/kuve
curl https://raw.githubusercontent.com/gabrie30/kuve/master/example_conf > $HOME/.kuve.yaml
```

# use

```
$ kuve --help
A collection of useful commands for working with kubernetes

Usage:
  kuve [command]

Available Commands:
  events      Get and filter events based off type from current context
  exec        Execs into the first running pod and container
  help        Help about any command
  images      Returns a list of images deployed into namespace
  logs        Get logs from pods and containers in a given namespace
  podnode     View which node a given pod in a given namespace is running on (gcp clusters only)
  pods        Returns pods given settings in .kuve.conf
  secrets     Base64 decode and view secrets from a given namespace
  version     A brief description of your command

Flags:
      --config string   config file (default is $HOME/.kuve.yaml)
  -h, --help            help for kuve
  -t, --toggle          Help message for toggle
```

```
$ kuve exec --help
Execs into the first running pod and container of a namespace

Usage:
  kuve exec [namespace] [flags]

Flags:
  -c, --container string   container to exec into
  -h, --help               help for exec
  -l, --selector string    selector (label query) to filter on
  -s, --shell string       shell to exec with (default "/bin/sh")
```

```
$ kuve exec helloworld --shell=/bin/bash
root@helloworld-768cc46c95-66vsm:/#
```
