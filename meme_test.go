package main

import (
	"fmt"
	"testing"
)

func TestGetMeme(t *testing.T) {
	url := getMeme("funny")
	fmt.Println(url)
}

func TestBadQuery(t *testing.T) {
	url := getMeme("jdwkdlawdl;awmdl;awmld;aw;ldmawl;dmawl;dm") //should return empty result from tenor
	fmt.Println(url)
}
