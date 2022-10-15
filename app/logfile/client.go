package logfile

import (
	"fmt"
	"os"

	"github.com/adobromilskiy/quake3-logcatcher/app/catcher"
)

type Client struct {
	file string
}

func NewClient(file string) (c *Client, err error) {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return nil, fmt.Errorf("file %s does not exist", file)
	}

	return &Client{
		file: file,
	}, nil
}

func (c *Client) Run() error {
	logs, err := os.ReadFile(c.file)
	if err != nil {
		return fmt.Errorf("can not read file %s: %s", c.file, err)
	}

	ct := catcher.Catcher{}
	if err := ct.Do(logs); err != nil {
		return fmt.Errorf("can not catch logs: %s", err)
	}

	return nil
}
