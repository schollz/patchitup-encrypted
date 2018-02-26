package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/schollz/patchitup-encrypted/patchitup"
)

var (
	doDebug    bool
	port       string
	dataFolder string
	server     bool
	rebuild    bool
	whoami     bool
	pathToFile string
	identity   string
	address    string
)

func main() {

	flag.StringVar(&port, "port", "8002", "port to run server")
	flag.StringVar(&pathToFile, "f", "", "path to the file to patch")
	flag.StringVar(&identity, "i", "", "identity on the cloud")
	flag.StringVar(&address, "s", "", "server name")
	flag.StringVar(&dataFolder, "data", "", "folder to data (default $HOME/.patchitup)")
	flag.BoolVar(&doDebug, "debug", false, "enable debugging")
	flag.BoolVar(&server, "host", false, "enable hosting")
	flag.BoolVar(&rebuild, "rebuild", false, "rebuild file")
	flag.BoolVar(&whoami, "whoami", false, "get identity")
	flag.Parse()

	if doDebug {
		patchitup.SetLogLevel("debug")
	} else {
		patchitup.SetLogLevel("info")
	}

	if dataFolder != "" {
		patchitup.DataFolder = dataFolder
	}

	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {
	var public, private string
	if len(strings.Split(identity, "-")) == 2 {
		public = strings.Split(identity, "-")[0]
		private = strings.Split(identity, "-")[1]
	}
	if server {
		patchitup.SetLogLevel("info")
		err := patchitup.Run(port)
		if err != nil {
			return err
		}
	} else if whoami {
		p, err := patchitup.New(patchitup.Configuration{
			PathToFile:    pathToFile,
			ServerAddress: address,
			PublicKey:     public,
			PrivateKey:    private,
		})
		if err != nil {
			return err
		}
		fmt.Println(p.Identity())
	} else if rebuild {
		p, err := patchitup.New(patchitup.Configuration{
			PathToFile:    pathToFile,
			ServerAddress: address,
			PublicKey:     public,
			PrivateKey:    private,
		})
		if err != nil {
			return err
		}
		latest, err := p.Rebuild()
		if err != nil {
			return err
		}
		fmt.Println(latest)
	} else {
		p, err := patchitup.New(patchitup.Configuration{
			PathToFile:    pathToFile,
			ServerAddress: address,
			PublicKey:     public,
			PrivateKey:    private,
		})
		if err != nil {
			return err
		}
		err = p.PatchUp()
		if err != nil {
			return err
		}
	}
	return nil
}
