package main

import (
	"log"
	"math"
	"time"

	bb "github.com/torotonnato/gobarebones"
)

func main() {
	config := bb.Config{
		Region: bb.DD_EU,
		APIKey: "(redacted)",
	}
	err := bb.Setup(&config)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("APIKey validated...")

	t, _ := bb.NewMetric("system.sensors.temp1", bb.METRIC_GAUGE)
	tags_t := bb.NewMetricTags().Add("sensor:temp1").Add("prod:sensor")
	t.SetTags(tags_t)

	l, _ := bb.NewMetric("system.sensors.light", bb.METRIC_GAUGE)
	tags_l := bb.NewMetricTags().Add("sensor:light").Add("prod:sensor")
	l.SetTags(tags_l)

	log.Println("Process started...")

	for count := 0; count < 100; count++ {
		temp := math.Sin(6.28 * float64(count) / 100.0)
		light := float64((count % 50) + 30)
		t.Sample(temp)
		l.Sample(light)
		time.Sleep(100 * time.Millisecond)
		log.Println(count)
	}
}
