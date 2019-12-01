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
