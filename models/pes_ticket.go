package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PesTicket struct {
	Id                   int    `orm:"column(ticket_id);auto"`
	TicketNumber         string `orm:"column(ticket_number);size(128)" description:"工单序号"`
	TicketTitle          string `orm:"column(ticket_title);size(255)" description:"工单标题"`
	TicketModelId        int    `orm:"column(ticket_model_id)" description:"对应的工单模型"`
	TicketStatus         int8   `orm:"column(ticket_status)" description:"工单状态,详情参考option中customstatus"`
	TicketSubmitTime     int    `orm:"column(ticket_submit_time)" description:"工单提交时间"`
	TicketReferTime      int    `orm:"column(ticket_refer_time)" description:"工单耗时参照时间"`
	TicketRunTime        int    `orm:"column(ticket_run_time)" description:"工单解决时长"`
	TicketCompleteTime   int    `orm:"column(ticket_complete_time)" description:"工单完成时间"`
	TicketRead           int8   `orm:"column(ticket_read)" description:"0:未读 1:已读"`
	UserId               int    `orm:"column(user_id);null" description:"工单操作者ID"`
	MemberId             int    `orm:"column(member_id)" description:"站内会员ID . -1表示匿名提交"`
	UserName             string `orm:"column(user_name);size(128)" description:"工单操作者名字"`
	TicketContact        int8   `orm:"column(ticket_contact)" description:"联系方式 1:邮箱 2:手机号码"`
	TicketContactAccount string `orm:"column(ticket_contact_account);size(128)" description:"联系账号"`
	TicketClose          int8   `orm:"column(ticket_close)" description:"0:正常 1:关闭"`
}

func (t *PesTicket) TableName() string {
	return "pes_ticket"
}

func init() {
	orm.RegisterModel(new(PesTicket))
}

// AddPesTicket insert a new PesTicket into database and returns
// last inserted Id on success.
func AddPesTicket(m *PesTicket) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPesTicketById retrieves PesTicket by Id. Returns error if
// Id doesn't exist
func GetPesTicketById(id int) (v *PesTicket, err error) {
	o := orm.NewOrm()
	v = &PesTicket{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPesTicket retrieves all PesTicket matches certain condition. Returns empty list if
// no records exist
func GetAllPesTicket(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PesTicket))
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

	var l []PesTicket
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

// UpdatePesTicket updates PesTicket by Id and returns error if
// the record to be updated doesn't exist
func UpdatePesTicketById(m *PesTicket) (err error) {
	o := orm.NewOrm()
	v := PesTicket{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePesTicket deletes PesTicket by Id and returns error if
// the record to be deleted doesn't exist
func DeletePesTicket(id int) (err error) {
	o := orm.NewOrm()
	v := PesTicket{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PesTicket{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
