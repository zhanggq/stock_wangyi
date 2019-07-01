package main

import (
	"PPGo_amaze/models"
	"PPGo_amaze/monitor"
	_ "PPGo_amaze/routers"

	"github.com/astaxie/beego"
	"github.com/robfig/cron"
)

func main() {
	// Initilize Database And Register Models to ORM
	models.Init()

	// Add cron job to update stock data every day
	go func() {
		cr := cron.New()
		//0 second, every minute for test
		//spec := "0 * * * * ?"
		//0 second, 0 minute, 2 hour, every day
		spec := "0 0 22 * * ?"
		cr.AddFunc(spec, func() {
			monitor.UpdateAndBackup()
		})
		cr.Start()
	}()

	// Add new goroutine to initilize all stock date
	// And it have to before beego.Run
	go func() {
		m := &monitor.Monitor{
			StartDate: "20100101",
			//StartDate: "20170101",
			EndDate: "20171231",
		}
		m.InitilizeDB()
	}()

	// Start listen port 80xx
	beego.Run()
}
