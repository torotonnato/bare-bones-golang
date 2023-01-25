package main

import (
	"encoding/json"
	"log"
	"os"

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

	l, _ := model.NewMetric("system.sensors.light", model.TYPE_GAUGE)
	tags_l := model.NewTags().Add("sensor:light").Add("prod:sensor")
	l.SetTags(tags_l)

	log.Println("Process started...")

	/*
		for count := 0; count < 100; count++ {
			temp := math.Sin(6.28 * float64(count) / 100.0)
			light := float64((count % 50) + 30)
			t.Sample(temp)
			l.Sample(light)
			time.Sleep(100 * time.Millisecond)
			log.Println(count)
		}
	*/
}
