package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type Admin struct {
	Id         int
	LoginName  string
	RealName   string
	Password   string
	RoleIds    string
	Phone      string
	Email      string
	Salt       string
	LastLogin  int64
	LastIp     string
	Status     int
	CreateId   int
	UpdateId   int
	CreateTime int64
	UpdateTime int64
}

func (a *Admin) TableName() string {
	return TableName("uc_admin")
}

func AdminAdd(a *Admin) (int64, error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	return o1.Insert(a)
}

func AdminGetByName(loginName string) (*Admin, error) {

	fmt.Println("AdminGetByName")
	a := new(Admin)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	fmt.Println("use hcs")
	err := o1.QueryTable(TableName("uc_admin")).Filter("login_name", loginName).One(a)
	if err != nil {
		fmt.Println("AdminGetByName Err `%v` ", err)
		return nil, err
	}
	return a, nil
}

func AdminGetList(page, pageSize int, filters ...interface{}) ([]*Admin, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Admin, 0)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("uc_admin"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)

	return list, total
}

func AdminGetById(id int) (*Admin, error) {
	r := new(Admin)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	err := o1.QueryTable(TableName("uc_admin")).Filter("id", id).One(r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (a *Admin) Update(fields ...string) error {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	if _, err := o1.Update(a, fields...); err != nil {
		return err
	}
	return nil
}

// func RoleAuthDelete(id int) (int64, error) {
// 	query := orm.NewOrm().QueryTable(TableName("role_auth"))
// 	return query.Filter("role_id", id).Delete()
// }

// func RoleAuthMultiAdd(ras []*RoleAuth) (n int, err error) {
// 	query := orm.NewOrm().QueryTable(TableName("role_auth"))
// 	i, _ := query.PrepareInsert()
// 	for _, ra := range ras {
// 		_, err := i.Insert(ra)
// 		if err == nil {
// 			n = n + 1
// 		}
// 	}
// 	i.Close() // 别忘记关闭 statement
// 	return n, err
// }
