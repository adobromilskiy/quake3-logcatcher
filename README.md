# quake3-logcatcher

Simple application to parse logs from dockerized [quake3-server](https://github.com/adobromilskiy/quake3-server) via docker [API](https://docs.docker.com/engine/api/) in realtime and save them to a mongodb.

### Application parameters

Required parameters has no default value.

| parameter | default | description |
|-----------|---------|-------------|
| dbconn    |          | database connection string |
| dbname    | `quake3` | database name |
| socket    | `/run/docker.sock` | path to docker.sock |
| container |          | container name for parsing |
| timeout | `10s`      | timoute between requests for parsing |


### Quick start

```console
docker run --network mongo_network \
	-v /var/run/docker.sock:/run/docker.sock \
	--restart always --name quake3-logcatcher \
	adobromilskiy/quake3-logcatcher:latest \
	--dbconn=mongodb://mongohost:27017 --dbname=quake3 --container=quake3-server
```