package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Auth struct {
	Id         int
	AuthName   string
	AuthUrl    string
	UserId     int
	Pid        int
	Sort       int
	Icon       string
	IsShow     int
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *Auth) TableName() string {
	return TableName("uc_auth")
}

func AuthGetList(page, pageSize int, filters ...interface{}) ([]*Auth, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Auth, 0)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("uc_auth"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("pid", "sort").Limit(pageSize, offset).All(&list)

	return list, total
}

func AuthGetListByIds(authIds string, userId int) ([]*Auth, error) {
	list1 := make([]*Auth, 0)
	var list []orm.Params
	//list:=[]orm.Params
	var err error
	o1 := orm.NewOrm()
	o1.Using("hcs")
	if userId == 1 {
		//超级管理员
		_, err = o1.Raw("select id,auth_name,auth_url,pid,icon,is_show from pp_uc_auth where status=? order by pid asc,sort asc", 1).Values(&list)
	} else {
		_, err = o1.Raw("select id,auth_name,auth_url,pid,icon,is_show from pp_uc_auth where status=1 and id in("+authIds+") order by pid asc,sort asc", authIds).Values(&list)
	}

	for k, v := range list {
		fmt.Println(k, v)
	}

	fmt.Println(list)
	return list1, err
}

func AuthAdd(auth *Auth) (int64, error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	return o1.Insert(auth)
}

func AuthGetById(id int) (*Auth, error) {
	a := new(Auth)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	err := o1.QueryTable(TableName("uc_auth")).Filter("id", id).One(a)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Auth) Update(fields ...string) error {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	if _, err := o1.Update(a, fields...); err != nil {
		return err
	}
	return nil
}
