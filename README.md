SimpleStatus
============

SimpleStatus provides a REST/HTTP API for obtaining certain system statistics from a node. SimpleStatus is written in [the Go Programming Language](http://golang.org/).

[![Build Status][1]][2]
[1]: https://travis-ci.org/cdarwin/simple_status.png
[2]: https://travis-ci.org/cdarwin/simple_status

## Endpoints

<table>
  <tr>
    <th>URI</th>
    <th>HTTP Verb</th>
    <th>Purpose</th>
  </tr>
  <tr>
    <td>/system</td>
    <td>GET</td>
    <td>List all available resources</td>
  </tr>
  <tr>
    <td>/system/host</td>
    <td>GET</td>
    <td>List node hostname</td>
  </tr>
  </tr>
  <tr>
    <td>/system/disk</td>
    <td>GET</td>
    <td>List node disk usage</td>
  </tr>
  <tr>
    <td>/system/load</td>
    <td>GET</td>
    <td>List node load averages</td>
  </tr>
  <tr>
    <td>/system/ram</td>
    <td>GET</td>
    <td>List node RAM usage</td>
  </tr>
  <tr>
    <td>/shell</td>
    <td>GET</td>
    <td>Execute arbitrary shell commands on the node</td>
  </tr>
</table>

## Authentication

SimpleStatus currently uses a very simple method for authenticating API requests. The daemon may be configured to require authentication at runtime using the `-t` switch. This token may be any arbitrary string at the time of this writing.

**Example Query with Parameter**:

`GET https://nodename.com:8080/1/api/system?token=abc4c7c627376858`

## Requests

Requests to the API are simple HTTP requests against the API endpoints.

All request bodies should be in JSON, with Content-Type of `application/json`.

### Base URL

A few parameters may be set at runtime which will affect the bare URL that you will use as the prefix for the desired endpoint.

* `-ssl`: Enforce encryption of communication with the API
* `-p`: Specify the port for communication with the API (defaults to 8080)

All endpoints should be prefixed with something similar to the following:

`{scheme}://{nodename}:{port}/1/api`

## Responses

All responses are in JSON, with Content-Type of `application/json`. A response is structured as follows:

`{ "resource_name": "resource value" }`

---

## System

Overview of all available resources

**Endpoint** 

`GET /system`

**Optional URI Parameters**

* `disk`: Specify a path or device for disk usage
    *  If none is specified, "/" is assumed

**Response**

    {
      "host": "mario",
      "disk": {
        "all": 117623562240,
        "used": 8339341312,
        "free": 109284220928
      },
      "cpuinfo": {
        "processors": 4,
        "siblings": 4,
        "cores": 2
      },
      "load": {
        "avg1": 0.26,
        "avg2": 0.23,
        "avg3": 0.23
      },
      "ram": {
        "free": 565848,
        "total": 7871876
      },
      "time": "2013 02/01 2323-54"
    }

## Host

System hostname

**Endpoint** 

`GET /system/host`

**Response**

    "nexus2"

## Disk

Used, free, and total disk space available for a given device or path. If no `disk` parameter is provided, the endpoint assumes the `/` path.

**Endpoint** 

`GET /system/disk`

**Optional URI Parameters**

* `disk`: Specify a path or device for disk usage
    *  If none is specified, "/" is assumed

**Response**

    {
      "all": 35439468544,
      "used": 20696563712,
      "free": 14742904832
    }

## CpuInfo

From `/etc/cpuinfo`

**Endpoint** 

`GET /system/cpuinfo`

Response:

    {
      "processors": 4,
      "siblings": 4,
      "cores": 2
    }

## Load

Load averages for the node

**Endpoint** 

`GET /system/load`

Response:

    {
      "avg1": 0.40,
      "avg2": 0.40,
      "avg3": 0.37
    }

## RAM

RAM usage for the node

**Endpoint** 

`GET /system/ram`

Response:

    {
      "free": 20608,
      "total": 2060976
    }

## Shell

Execute arbitrary shell commands on the node. This endpoint is only accessible if a token was specified at runtime

**Endpoint** 

`GET /shell`

**URI Parameters**

* `exec`: Specify the command to be executed on the node
* `token`: Mandatory authentication token

**Response**

    "command response"
