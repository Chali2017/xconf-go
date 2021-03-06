package main

// example

// Usually, you just copy 'example/config.go' and change the 'XconfConfig' struct define.
// After 'XconfInit', you can use Xconf() got the Config object.

import (
	"log"

	"github.com/one-go/xconf-go"
)

func main() {
	group := "xconf-sdk"
	name := "xconftest-test2.json"
	if err := XconfInit(&xconf.Options{
		Endpoints: []string{"test.riodev.oa.com:2379"},
		Username:  "",
		Password:  "",
		ID:        "", // use os.Hostname() when ID empty
	}, group, name); err != nil {
		log.Fatalf("xconf init fail err=%s", err.Error())
	}

	// use
	log.Printf("name: %s", Xconf().Name)
}
