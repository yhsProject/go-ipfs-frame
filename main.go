// the first
package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

var sh *shell.Shell

func main() {
	sh = shell.NewShell("localhost:5001")

	// Add
	hash, err := sh.Add(strings.NewReader("hello,world"))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(hash)

	// Get
	// sh.Get(hash, "hash.txt")

	// Cat
	readData, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(readData)
	fmt.Println(string(data))

}
