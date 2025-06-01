package coreutils

var (
	//ServerPort The port on which the application will run
	ServerPort = ":8080"
	//Basepath The basepath to which callback requests will group
	Basepath = "/"
	// HealthCheckBasepath is the basepath for the health check
	HealthCheckBasepath = "/health"
)
