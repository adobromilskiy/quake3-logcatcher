package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

type Client struct {
	endpoint string
	client   *http.Client
	timeout  time.Duration
}

type Container struct {
	ID    string   `json:"Id"`
	Names []string `json:"Names"`
}

func NewClient(socket string, timeout time.Duration) (c *Client, err error) {
	if _, err := os.Stat(socket); os.IsNotExist(err) {
		return nil, fmt.Errorf("socket %s does not exist", socket)
	}

	tr := &http.Transport{
		Dial: func(proto, addr string) (conn net.Conn, err error) {
			return net.Dial("unix", socket)
		},
	}

	return &Client{
		endpoint: "http://localhost",
		client:   &http.Client{Transport: tr, Timeout: time.Second * 5},
		timeout:  timeout,
	}, nil
}

func (c *Client) Run() (err error) {
	for {
		containers, err := c.getcontainers()
		if err != nil {
			return err
		}

		log.Println(containers)

		time.Sleep(c.timeout)
	}
}

func (c *Client) getcontainers() (res []Container, err error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/containers/json", c.endpoint))
	if err != nil {
		return res, fmt.Errorf("api.getcontainers: can not get list of containers: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("api.getcontainers: can not read response body: %s", err)
	}

	if err := json.Unmarshal(body, &res); err != nil {
		return res, fmt.Errorf("api.getcontainers: can not unmarshal response body: %s", err)
	}

	return res, nil
}
