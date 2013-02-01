# Installation

We have included a very simple bootstrap [script](https://github.com/cdarwin/simple_status/blob/master/install.sh). You should take a careful look at it and modify it for your own environment before blindly deploying it. This is just a starting template to help deploy `simple_status` to all of your servers.

## Upstart

One simple way to manage running SimpleStatus in the background is to use [upstart](http://upstart.ubuntu.com/). I have included a [sample](https://github.com/cdarwin/simple_status/blob/master/simple_status.conf) [job configuration file](http://upstart.ubuntu.com/cookbook/#job-configuration-file) as a template for what you might use in the real world.

## Usage

A typical `exec` line might look something like this:

    simple_status -ssl -p :9090 -t foobarbaz

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
