package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/rinosukmandityo/atlasgo/helper"
)

func main() {
	var configLoc string
	flag.StringVar(&configLoc, "config", filepath.Join(helper.WD, "configs", "configs.json"), "config file location")
	flag.Parse()

	configs := new(helper.AppConfig)
	helper.ReadJsonFile(configLoc, configs)

	t := time.Now()
	conn, e := helper.NewConnection(configs.URI)
	if e != nil {
		os.Exit(0)
	}
	defer conn.Close()
	log.Println("connection duration", time.Since(t))

	t = time.Now()
	conn.ShowResult = configs.ShowResult
	if configs.PerformanceTest {
		conn.PerformanceTest(configs.CollectionTest)
		log.Println("TOTAL PERFORMANCE TEST DURATION ======", time.Since(t))
	}
}
