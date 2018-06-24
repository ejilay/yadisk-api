package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"github.com/ejilay/yadisk-api"
	"net/http"
	"time"
	"path"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <token> <disk_path>\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
		os.Exit(2)
		return
	}

	oAuthToken := flag.Arg(0)
	path := flag.Arg(1)


	client := &http.Client{
		Timeout: 600 * time.Second, // example
	}
	api := yadisk.NewAPI(oAuthToken, client)


	res, err := api.Share(path)
	if err != nil {
		log.SetPrefix("error: ")
		log.Println(err)
	}

	fmt.Println(res)
}
