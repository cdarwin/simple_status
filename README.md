SimpleStatus
============

This repo contains a very simple tool written in [the Go Programming Language](http://golang.org/). It is intended to be a RESTful interface for obtaining certain system statistics from a server. It returns results in [JSON](http://www.json.org/) format.

## Configuration

`simple_status` accepts a few command line arguments to configure how it is accessed. Running `simple_status --help` will show you the following:

    Usage of ./simple_status:
    -p=":8080": http service address
    -ssl=false: TLS boolean flag
    -t="": http auth token

### Port

You may set the port to run the service on. This option is a string and is required to prepend with the ":" (colon) character. 

The default port is 8080 if left unset.

    -p=":8080": http service address

### Encryption

You may choose to wrap your connection in SSL for added security. This option is a boolean value of either `true` or `false`. 

This option is left off by defaults for simplicity.

    -ssl=false: TLS boolean flag

### Authentication

You may choose to set an authentication token for added security. This option is any string.

This option is left off by default.

    -t="": http auth token

## Upstart

We are using [upstart](http://upstart.ubuntu.com/) to manage running it in the background and have included a [sample](https://github.com/cdarwin/simple_status/blob/master/simple_status.conf) [job configuration file](http://upstart.ubuntu.com/cookbook/#job-configuration-file) to help facilitate this.

## Installation

We have included a very simple bootstrap [script](https://github.com/cdarwin/simple_status/blob/master/install.sh). You should take a careful look at it and modify it for your own environment before blindly deploying it. This is just a starting template to help deploy `simple_status` to all of your servers.

## Usage

You've got a server that you want to know some system stats about. You'll want to run the process in the background with something like the upstart script mentioned above. A typical `exec` line might look something like this:

    simple_status -ssl -p :9090 -t foobarbaz

Once you've got the status daemon running on the machines you'd like to monitor, next you'll want to request some useful information about them. Issuing a `GET` request to your host might look something like this:

    https://myhostname.com:9090/1/api/system?token=foobarbz

This should return a json blob of something to the effect of:

    {
      "host": "myhostname",
      "load": {
        "avg1": "0.56",
        "avg2": "0.51",
        "avg3": "0.53"
      },
      "ram": {
        "free": "4937256",
        "total": "8062936"
      },
      "time": "2012 08/01 0055-43"
    }

To get just the load averages, you would use `https://myhostname.com:9090/1/api/system/load?token=foobarbz`

A new `shell` endpoint allows you to execute arbitrary commands on the node this daemon is running on.
    
    curl -k -d "exec=whoami" -d "token=foobarbaz" https://myhostname.com:9090/1/api/shell
