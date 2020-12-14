package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PesTicketForm struct {
	Id                    int    `orm:"column(ticket_form_id);auto"`
	TicketFormModelId     int    `orm:"column(ticket_form_model_id)" description:"对应的工单模型ID"`
	TicketFormName        string `orm:"column(ticket_form_name);size(128)" description:"工单表单名词"`
	TicketFormDescription string `orm:"column(ticket_form_description);size(128)" description:"工单表单显示名称"`
	TicketFormExplain     string `orm:"column(ticket_form_explain);size(128)" description:"工单表单说明"`
	TicketFormMsg         string `orm:"column(ticket_form_msg);size(128)" description:"提示信息"`
	TicketFormType        string `orm:"column(ticket_form_type);size(16)" description:"工单表单类型"`
	TicketFormOption      string `orm:"column(ticket_form_option)" description:"工单表单的选项值"`
	TicketFormVerify      string `orm:"column(ticket_form_verify);size(32)" description:"工单表单的验证类型"`
	TicketFormRequired    int8   `orm:"column(ticket_form_required)" description:"是否必填 0: 否 1:必填"`
	TicketFormStatus      int8   `orm:"column(ticket_form_status)" description:"是否启用 0:否 1:启用"`
	TicketFormListsort    int    `orm:"column(ticket_form_listsort)" description:"动态表单的排序值（升值））"`
	TicketFormBind        int    `orm:"column(ticket_form_bind)" description:"绑定的联动表单"`
	TicketFormBindValue   string `orm:"column(ticket_form_bind_value);size(255)" description:"联动触发值"`
}

func (t *PesTicketForm) TableName() string {
	return "pes_ticket_form"
}

func init() {
	orm.RegisterModel(new(PesTicketForm))
}

// AddPesTicketForm insert a new PesTicketForm into database and returns
// last inserted Id on success.
func AddPesTicketForm(m *PesTicketForm) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPesTicketFormById retrieves PesTicketForm by Id. Returns error if
// Id doesn't exist
func GetPesTicketFormById(id int) (v *PesTicketForm, err error) {
	o := orm.NewOrm()
	v = &PesTicketForm{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPesTicketForm retrieves all PesTicketForm matches certain condition. Returns empty list if
// no records exist
func GetAllPesTicketForm(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PesTicketForm))
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

	var l []PesTicketForm
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

// UpdatePesTicketForm updates PesTicketForm by Id and returns error if
// the record to be updated doesn't exist
func UpdatePesTicketFormById(m *PesTicketForm) (err error) {
	o := orm.NewOrm()
	v := PesTicketForm{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePesTicketForm deletes PesTicketForm by Id and returns error if
// the record to be deleted doesn't exist
func DeletePesTicketForm(id int) (err error) {
	o := orm.NewOrm()
	v := PesTicketForm{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PesTicketForm{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
