package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PesField struct {
	Id               int    `orm:"column(field_id);auto"`
	FieldModelId     int    `orm:"column(field_model_id)"`
	FieldName        string `orm:"column(field_name);size(128)"`
	FieldDisplayName string `orm:"column(field_display_name);size(128)"`
	FieldType        string `orm:"column(field_type);size(128)"`
	FieldOption      string `orm:"column(field_option)"`
	FieldExplain     string `orm:"column(field_explain);size(128)"`
	FieldDefault     string `orm:"column(field_default);size(128)"`
	FieldRequired    int8   `orm:"column(field_required)"`
	FieldListsort    int    `orm:"column(field_listsort)"`
	FieldList        int8   `orm:"column(field_list)" description:"是否显示于列表"`
	FieldForm        int8   `orm:"column(field_form)" description:"是否显示于表单 0:否 1:显示"`
	FieldStatus      int8   `orm:"column(field_status)"`
}

func (t *PesField) TableName() string {
	return "pes_field"
}

func init() {
	orm.RegisterModel(new(PesField))
}

// AddPesField insert a new PesField into database and returns
// last inserted Id on success.
func AddPesField(m *PesField) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPesFieldById retrieves PesField by Id. Returns error if
// Id doesn't exist
func GetPesFieldById(id int) (v *PesField, err error) {
	o := orm.NewOrm()
	v = &PesField{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPesField retrieves all PesField matches certain condition. Returns empty list if
// no records exist
func GetAllPesField(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PesField))
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

	var l []PesField
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

// UpdatePesField updates PesField by Id and returns error if
// the record to be updated doesn't exist
func UpdatePesFieldById(m *PesField) (err error) {
	o := orm.NewOrm()
	v := PesField{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePesField deletes PesField by Id and returns error if
// the record to be deleted doesn't exist
func DeletePesField(id int) (err error) {
	o := orm.NewOrm()
	v := PesField{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PesField{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
