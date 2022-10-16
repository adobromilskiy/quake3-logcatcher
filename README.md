# quake3-logcatcher

Simple application to parse logs from dockerized [quake3-server](https://github.com/adobromilskiy/quake3-server) via docker [API](https://docs.docker.com/engine/api/) in realtime and save them to a mongodb. Or you can parse qconsole.log file.

### Application parameters

Required parameters has no default value.

| parameter | default | description |
|-----------|---------|-------------|
| dbconn    |          | database connection string |
| dbname    | `quake3` | database name |
| path      |          | path to docker.sock or logfile qconsole.log |
| socket    |          | choose read from socket or qconsole.log file |
| container | `quake3-server` | container name for parsing |
| interval  | `10s`    | interval between requests for parsing |


### Quick start

Run and parse logs via docker api:

```console
docker run --network mongo_network \
	-v /var/run/docker.sock:/run/docker.sock \
	--restart always --name quake3-logcatcher \
	adobromilskiy/quake3-logcatcher:latest \
	--dbconn=mongodb://mongohost:27017 --dbname=quake3 --container=quake3-server --socket --path="/run/docker.sock"
```

or via qconsole.log file:

```console
docker run --network mongo_network \
	-v ~/projects/quake3-logcatcher/qconsole.log:/qconsole.log \
	--name quake3-logcatcher \
	adobromilskiy/quake3-logcatcher:latest \
	--dbconn=mongodb://mongohost:27017 --dbname=quake3 --path="/qconsole.log"
```

Use [quake3-stats](https://github.com/adobromilskiy/quake3-stats) application to view game statistics.