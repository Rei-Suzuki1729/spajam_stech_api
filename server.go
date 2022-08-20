package main

import (
	
	"github.com/Rei-Suzuki1729/spajam_stech_api/router"
	
)

func main() {
	
	e := router.Initialize()
	e.Logger.Fatal(e.Start(":1323"))
	
}
