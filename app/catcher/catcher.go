package catcher

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Catcher struct {
	ctx       context.Context
	col       *mongo.Collection
	isStarted bool
	data      Gamelog
}

type Kill struct {
	Killer string `bson:"killer"`
	Victim string `bson:"victim"`
}

type Gamelog struct {
	Kills []Kill    `bson:"kills"`
	Date  time.Time `bson:"date"`
}

func New(ctx context.Context, dbconn, dbname string) (ct *Catcher, err error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbconn))
	if err != nil {
		return ct, err
	}
	var ctr Catcher
	ctr.col = client.Database(dbname).Collection("logs")
	ctr.ctx = ctx

	return &ctr, nil
}

func (c *Catcher) Do(data []byte) error {
	reader := bytes.NewReader(data)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		if c.isNewGame(scanner.Text()) {
			c.isStarted = true
			c.data = Gamelog{
				Date: time.Now(),
			}
		}

		kill := c.getKill(scanner.Text())
		if kill != nil && c.isStarted && kill.Killer != kill.Victim {
			c.data.Kills = append(c.data.Kills, *kill)
		}

		if c.isEndGame(scanner.Text()) && c.isStarted && len(c.data.Kills) > 0 {
			_, err := c.col.InsertOne(c.ctx, c.data)
			if err != nil {
				return fmt.Errorf("catcher.Do: can not insert data: %s", err)
			}
			c.isStarted = false
			c.data = Gamelog{}
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("can not scan data: %s", err)
	}

	return nil
}

func (c *Catcher) isNewGame(data string) bool {
	matched, err := regexp.MatchString(`Game_Start:\s\w`, data)
	if err != nil {
		log.Printf("[WARN] catcher.isNewGame: can not match string: %s", err)
	}

	return matched
}

func (c *Catcher) isEndGame(data string) bool {
	matched, err := regexp.MatchString(`Exit:\s`, data)
	if err != nil {
		log.Printf("[WARN] catcher.isEndGame: can not match string: %s", err)
	}

	return matched
}

func (c *Catcher) getKill(data string) *Kill {
	re, _ := regexp.Compile(`(?i)([a-z]+)\skilled\s([a-z]+)`)
	res := re.FindAllStringSubmatch(data, -1)
	if res == nil {
		return nil
	}

	k := Kill{
		Killer: res[0][1],
		Victim: res[0][2],
	}

	return &k
}
