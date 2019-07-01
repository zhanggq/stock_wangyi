package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"PPGo_amaze/models"
	//"github.com/astaxie/beego"
)

type Stockindex struct {
	BaseController
}

// [date, open, close, lowest, highest, volumn]
// define(
// [["2004-01-02",10452.74,10409.85,10367.41,10554.96,168890000],["2004-01-05",10411.85,10544.07,10411.85,10575.92,221290000]]

func (self *Stockindex) Index399() {
	fmt.Println("Into Index399")
	valueList := make([][]string, 0)
	self.Data["pageTitle"] = "深成指"

	list, total := models.ValueGet("399001")
	fmt.Println("Total %d", total)
	for _, v := range list {
		value := make([]string, 0)
		value = append(value, v.Date)
		value = append(value, strconv.FormatFloat(v.Topen, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.Tclose, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.Low, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.High, 'f', 2, 64))
		value = append(value, strconv.FormatInt(v.Voturnover, 10))
		valueList = append(valueList, value)
		//fmt.Println(value)
		//fmt.Println(valueList)
	}
	jValue, err := json.Marshal(valueList)
	if err != nil {
		fmt.Println("encoding faild")
	}
	//else {
	//fmt.Println("encoded data : ")
	//fmt.Println(jValue)
	//fmt.Println(string(jValue))
	//}
	self.Data["StockIndex"] = string(jValue)
	//self.ServeJSON()
	//self.display()
	self.TplName = "stockindex/399001.html"
}

func (self *Stockindex) Index300() {
	fmt.Println("Into Index300")
	valueList := make([][]string, 0)
	self.Data["pageTitle"] = "创业版"

	list, total := models.ValueGet("399006")
	fmt.Println("Total %d", total)
	for _, v := range list {
		value := make([]string, 0)
		value = append(value, v.Date)
		value = append(value, strconv.FormatFloat(v.Topen, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.Tclose, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.Low, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.High, 'f', 2, 64))
		value = append(value, strconv.FormatInt(v.Voturnover, 10))
		valueList = append(valueList, value)
		//fmt.Println(value)
		//fmt.Println(valueList)
	}
	jValue, err := json.Marshal(valueList)
	if err != nil {
		fmt.Println("encoding faild")
	}
	//else {
	//fmt.Println("encoded data : ")
	//fmt.Println(jValue)
	//fmt.Println(string(jValue))
	//}
	self.Data["StockIndex"] = string(jValue)
	//self.ServeJSON()
	//self.display()
	self.TplName = "stockindex/399006.html"
}

func (self *Stockindex) Index() {
	fmt.Println("test Stockindex")

	stockCode := strings.TrimSpace(self.GetString("code"))
	stockType := strings.TrimSpace(self.GetString("type"))
	stockFrom := strings.TrimSpace(self.GetString("from"))
	fmt.Println("Stockindex %s %s %s", stockCode, stockType, stockFrom)
	if 6 != len(stockCode) {
		stockCode = "000001"
	}

	valueList := make([][]string, 0)
	self.Data["pageTitle"] = "个股"

	stockName := "平安银行"
	if "创业板" == stockFrom {
		newName, err := models.Name300GetByCode(stockCode)
		if err != nil {
			fmt.Println("Get [%s] failed [%v]\n", stockCode, err)
			stockCode = "000001"
		} else {
			stockName = newName.Name
		}
	} else {
		newName, err := models.Name000GetByCode(stockCode)
		if err != nil {
			fmt.Println("Get [%s] failed [%v]\n", stockCode, err)
			stockCode = "000001"
		} else {
			stockName = newName.Name
		}
	}

	list, total := models.ValueGet(stockCode)
	fmt.Println("Code %s, Name %s, Total %d", stockCode, stockName, total)
	for _, v := range list {
		value := make([]string, 0)
		value = append(value, v.Date)
		value = append(value, strconv.FormatFloat(v.Topen, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.Tclose, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.Low, 'f', 2, 64))
		value = append(value, strconv.FormatFloat(v.High, 'f', 2, 64))
		value = append(value, strconv.FormatInt(v.Voturnover, 10))
		valueList = append(valueList, value)
		//fmt.Println(value)
		//fmt.Println(valueList)
	}
	jValue, err := json.Marshal(valueList)
	if err != nil {
		fmt.Println("encoding faild")
	}
	//else {
	//fmt.Println("encoded data : ")
	//fmt.Println(jValue)
	//fmt.Println(string(jValue))
	//}
	self.Data["StockIndex"] = string(jValue)
	if "click" == stockType {
		self.ajaxMsgCode(string(jValue), stockName)
		return
	}
	self.TplName = "stockindex/stockindex.html"
}
