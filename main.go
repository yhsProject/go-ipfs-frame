// the first
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/whyrusleeping/tar-utils"
)

var sh *shell.Shell

func main() {
	sh = shell.NewShell("localhost:5001")

	// upload
	// hash := IPFSUpload("hello,world")
	// fmt.Println(hash)

	// cat
	// readData := IPFSCat(hash)
	// fmt.Println(readData)

	// get
	// get := IPFSGet(hash, "hash.txt")

}

// IPFSGet Download hash content to file
func IPFSGet(hash, outdir string) error {
	resp, err := sh.Request("get", hash).Option("create", true).Send(context.Background())
	if err != nil {
		return err
	}
	defer resp.Close()
	if resp.Error != nil {
		return resp.Error
	}

	extractor := &tar.Extractor{Path: outdir}
	return extractor.Extract(resp.Output)
}

// IPFSUpload Upload string to ipfs
func IPFSUpload(str string) string {
	cid, err := sh.Add(strings.NewReader(str))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%s", err)
	}
	return string(cid)
}

//IPFSCat cat data in hash
func IPFSCat(hash string) string {
	read, err := sh.Cat(hash)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(read)
	return string(data)
}
