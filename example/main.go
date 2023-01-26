package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	agent "github.com/torotonnato/gobarebones/agent"
	api "github.com/torotonnato/gobarebones/api"
	config "github.com/torotonnato/gobarebones/config"
	model "github.com/torotonnato/gobarebones/model"
)

// Reads configuration file and checks API key validity
func setup() {
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
	valid, err := api.Validate()
	if !valid {
		log.Fatal(err.Error())
	}
	log.Println("DataDog API key validated.")
}

func main() {
	setup()

	t, _ := model.NewMetric("system.sensors.temp1", model.TYPE_GAUGE)
	tags_t := model.NewTags().Add("sensor:temp1").Add("prod:sensor")
	t.SetTags(tags_t)
	err := agent.RegisterMetric(t)
	if err != nil {
		log.Println(err)
	}

	l, _ := model.NewMetric("system.sensors.light", model.TYPE_GAUGE)
	tags_l := model.NewTags().Add("sensor:light").Add("prod:sensor")
	l.SetTags(tags_l)
	err = agent.RegisterMetric(l)
	if err != nil {
		log.Println(err)
	}
	log.Println("Process started...")
	agent.Start()

	count := 0
	for {
		temp1 := math.Sin(6.28 * float64(count) / 100.0)
		agent.PushMetric(t, temp1)
		light := float64((count % 50) + 30)
		agent.PushMetric(l, light)
		time.Sleep(1 * time.Second)
		status := fmt.Sprintf("tick= %d, temp1= %.2f, light= %.2f", count, temp1, light)
		log.Println(status)
		count++
	}
}
