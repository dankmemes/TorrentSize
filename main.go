package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	humanize "github.com/dustin/go-humanize"
	"github.com/labstack/gommon/color"
)

var stats = struct {
	Size    int
	Index   int
	NbFiles int
}{}

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]")
var tildPre = color.Yellow("[") + color.Green("~") + color.Yellow("]")
var crossPre = color.Yellow("[") + color.Red("✗") + color.Yellow("]")

func main() {
	var worker sync.WaitGroup
	var count int
	var err error

	stats.Index = 1

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
	stats.NbFiles = len(files)
	for _, f := range files {
		worker.Add(1)
		count++
		go calculateSize(f.Name(), &worker)
		if count == arguments.Concurrency {
			worker.Wait()
			count = 0
		}
	}

	worker.Wait()

	fmt.Println(checkPre +
		color.Yellow(" [") +
		color.Green(stats.Index-1) +
		color.Yellow("/") +
		color.Green(stats.NbFiles) +
		color.Yellow("] ") +
		color.Green(" Total size: ") +
		color.Yellow(humanize.Bytes(uint64(stats.Size))))
}
