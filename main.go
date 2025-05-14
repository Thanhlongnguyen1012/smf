package main

import (
	"smf/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080") // REST API chạy trên port 8080
}
