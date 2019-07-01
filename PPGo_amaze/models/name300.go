package models

import (
	//"fmt"
	//"os"

	"github.com/astaxie/beego/orm"
)

type Name300 struct {
	StockName
}

func (a *Name300) TableName() string {
	return TableName("name_300")
}

func Name300Add(a *Name300) (int64, error) {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	return o1.Insert(a)
}

func Name300GetList() ([]*Name300, int64) {
	list := make([]*Name300, 0)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	query := o1.QueryTable(TableName("name_300"))
	total, _ := query.Count()
	//fmt.Fprintf(os.Stdout, "total[%d]\n", total)

	query.OrderBy("-id").All(&list)
	//for id := 0; id < len(list); id++ {
	//	fmt.Fprintf(os.Stdout, "ID[%d], [%s]\n", id, list[id])
	//}

	return list, total
}

func Name300GetById(id int) (*Name300, error) {
	name := new(Name300)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	err := o1.QueryTable(TableName("name_300")).Filter("id", id).One(name)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func Name300GetByCode(code string) (*Name300, error) {
	name := new(Name300)
	o1 := orm.NewOrm()
	o1.Using("hcs")
	err := o1.QueryTable(TableName("name_300")).Filter("code", code).One(name)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func (a *Name300) Update(fields ...string) error {
	o1 := orm.NewOrm()
	o1.Using("hcs")
	if _, err := o1.Update(a, fields...); err != nil {
		return err
	}
	return nil
}
