package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/anacrolix/torrent"
	humanize "github.com/dustin/go-humanize"
	"github.com/labstack/gommon/color"
)

var client = http.Client{}

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]")
var tildPre = color.Yellow("[") + color.Green("~") + color.Yellow("]")
var crossPre = color.Yellow("[") + color.Red("✗") + color.Yellow("]")

func main() {
	var size int64
	i := 1
	c, _ := torrent.NewClient(nil)
	defer c.Close()

	// Parse arguments
	parseArgs(os.Args)

	// Check if input folder exist
	if _, err := os.Stat(arguments.Input); os.IsNotExist(err) {
		fmt.Println(crossPre +
			color.Yellow(" [") +
			color.Red(arguments.Input) +
			color.Yellow("] ") +
			color.Red("Invalid input folder."))
	}

	files, err := ioutil.ReadDir(arguments.Input)
	if err != nil {
		log.Fatal(err)
	}
	nbFiles := len(files)
	for _, f := range files {
		torrent, err := c.AddTorrentFromFile(arguments.Input + "/" + f.Name())
		if err != nil {
			fmt.Println(crossPre +
				color.Yellow(" [") +
				color.Red(i) +
				color.Yellow("/") +
				color.Red(nbFiles) +
				color.Yellow("] ") +
				color.Red("Error: ") +
				color.Yellow(err.Error()))
		}
		<-torrent.GotInfo()
		for _, file := range torrent.Files() {
			size += file.Length()
		}
		fmt.Println(checkPre +
			color.Yellow(" [") +
			color.Green(i) +
			color.Yellow("/") +
			color.Green(nbFiles) +
			color.Yellow("] ") +
			color.Green(" Extracted size from ") +
			color.Yellow(f.Name()))
		torrent.Drop()
		i++
	}

	fmt.Println(checkPre +
		color.Yellow(" [") +
		color.Green(i-1) +
		color.Yellow("/") +
		color.Green(nbFiles) +
		color.Yellow("] ") +
		color.Green(" Total size: ") +
		color.Yellow(humanize.Bytes(uint64(size))))
}
