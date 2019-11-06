package main

import (
	"github.com/IamStubborN/calendar/notifier/app"
	_ "github.com/lib/pq"
)

func main() {
	app.NewApp().Run()
}
