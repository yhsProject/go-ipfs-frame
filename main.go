// the first
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func main() {

	hash := IPFSUpload("hello")
	fmt.Println(hash)
	readData := IPFSCat(hash)
	fmt.Println(readData)
}

// IPFSUpload string
func IPFSUpload(str string) string {
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader(str))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%s", err)
	}
	return string(cid)
}

//IPFSCat download data
func IPFSCat(hash string) string {
	sh := shell.NewShell("localhost:5001")
	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(read)
	return string(data)
}
