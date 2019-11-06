package main

import (
	"github.com/IamStubborN/calendar/reminder/app"
	_ "github.com/lib/pq"
)

func main() {
	app.NewApp().Run()
}
