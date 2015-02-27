# Docker stats to backstop

Feeds [graphit](http://graphite.readthedocs.org/en/latest/) thru [backstop](https://github.com/obfuscurity/backstop) with [docker](https://www.docker.com/) stats

"Docker stats to backstop" is a litle daemon they collect stats from docker daemon and they translate it in [carbon](http://graphite.readthedocs.org/en/latest/feeding-carbon.html) data structure before send it in JSON to Backstop daemon.

# Requirement 

docker stats API is available from docker api v1.17, deliver in docker 1.5

# Deployment

Docker stats to backstop support 3 parameters :
* --backstop : URL for connecting backsop server
* --docker : URL for connecting docker server like "unix:///var/run/docker.sock" or "http://docker:2375"
* --prefix : JSON containing 'regexp' and 'into' to rewrite the container name into graphite identifier like {"regexp":"(.*)\\..*\\.(.*)\\..*","into":"$1.$2"} to have `my.container.name` identify in graphite by `my.name`
* --poll : Polling interval. Valid time units are 'ns', 'us' (or 'µs'), 'ms', 's', 'm', 'h'"

All this parameters can be set by environement :
* SRV_BACKSTOP for --backstop option
* SRV_DOCKER for --docker option
* SRV_PREFIX for --prefix option
* SRV_POLL for --poll
 
To run it you can use : 

    ./docker-stats-to-backstop \
      --backstop=http://backstop.server:8080/publish/docker \
      --prefix='{"regexp":"(.*)\\..*\\.(.*)\\..*","into":"$1.$2"}' \
      --poll=60s

If you prefer run it with systemd :

    [Unit]
    Description=Docker stats collector
    After=docker.service
    Requires=docker.service
    
    [Service]
    ExecStart=/usr/bin/docker-stats-to-backstop \
        --backstop=https://user:password@backstop.server/publish/docker \
        --prefix={"regexp":"(.*)\\\\..*\\\\.(.*)\\\\..*","into":"$1.$2"} \
        --poll=60s
    Restart=on-failure
    
    [Install]
    WantedBy=multi-user.target

# Build

Install and build it with : 

    go get github.com/deliverous/docker-stats-to-backstop

# Godo
We love [godo](https://github.com/go-godo/godo) for continuous integration on our desktop. 
Install it with : 

    go get -u gopkg.in/godo.v1/cmd/godo

and run it

    godo -w
