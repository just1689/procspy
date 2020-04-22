![Go](https://github.com/lunatik13/procspy/workflows/Go/badge.svg?branch=master)
Go module to list all TCP connections, with an option to try to find the owning PID and processname.

Works by reading /proc directly on Linux, and by executing `netstat` and `lsof -i` on Darwin.

Works for IPv4 and IPv6 TCP connections. Only established connections are listed; ports where something is only listening or TIME_WAITs are skipped.

If you want to find all processes you'll need to run this as root.

Status:
-------

Tested on Linux and Darwin (10.9).

Install:
--------

`go install`

Usage:
------

Only list the connections:

```
cs, err := procspy.Connections(false)
for c := cs.Next(); c != nil; c = cs.Next() {
    ...
}
```

List the connections and try to find the owning process:

```
cs, err := procspy.Connections(true)
for c := cs.Next(); c != nil; c = cs.Next() {
    ...
}
```

(See ./example\_test.go)

``` go

package main

import (
	"fmt"
	"github.com/lunatik13/procspy"
)

func main() {
	var tcpCloseWait uint = 8
	lookupProcesses := true
	var result string
	cs, err := procspy.Connections(lookupProcesses, tcpCloseWait)
	if err != nil {
		panic(err)
	}

	fmt.Println("TCP Connections:")
	for c := cs.Next(); c != nil; c = cs.Next() {
	    result =  result + fmt.Sprintf(" - %v\n", c)
	}
	fmt.Println(result)
}
```
