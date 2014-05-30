fluentail
=========

Receives UDP packets from [udp_stream](https://bitbucket.org/winebarrel/fluent-plugin-udp-stream) plugin and displays to standard output in JSON format

[ ![Download](https://api.bintray.com/packages/nyushi/fluentail/fluentail/images/download.png) ](https://bintray.com/nyushi/fluentail/fluentail/_latestVersion)

build
-----

```
go get -d
go build
```

usage
-----

### setup `udp_stream` in `fluentd`

```
<match any_tag.**>
  type udp_stream
  host <target_host>
</match>
```

### run `fluentail`

```
user@<target_host>$ fluentail
user@<target_host>$ fluentail | jq .
```
