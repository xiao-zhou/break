package main

import (
	"testing"
)

func Test_ExecFile(t *testing.T) {
	c := config{}
	execFile(&c)
}

func Test_ExecFile_File(t *testing.T) {
	c := config{appFile: "appFile"}
	execFile(&c)
}
