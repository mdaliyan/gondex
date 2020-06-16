package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"testing"
)

var tmp = path.Join(os.TempDir(), "gondex")

func TestMain(m *testing.M) {
	os.RemoveAll(tmp)
	os.MkdirAll(tmp, 0755)
	for k := 0; k < 40; k++ {
		file, err := os.Create(fmt.Sprintf(tmp+"/%d.txt", k))
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
	}
	os.Exit(m.Run())
}

func TestRead(t *testing.T) {

}
