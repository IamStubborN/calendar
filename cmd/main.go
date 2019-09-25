package main

import (
	"github.com/IamStubborN/calendar/app"
	_ "github.com/lib/pq"
)

func main() {
	app.NewApp().Run()
}
