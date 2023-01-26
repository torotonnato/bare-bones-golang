package main

import (
	"fmt"
	"log"
	"math"
)

type procData struct {
	light, temp1 float64
}

func (p procData) log(tick int) procData {
	status := fmt.Sprintf(
		"tick= %d, temp1= %.2f, light= %.2f",
		tick,
		p.temp1,
		p.light,
	)
	log.Println(status)
	return p
}

func process(tick int) procData {
	return procData{
		light: math.Sin(6.28 * float64(tick) / 100.0),
		temp1: float64((tick % 50) + 30),
	}
}
