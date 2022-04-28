package main

import (
	"fmt"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

// use os package to get the env variable which is already set
func envVariable(key string) string {

	// set env variable using os package
	os.Setenv(key, "go-setting-cronjob-process")

	// return the env variable using os package
	return os.Getenv(key)
}

func main() {
	// os package
	fmt.Println("--------- os read env variable ---------")
	value := envVariable("name")

	fmt.Printf("os package: name = %s \n", value)
	fmt.Printf("environment = %s \n", os.Getenv("APP_CRONJOB"))

	if event := os.Getenv("APP_CRONJOB"); event == "start" {
		log.Info("--------- Create new cron ---------")
		c := cron.New()

		// c.AddFunc("*/1 * * * *", printJobFirst)
		setCronJob(c)

		// Start cron with one scheduled job
		log.Info("Start cron")
		c.Start()
		printCronEntries(c.Entries())
	}

	log.Info("Start Something")

	time.Sleep(3 * time.Minute)
	fmt.Println("--------- programe finish ---------")

	// Funcs may also be added to a running Cron
	// log.Info("Add new job to a running cron")
	// entryID2, _ := c.AddFunc("*/2 * * * *", printJobSecond)
	// printCronEntries(c.Entries())
	// time.Sleep(5 * time.Minute)

	//Remove Job2 and add new Job2 that run every 1 minute
	// log.Info("Remove Job2 and add new Job2 with schedule run every minute")
	// c.Remove(entryID2)
	// c.AddFunc("*/1 * * * *", printJobSecond)
	// printCronEntries(c.Entries())
	// time.Sleep(5 * time.Minute)
}

func printCronEntries(cronEntries []cron.Entry) {
	log.Infof("Cron Info: %+v\n", cronEntries)
}

func printJobFirst() {
	log.Info("[Job 1]-Every minute job\n")
}

func printJobSecond() {
	log.Info("[Job 2]-Every two minutes job\n")
}

func setCronJob(c *cron.Cron) {
	c.AddFunc("*/1 * * * *", printJobFirst)
	// c.AddFunc("*/2 * * * *", printJobSecond)
}
