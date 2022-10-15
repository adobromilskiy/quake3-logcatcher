package catcher

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
)

type Catcher struct{}

func (c *Catcher) Do(data []byte) error {

	reader := bytes.NewReader(data)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		log.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("can not scan data: %s", err)
	}

	return nil
}
