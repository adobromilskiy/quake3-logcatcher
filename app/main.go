package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/adobromilskiy/quake3-logcatcher/app/api"
	"github.com/adobromilskiy/quake3-logcatcher/app/catcher"
	"github.com/adobromilskiy/quake3-logcatcher/app/logfile"
	"github.com/jessevdk/go-flags"
)

var revision = "unknown"

var opts struct {
	DbConn    string        `long:"dbconn" description:"Database connection." required:"true"`
	DbName    string        `long:"dbname" description:"Database name." default:"quake3"`
	Path      string        `long:"path" description:"Path to the docker cocket or logfile." required:"true"`
	Socket    bool          `long:"socket" description:"Use socket connection or parse flogile."`
	Container string        `long:"container" description:"Container name." default:"quake3-server"`
	Interval  time.Duration `long:"interval" description:"Interval for api client runner." default:"10s"`
}

type Logcatcher interface {
	Run(ct *catcher.Catcher) error
}

func main() {
	fmt.Println("Revision:", revision)

	var err error

	if _, err := flags.Parse(&opts); err != nil {
		log.Fatalf("[ERROR] flags.Parse: %s", err)
	}

	var lc Logcatcher

	lc, err = logfile.NewClient(opts.Path)
	if opts.Socket {
		lc, err = api.NewClient(opts.Path, opts.Container, opts.Interval)
	}

	if err != nil {
		log.Fatalf("[ERROR] NewClient: %s", err)
	}

	catcher, err := catcher.New(context.Background(), opts.DbConn, opts.DbName)
	if err != nil {
		log.Fatalf("[ERROR] catcher.New: %s", err)
	}

	if err := lc.Run(catcher); err != nil {
		log.Fatalf("[ERROR] client.Run: %s", err)
	}
}
