// the first
package main

import (
	"strings"
	"fmt"
	shell "github.com/ipfs/go-ifs-api"
)

func main() {
sh:=shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader("hello world!"))
	if err != nil {
		fmt.Fprontf(os.stderr,"error:%s",err)
		os.Exit(1)
	}
	fmt.Printf("added %s",cid)
}
