package models

import (
	//"fmt"
	//"os"

	"github.com/astaxie/beego/orm"
)

type Name000 struct {
	StockName
}

func (a *Name000) TableName() string {
	return TableName("name_000")
}

func Name000Add(a *Name000) (int64, error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	return o1.Insert(a)
}

func Name000GetList() ([]*Name000, int64) {
	list := make([]*Name000, 0)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("name_000"))
	total, _ := query.Count()
	//fmt.Fprintf(os.Stdout, "total[%d]\n", total)

	query.OrderBy("-id").All(&list)
	//for id := 0; id < len(list); id++ {
	//	fmt.Fprintf(os.Stdout, "ID[%d], [%s]\n", id, list[id])
	//}

	return list, total
}

func Name000GetById(id int) (*Name000, error) {
	name := new(Name000)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	err := o1.QueryTable(TableName("name_000")).Filter("id", id).One(name)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func Name000GetByCode(code string) (*Name000, error) {
	name := new(Name000)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	err := o1.QueryTable(TableName("name_000")).Filter("code", code).One(name)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func (a *Name000) Update(fields ...string) error {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	if _, err := o1.Update(a, fields...); err != nil {
		return err
	}
	return nil
}
