package main

import (
	"flag"
	"log"

	_ "github.com/kyoukaya/ak-discordrpc/discord"

	"github.com/kyoukaya/rhine/proxy"
)

func main() {
	logPath := flag.String("log-path", "logs/proxy.log", "file to output the log to")
	silent := flag.Bool("silent", false, "don't print anything to stdout")
	filter := flag.Bool("filter", false, "enable the host filter")
	verbose := flag.Bool("v", false, "print Rhine verbose messages")
	verboseGoProxy := flag.Bool("v-goproxy", false, "print verbose goproxy messages")
	host := flag.String("host", ":8080", "hostname:port")
	flag.Parse()

	options := &proxy.Options{
		LogPath:          *logPath,
		LogDisableStdOut: *silent,
		EnableHostFilter: *filter,
		LoggerFlags:      log.Ltime | log.Lshortfile,
		Verbose:          *verbose,
		VerboseGoProxy:   *verboseGoProxy,
		Address:          *host,
	}
	rhine := proxy.NewProxy(options)
	rhine.Start()
}
