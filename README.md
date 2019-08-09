# kuve

> Useful commands for working with kubernetes

# install

```
go get github.com/gabrie30/kuve
```

# use

```
$ kuve --help
Available Commands:
  exec        Execs into the first running pod and container of a namespace
  help        Help about any command
  images      Returns a list of images deployed into namespace
  logs        Get logs from pods and containers in a given namespace
  secrets     Base64 decode and view secrets from a given namespace
```

```
$ kuve exec --help
Execs into the first running pod and container of a namespace

Usage:
  kuve exec [namespace] [flags]

Flags:
  -c, --container string   container to exec into
  -h, --help               help for exec
  -l, --selector string    selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)
  -s, --shell string       shell to exec with (default "/bin/sh")
```

```
kuve exec helloworld --shell=/bin/bash
root@helloworld-768cc46c95-66vsm:/#
```
