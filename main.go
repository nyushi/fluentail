package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"regexp"
	"time"

	"github.com/ugorji/go/codec"
)

func bytes2string(i interface{}) interface{} {
	switch v := i.(type) {
	case map[string]interface{}:
		for key, next := range v {
			v[key] = bytes2string(next)
		}
		return v
	case []interface{}:
		for key, next := range v {
			v[key] = bytes2string(next)
		}
		return v
	case []byte:
		str := string(v)
		return str
	default:
		return v
	}
}

type inputData struct {
	Tag    []byte                 `codec:"toarray"`
	Time   float64                `codec:"toarray"`
	Record map[string]interface{} `codec:"toarray"`
}

func (e *inputData) TimeString() string {
	return time.Unix(int64(e.Time), 0).Format("2006/01/02 15:04:05 MST")
}

func (e *inputData) Readable() map[string]interface{} {
	o := map[string]interface{}{
		"tag":      e.Tag,
		"time":     e.TimeString(),
		"unixtime": e.Time,
		"record":   e.Record,
	}
	return bytes2string(o).(map[string]interface{})
}

func decodeInput(b []byte) (map[string]interface{}, error) {
	var (
		mh codec.MsgpackHandle
	)
	v := inputData{}

	dec := codec.NewDecoderBytes(b, &mh)
	err := dec.Decode(&v)
	if err != nil {
		return nil, err
	}

	return v.Readable(), nil
}

func setupSocket(addrstr string) (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", addrstr)
	if err != nil {
		return nil, err
	}
	return net.ListenUDP("udp", addr)
}

func main() {
	var (
		listenPortFlag int
		tagRegexpFlag  string
		remoteAddrFlag string
	)
	flag.IntVar(&listenPortFlag, "l", 25000, "udp port for listen.")
	flag.StringVar(&tagRegexpFlag, "t", ".*", "filter regexp for tag. (e.g. 'warning$')")
	flag.StringVar(&remoteAddrFlag, "r", "", "filter string for remote addr. (e.g. '127.0.0.1')")
	flag.Parse()

	tagFilter := regexp.MustCompile(tagRegexpFlag)

	var buf []byte
	buf = make([]byte, 65507)
	addr := fmt.Sprintf(":%d", listenPortFlag)
	sock, err := setupSocket(addr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		rlen, remote, err := sock.ReadFromUDP(buf)
		if err != nil {
			log.Printf("read error: %s", err.Error())
		}

		decoded, err := decodeInput(buf[:rlen])
		if err != nil {
			log.Printf("msgpack decode error: %s", err.Error())
			continue
		}

		if !tagFilter.MatchString(decoded["tag"].(string)) {
			continue
		}
		if remoteAddrFlag != "" && remoteAddrFlag != fmt.Sprint(remote.IP) {
			continue
		}

		decoded["remote"] = fmt.Sprintf("%s:%d", remote.IP, remote.Port)

		b, err := json.Marshal(decoded)
		if err != nil {
			log.Printf("json encode error: %s", err.Error())
		}

		fmt.Println(string(b))
	}
}
