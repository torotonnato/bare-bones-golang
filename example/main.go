package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	agent "github.com/torotonnato/gobarebones/agent"
	api "github.com/torotonnato/gobarebones/api"
	config "github.com/torotonnato/gobarebones/config"
	model "github.com/torotonnato/gobarebones/model"
)

// Reads configuration file and checks API key validity
func readConfig() {
	// Don't use json as configuration, this is just an
	// example
	confData, err := os.ReadFile("config.json")
	if err != nil {
		log.Fatal("Can't read from configuration file")
	}
	env := config.Config{}
	err = json.Unmarshal(confData, &env)
	if err != nil {
		log.Fatal(err)
	}
	config.Setup(&env)
}

func main() {
	readConfig()
	valid, err := api.Validate()
	if !valid {
		log.Fatal(err.Error())
	}
	log.Println("DataDog API key validated.")

	t, _ := model.NewMetric("system.sensors.temp1", model.TYPE_GAUGE)
	tTags := model.NewTags().Add("sensor:temp1").Add("prod:sensor")
	t.SetTags(tTags)
	if err := agent.RegisterMetric(t); err != nil {
		log.Println(err)
	}

	l, _ := model.NewMetric("system.sensors.light", model.TYPE_GAUGE)
	lTags := model.NewTags().Add("sensor:light").Add("prod:sensor")
	l.SetTags(lTags)
	if err := agent.RegisterMetric(l); err != nil {
		log.Println(err)
	}

	agent.Start()

	log.Println("Process started...")
	ticks := 0
	for {
		data := process(ticks).log(ticks)
		agent.PushMetric(l, data.light)
		agent.PushMetric(t, data.temp1)
		time.Sleep(1 * time.Second)
		ticks++
	}
}
