package models

import (
	"github.com/astaxie/beego/orm"
)

type Value struct {
	Id         int
	Code       string
	Date       string
	Tclose     float64
	High       float64
	Low        float64
	Topen      float64
	Lclose     float64
	Chg        float64
	Pchg       float64
	Turnover   float64
	Voturnover int64
	Vaturnover int64
	Tcap       int64
	Mcap       int64
}

func (a *Value) TableName() string {
	return TableName("value")
}

func ValueAdd(a *Value) (int64, error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	return o1.Insert(a)
}

func ValueGet(CodeName string) ([]*Value, int64) {
	list := make([]*Value, 0)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("value")).Filter("code", CodeName)
	total, _ := query.Count()
	//query.OrderBy("-id").Limit(10).All(&list)
	query.OrderBy("-id").All(&list)
	return list, total
}

//func CodeGetByName(CodeName string) (*Code, error) {
//	a := new(Code)
//	o1 := orm.NewOrm()
//	o1.Using("hcs")
//	err := o1.QueryTable(TableName("set_code")).Filter("code", CodeName).One(a)
//	if err != nil {
//		return nil, err
//	}
//	return a, nil
//}

//func CodeGetList(page, pageSize int, filters ...interface{}) ([]*Code, int64) {
//	offset := (page - 1) * pageSize
//	list := make([]*Code, 0)
//	o1 := orm.NewOrm()
//	o1.Using("hcs")
//	query := o1.QueryTable(TableName("set_code"))
//	if len(filters) > 0 {
//		l := len(filters)
//		for k := 0; k < l; k += 2 {
//			query = query.Filter(filters[k].(string), filters[k+1])
//		}
//	}
//	total, _ := query.Count()
//	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

//	return list, total
//}

//func CodeGetById(id int) (*Code, error) {
//	r := new(Code)
//	o1 := orm.NewOrm()
//	o1.Using("hcs")
//	err := o1.QueryTable(TableName("set_code")).Filter("id", id).One(r)
//	if err != nil {
//		return nil, err
//	}
//	return r, nil
//}

//func (a *Code) Update(fields ...string) error {
//	o1 := orm.NewOrm()
//	o1.Using("hcs")
//	if _, err := o1.Update(a, fields...); err != nil {
//		return err
//	}
//	return nil
//}
