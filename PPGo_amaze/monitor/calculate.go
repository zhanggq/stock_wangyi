package monitor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	//"reflect"
	"strconv"
	"strings"
	"time"

	"PPGo_amaze/models"

	"github.com/astaxie/beego"
	iconv "github.com/djimenez/iconv-go"
)

const (
	urlBegin = "http://quotes.money.163.com/service/chddata.html?"
	urlEnd   = "&fields=TCLOSE;HIGH;LOW;TOPEN;LCLOSE;CHG;PCHG;TURNOVER;VOTURNOVER;VATURNOVER;TCAP;MCAP"
)

func sendMsg(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Header.Set("Content-Type", "application/json;charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		errText := fmt.Sprintf("status code: %d", resp.StatusCode)
		return nil, errors.New(errText)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//fmt.Println(len(body))
	//normal buffer len is 255551
	out := make([]byte, (len(body) + 10000))
	out = out[:]
	iconv.Convert(body, out, "gb2312", "utf-8")

	//fmt.Println(string(out))
	return out, nil
}

func checkDate(code, lastDate string) (bool, string, string) {
	timeNow := time.Now().Format("2006-01-02 15:04:05")
	timeDate := strings.Split(timeNow, " ")[0]
	year := strings.Split(timeDate, "-")[0]
	mon := strings.Split(timeDate, "-")[1]
	day := strings.Split(timeDate, "-")[2]
	endDate := year + mon + day
	startDate := "20100101"

	if timeDate > lastDate {
		t, _ := time.Parse("2006-01-02 15:04:05", lastDate+" 15:04:05")
		timeAfter := t.Add(24 * time.Hour).Format("2006-01-02 15:04:05")
		timeDate = strings.Split(timeAfter, " ")[0]
		year = strings.Split(timeDate, "-")[0]
		mon = strings.Split(timeDate, "-")[1]
		day = strings.Split(timeDate, "-")[2]
		startDate = year + mon + day
	} else {
		return false, "", ""
	}
	return true, startDate, endDate
}

/**********************************************
** 结构体可以作为函数的参数和返回值，如果结构体较大，一般使用指针参数，而且如果要在函数修改结构体，则必须使用指针形式
***********************************************/
func CalculateStock(stock models.StockName) (bool, string) {
	code := stock.Code
	lastDate := stock.LastDate
	res, startDate, endDate := checkDate(code, stock.LastDate)
	if res {
		beego.Info(fmt.Sprintf("[%s] start=%s end=%s.! ", code, startDate, endDate))
		urlBody := fmt.Sprintf("code=1%s&start=%s&end=%s", code, startDate, endDate)
		url := urlBegin + urlBody + urlEnd
		//fmt.Println("Req URL '`%s' ", url)
		respj, err := sendMsg(url)
		if err != nil {
			fmt.Fprintf(os.Stdout, "Err [%v]\n", err)
		}

		//insert into value table
		s := string(respj)
		//fmt.Println(reflect.TypeOf(s))
		ss := strings.Split(s, "\n")
		beego.Info(fmt.Sprintf("[%s] get data num[%d].! ", code, len(ss)))
		if len(ss) < 3 {
			beego.Info(fmt.Sprintf("[%s] no data, continue.! ", code))
			return false, ""
		}

		for idx := 1; idx < len(ss); idx++ {
			Value := new(models.Value)
			sss := strings.Split(ss[idx], ",")
			if len(sss) < 2 {
				continue
			}
			Value.Code = sss[1][1:]
			Value.Date = sss[0]
			Value.Tclose, _ = strconv.ParseFloat(sss[3], 64)
			Value.High, _ = strconv.ParseFloat(sss[4], 64)
			Value.Low, _ = strconv.ParseFloat(sss[5], 64)
			Value.Topen, _ = strconv.ParseFloat(sss[6], 64)
			Value.Lclose, _ = strconv.ParseFloat(sss[7], 64)
			Value.Chg, _ = strconv.ParseFloat(sss[8], 64)
			Value.Pchg, _ = strconv.ParseFloat(sss[9], 64)
			Value.Turnover, _ = strconv.ParseFloat(sss[10], 64)
			Value.Voturnover, _ = strconv.ParseInt(sss[11], 10, 64)
			Value.Vaturnover, _ = strconv.ParseInt(sss[12], 10, 64)
			Value.Tcap, _ = strconv.ParseInt(sss[13], 10, 64)
			Value.Mcap, _ = strconv.ParseInt(sss[14], 10, 64)

			if _, err := models.ValueAdd(Value); err != nil {
				beego.Error(fmt.Sprintf("[%s] Insert DB failed [%v] Error: [%v]! ", Value, err))
				return false, ""
			}

			if Value.Date > lastDate {
				lastDate = Value.Date
			}
		}
		beego.Info(fmt.Sprintf("[%s] insert into db finish.Last date [%s] ", code, lastDate))
	}
	return true, lastDate
}
