package monitor

import (

	//"reflect"

	"PPGo_amaze/comm"
	"PPGo_amaze/models"

	//"github.com/astaxie/beego"

	//log "github.com/Sirupsen/logrus"

	"github.com/astaxie/beego"
)

type Monitor struct {
	StartDate string
	EndDate   string
}

func InitializeChuangye() {
	list, total := models.Name300GetList()
	//beego.Info("InitializeChuangye Goroutine Id: ", comm.GetSlow())
	beego.Info("ChuangYe total :", total)
	for _, v := range list {
		res, lastDate := CalculateStock(v.StockName)

		if res {
			//update latest date to pp_name_300 table
			newName, err := models.Name300GetById(v.Id)
			if err != nil {
				beego.Error("Get [%s] by ID [%d] failed [%v] ", v.Code, v.Id, err)
				continue
			}
			newName.LastDate = lastDate
			beego.Info("Update [", v.Code, "] Date :", lastDate)
			if err = newName.Update(); err != nil {
				beego.Error("Update [", v.Code, "] failed :", err)
			}
		}
	}
}

func InitializeShenzhen() {
	list, total := models.Name000GetList()
	//beego.Info("InitializeShenzhen Goroutine Id: ", comm.GetSlow())
	beego.Info("ShenZhen total :", total)
	for _, v := range list {
		res, lastDate := CalculateStock(v.StockName)
		if res {
			//update latest date to pp_name_000 table
			newName, err := models.Name000GetById(v.Id)
			if err != nil {
				beego.Error("Get [%s] by ID [%d] failed [%v] ", v.Code, v.Id, err)
				continue
			}
			newName.LastDate = lastDate
			beego.Info("Update [", v.Code, "] Date :", lastDate)
			if err = newName.Update(); err != nil {
				beego.Error("Update [", v.Code, "] failed :", err)
			}
		}
	}
}

func (m *Monitor) InitilizeDB() {
	beego.Info("InitilizeDB Goroutine Id: ", comm.GetSlow())
	beego.Info("Begin Initilize Database! ")
	go func() {
		InitializeChuangye()
	}()
	go func() {
		InitializeShenzhen()
	}()
	beego.Info("Initilize Finish! ")
}
