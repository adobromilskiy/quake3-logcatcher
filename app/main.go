package main

import (
	"io"
	"net"
	"net/http"
)

const sock = "/docker.sock"

func main() {
	tr := &http.Transport{
		Dial: fakeDial,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get("http://localhost/containers/json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// read resp.Body to string
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))

}

func fakeDial(proto, addr string) (conn net.Conn, err error) {
	return net.Dial("unix", sock)
}
