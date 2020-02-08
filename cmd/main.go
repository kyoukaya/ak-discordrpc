package main

import (
	"log"

	_ "github.com/kyoukaya/ak-discordrpc/discord"

	"github.com/kyoukaya/rhine/proxy"
)

func main() {
	options := &proxy.Options{
		LoggerFlags: log.Ltime | log.Lshortfile,
	}
	rhine := proxy.NewProxy(options)
	rhine.Start()
}
