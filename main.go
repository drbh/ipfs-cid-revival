package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	// "github.com/davecgh/go-spew/spew"
	shell "github.com/ipfs/go-ipfs-api"

	"github.com/ipfs/go-ipfs-api/options"

	"text/tabwriter"
)

func main() {

	s := shell.NewLocalShell()
	data, e := ioutil.ReadFile("file.json")

	if e != nil {
		panic(e)
	}

	fmt.Println("THE DATA")
	fmt.Println(data, "\n")

	hashing_algos := []string{
		"id",
		"sha1",
		"sha2-256",
		"sha2-512",
		"sha3-224",
		"sha3-256",
		"sha3-384",
		"sha3-512",
		"murmur3",
		"keccak-224",
		"keccak-256",
		"keccak-384",
		"keccak-512",
		"shake-128",
		"shake-256",
	}
	w := new(tabwriter.Writer)
	const padding = 6
	w.Init(os.Stdout, 0, 0, padding, ' ', 0)

	fmt.Println("THE HASHES")

	for _, element := range hashing_algos {
		c, err := s.DagPutWithOpts(
			data,
			options.Dag.Pin("true"),
			options.Dag.InputEnc("json"),
			options.Dag.Kind("dag-pb"),
			options.Dag.Hash(element),
		)
		if err != nil {
			fmt.Println(err)
		}
		s := []string{element, c}
		fmt.Fprintln(w, strings.Join(s, "\t"))
	}
	w.Flush()

}
