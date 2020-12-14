package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PesSend struct {
	Id          int    `orm:"column(send_id);auto"`
	SendAccount string `orm:"column(send_account);size(255)"`
	SendTitle   string `orm:"column(send_title);size(255)" description:"待发送标题"`
	SendContent string `orm:"column(send_content)" description:"待发送的内容"`
	SendTime    int    `orm:"column(send_time)" description:"发送时间"`
	SendType    int8   `orm:"column(send_type)" description:"1:邮箱 2:手机 .."`
}

func (t *PesSend) TableName() string {
	return "pes_send"
}

func init() {
	orm.RegisterModel(new(PesSend))
}

// AddPesSend insert a new PesSend into database and returns
// last inserted Id on success.
func AddPesSend(m *PesSend) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPesSendById retrieves PesSend by Id. Returns error if
// Id doesn't exist
func GetPesSendById(id int) (v *PesSend, err error) {
	o := orm.NewOrm()
	v = &PesSend{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPesSend retrieves all PesSend matches certain condition. Returns empty list if
// no records exist
func GetAllPesSend(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PesSend))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []PesSend
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePesSend updates PesSend by Id and returns error if
// the record to be updated doesn't exist
func UpdatePesSendById(m *PesSend) (err error) {
	o := orm.NewOrm()
	v := PesSend{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePesSend deletes PesSend by Id and returns error if
// the record to be deleted doesn't exist
func DeletePesSend(id int) (err error) {
	o := orm.NewOrm()
	v := PesSend{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PesSend{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
