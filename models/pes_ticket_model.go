package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PesTicketModel struct {
	Id                  int    `orm:"column(ticket_model_id);auto"`
	TicketModelNumber   string `orm:"column(ticket_model_number);size(32);null" description:"每个用户看到的唯一工单模型ID"`
	TicketModelName     string `orm:"column(ticket_model_name);size(128)" description:"工单模型名称"`
	TicketModelStatus   int8   `orm:"column(ticket_model_status)" description:"工单模型是否启用"`
	TicketModelLogin    int    `orm:"column(ticket_model_login)"`
	TicketModelVerify   int    `orm:"column(ticket_model_verify)"`
	TicketModelCid      int    `orm:"column(ticket_model_cid)"`
	TicketModelListsort int    `orm:"column(ticket_model_listsort)"`
	TicketModelExplain  string `orm:"column(ticket_model_explain)"`
	TicketModelOwner    string `orm:"column(ticket_model_owner);size(20);null" description:"工单负责人"`
}

func (t *PesTicketModel) TableName() string {
	return "pes_ticket_model"
}

func init() {
	orm.RegisterModel(new(PesTicketModel))
}

// AddPesTicketModel insert a new PesTicketModel into database and returns
// last inserted Id on success.
func AddPesTicketModel(m *PesTicketModel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPesTicketModelById retrieves PesTicketModel by Id. Returns error if
// Id doesn't exist
func GetPesTicketModelById(id int) (v *PesTicketModel, err error) {
	o := orm.NewOrm()
	v = &PesTicketModel{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPesTicketModel retrieves all PesTicketModel matches certain condition. Returns empty list if
// no records exist
func GetAllPesTicketModel(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PesTicketModel))
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

	var l []PesTicketModel
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

// UpdatePesTicketModel updates PesTicketModel by Id and returns error if
// the record to be updated doesn't exist
func UpdatePesTicketModelById(m *PesTicketModel) (err error) {
	o := orm.NewOrm()
	v := PesTicketModel{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePesTicketModel deletes PesTicketModel by Id and returns error if
// the record to be deleted doesn't exist
func DeletePesTicketModel(id int) (err error) {
	o := orm.NewOrm()
	v := PesTicketModel{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PesTicketModel{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
