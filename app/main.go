package main

import (
	"fmt"
	"log"
	"time"

	"github.com/adobromilskiy/quake3-logcatcher/app/api"
	"github.com/jessevdk/go-flags"
)

var revision = "unknown"

var opts struct {
	DbConn    string        `long:"dbconn" description:"Database connection." required:"true"`
	DbName    string        `long:"dbname" description:"Database name." default:"quake3"`
	Socket    string        `long:"path" description:"Path to the docker cocket file." default:"/run/docker.sock"`
	Container string        `long:"container" description:"Container name." required:"true"`
	Timeout   time.Duration `long:"timeout" description:"Timeout for api client runner." default:"10s"`
}

func main() {
	fmt.Println("Revision:", revision)

	if _, err := flags.Parse(&opts); err != nil {
		log.Fatalf("[ERROR] flags.Parse: %s", err)
	}

	client, err := api.NewClient(opts.Socket, opts.Timeout)
	if err != nil {
		log.Fatalf("[ERROR] api.NewClient: %s", err)
	}

	if err := client.Run(); err != nil {
		log.Fatalf("[ERROR] client.Run: %s", err)
	}
}
