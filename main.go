package main

import (
	appInit "targeting-engine/init/app"
)

// Start Application
func main() {
	appInit.InitPrometheus()
	appInit.InitEnvironment()

}
