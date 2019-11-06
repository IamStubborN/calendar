package main

import (
	"github.com/IamStubborN/calendar/api/app"
	_ "github.com/lib/pq"
)

func main() {
	app.NewApp().Run()
}
