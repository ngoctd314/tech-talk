package main

import "kafka-talk/demo"

func main() {
	// demo.PartitionRoundRobin()
	// demo.PartitionMurmur2()
	// demo.PartitionCustom()
	// demo.ProduceWithoutBatching()
	demo.ProduceWithBatching()
}
