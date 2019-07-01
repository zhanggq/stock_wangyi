package models

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/astaxie/beego/orm"
)

type RoleAuth struct {
	AuthId int `orm:"pk"`
	RoleId int64
}

func (ra *RoleAuth) TableName() string {
	return TableName("uc_role_auth")
}

func RoleAuthAdd(ra *RoleAuth) (int64, error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	return o1.Insert(ra)
}

func RoleAuthGetById(id int) ([]*RoleAuth, error) {
	list := make([]*RoleAuth, 0)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("uc_role_auth"))
	_, err := query.Filter("role_id", id).All(&list, "AuthId")
	if err != nil {
		return nil, err
	}
	return list, nil
}

func RoleAuthDelete(id int) (int64, error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("uc_role_auth"))
	return query.Filter("role_id", id).Delete()
}

//获取多个
func RoleAuthGetByIds(RoleIds string) (Authids string, err error) {
	list := make([]*RoleAuth, 0)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("uc_role_auth"))
	ids := strings.Split(RoleIds, ",")
	_, err = query.Filter("role_id__in", ids).All(&list, "AuthId")
	if err != nil {
		return "", err
	}
	b := bytes.Buffer{}
	for _, v := range list {
		if v.AuthId != 0 && v.AuthId != 1 {
			b.WriteString(strconv.Itoa(v.AuthId))
			b.WriteString(",")
		}
	}
	Authids = strings.TrimRight(b.String(), ",")
	return Authids, nil
}

func RoleAuthMultiAdd(ras []*RoleAuth) (n int, err error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("uc_role_auth"))
	i, _ := query.PrepareInsert()
	for _, ra := range ras {
		_, err := i.Insert(ra)
		if err == nil {
			n = n + 1
		}
	}
	i.Close() // 别忘记关闭 statement
	return n, err
}
