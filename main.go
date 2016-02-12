package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/coreos/go-semver/semver"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	De "github.com/tj/go-debug"
)

var debug = De.Debug("add-bytes:main")

func main() {
	app := cli.NewApp()
	app.Name = "add-bytes"
	app.Usage = "cat file | add-bytes"
	app.Version = version()
	app.Action = run
	app.Flags = []cli.Flag{}
	app.Run(os.Args)
}

func run(context *cli.Context) {
	totalSize := uint64(0)
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		bytes, err := humanize.ParseBytes(line)
		FatalIfError(fmt.Sprintf("Failed to convert to bytes: %v", line), err)
		totalSize += bytes
	}

	if err := scanner.Err(); err != nil {
		FatalIfError("Failed to read from stdin", err)
	}

	fmt.Println(humanize.Bytes(totalSize))
}

func version() string {
	version, err := semver.NewVersion(VERSION)
	if err != nil {
		errorMessage := fmt.Sprintf("Error with version number: %v", VERSION)
		log.Panicln(errorMessage, err.Error())
	}
	return version.String()
}

// FatalIfError logs the error, then exits 1
func FatalIfError(message string, err error) {
	if err == nil {
		return
	}

	log.Fatalln(color.RedString(message), err.Error())
}
