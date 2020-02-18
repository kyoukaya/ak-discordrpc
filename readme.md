# ak-discordrpc

![discord rich pressence](https://i.imgur.com/HBCna5J.jpg)

A [rhine](https://github.com/kyoukaya/rhine) module to update a user's Discord rich presence when logging into the game.
The status text is limited to combat (practice, auto, fighting) and non-combat (idling only).

## Usage

Download the latest release from the [releases page](https://github.com/kyoukaya/ak-discordrpc/releases) or run from source with `go build cmd/main.go && ./main.exe`.

If you're using the program for the first time, you'll need to setup the certificate and change the proxy settings in your Android emulator/phone. [See here](https://github.com/kyoukaya/rhine/wiki/First-Time-Setup) for more information.

```
$ ./main.exe -help
Usage of C:\Users\kaya\Desktop\discordrpc\main.exe:
  -filter
        enable the host filter
  -host string
        hostname:port (default ":8080")
  -log-path string
        file to output the log to (default "logs/proxy.log")
  -silent
        don't print anything to stdout
  -v    print Rhine verbose messages
  -v-goproxy
        print verbose goproxy messages
```
