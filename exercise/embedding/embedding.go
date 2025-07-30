//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import (
	"fmt"
)

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandwidth() int {
	sum := 0
	for _, value := range b.amount {
		sum += int(value)
	}
	return sum / len(b.amount)
}

func (c *CpuTemp) AverageTemp() int {
	sum := 0
	for _, value := range c.temp {
		sum += int(value)
	}
	return sum / len(c.temp)
}

func (m *MemoryUsage) AverageMemory() int {
	sum := 0
	for _, value := range m.amount {
		sum += int(value)
	}
	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}
	dashboard := Dashboard{bandwidth, temp, memory}
	fmt.Println(dashboard.AverageBandwidth())
	fmt.Println(dashboard.AverageTemp())
	fmt.Println(dashboard.AverageMemory())
}

