package main

import (
	appInit "targeting-engine/init/app"
	prometheusInit "targeting-engine/init/prometheous"
)

// Start Application
func main() {
	prometheusInit.InitPrometheus()
	appInit.InitEnvironment()

}
