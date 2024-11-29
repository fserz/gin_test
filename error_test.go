package main

import (
	"fmt"
	"log"
	"testing"
)

func TestError(t *testing.T) {
	//bufio.NewReader()
	n, _ := fmt.Sscan("sheaetgbjlekag " +
		"asdfjlkadfnadkljfgn dagd")
	log.Print(n)
	//fmt.Println(n)
}
