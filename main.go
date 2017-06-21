// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"flag"
	"os/exec"
	"time"
)

var (
	doneChan chan bool
)

type config struct {
	appFile string        `json:"appFile"`
}

func execFile(c *config) {
	if c.appFile == "" {
		fmt.Print("usage: -appFile appFile\n")
		os.Exit(1)
	}
	cmd := "./" + c.appFile
	out, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println("error occured")
		fmt.Printf("%s", err)
	}
	fmt.Printf("%s", out)

}

func makeFlags(c *config, fs *flag.FlagSet) {
	fs.StringVar(&c.appFile, "appFile", "", "App File will be executed.")
}


func main() {
	c := &config{}
	makeFlags(c, flag.CommandLine)
	flag.Parse()
	fmt.Printf("app file: %v", c.appFile)

	tickChan := time.NewTicker(time.Second).C
	doneChan = make(chan bool)

	for {
		select {
		case <- tickChan:
			execFile(c)
		case <- doneChan:
			return
		}
	}
}

