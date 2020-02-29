package main

import (
	"flag"
	"log"

	_ "github.com/kyoukaya/ak-discordrpc/discord"

	"github.com/kyoukaya/rhine/proxy"
)

var logPath = flag.String("log-path", "logs/proxy.log", "file to output the log to")
var silent = flag.Bool("silent", false, "don't print anything to stdout")
var filter = flag.Bool("filter", false, "enable the host filter")
var verbose = flag.Bool("v", false, "print Rhine verbose messages")
var verboseGoProxy = flag.Bool("v-goproxy", false, "print verbose goproxy messages")
var host = flag.String("host", ":8080", "hostname:port")
var disableCertStore = flag.Bool("disable-cert-store", false, "disables the built in certstore, reduces memory usage but increases HTTP latency and CPU usage")

func main() {
	flag.Parse()

	options := &proxy.Options{
		LogPath:          *logPath,
		LogDisableStdOut: *silent,
		EnableHostFilter: *filter,
		LoggerFlags:      log.Ltime | log.Lshortfile,
		Verbose:          *verbose,
		VerboseGoProxy:   *verboseGoProxy,
		Address:          *host,
		DisableCertStore: *disableCertStore,
	}
	rhine := proxy.NewProxy(options)
	rhine.Start()
}
