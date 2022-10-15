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

	"github.com/adobromilskiy/quake3-logcatcher/app/catcher"
)

type Client struct {
	endpoint  string
	client    *http.Client
	timeout   time.Duration
	container string
}

type Container struct {
	ID    string   `json:"Id"`
	Names []string `json:"Names"`
}

func NewClient(socket, container string, timeout time.Duration) (c *Client, err error) {
	if _, err := os.Stat(socket); os.IsNotExist(err) {
		return nil, fmt.Errorf("socket %s does not exist", socket)
	}

	tr := &http.Transport{
		Dial: func(proto, addr string) (conn net.Conn, err error) {
			return net.Dial("unix", socket)
		},
	}

	return &Client{
		endpoint:  "http://localhost",
		client:    &http.Client{Transport: tr, Timeout: time.Second * 5},
		timeout:   timeout,
		container: fmt.Sprintf("/%s", container),
	}, nil
}

func (c *Client) Run() (err error) {
	for {
		container, err := c.getcontainer()
		if err != nil {
			return err
		}

		if container.ID == "" {
			log.Printf("[WARN] api.Run: container %s does not exist", c.container)
			time.Sleep(c.timeout)
			continue
		}

		logs, err := c.getlogs(container.ID)
		if err != nil {
			return err
		}

		ct := catcher.Catcher{}
		if err := ct.Do(logs); err != nil {
			return fmt.Errorf("can not catch logs: %s", err)
		}

		time.Sleep(c.timeout)
	}
}

func (c *Client) getcontainer() (res Container, err error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/containers/json", c.endpoint))
	if err != nil {
		return res, fmt.Errorf("api.getcontainers: can not get list of containers: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("api.getcontainers: can not read response body: %s", err)
	}

	var containers []Container

	if err := json.Unmarshal(body, &containers); err != nil {
		return res, fmt.Errorf("api.getcontainers: can not unmarshal response body: %s", err)
	}

	for _, container := range containers {
		if container.Names[0] == c.container {
			return container, nil
		}
	}

	return res, nil
}

func (c *Client) getlogs(container string) (res []byte, err error) {
	resp, err := c.client.Get(fmt.Sprintf("%s/containers/%s/logs?stdout=1&stderr=1", c.endpoint, container))
	if err != nil {
		return res, fmt.Errorf("api.getlogs: can not get logs for container %s: %s", container, err)
	}
	defer resp.Body.Close()

	res, err = io.ReadAll(resp.Body)
	if err != nil {
		return res, fmt.Errorf("api.getlogs: can not read response body: %s", err)
	}

	return res, nil
}
