fluentail
=========

Receives UDP packets from [udp_stream](https://bitbucket.org/winebarrel/fluent-plugin-udp-stream) plugin and displays to standard output in JSON format

build
-----

```
go get -d
go build
```

usage
-----

1. setup udp_stream in `fluentd`

```
<match any_tag.**>
  type udp_stream
  host <target_host>
  port <target_port>
</match>
```

1. run `fluentail`

```
user@target_host$ fluentail -l <target_port>
user@target_host$ fluentail -l <target_port> | jq .
```
