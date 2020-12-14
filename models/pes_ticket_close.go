package models

type PesTicketClose struct {
	TicketNumber              string `orm:"column(ticket_number);size(50)"`
	TicketDepartmentId        int    `orm:"column(ticket_department_id);null"`
	TicketCategoryId          int    `orm:"column(ticket_category_id);null" description:"工单所属业务ID"`
	TicketRelatedTicketNumber string `orm:"column(ticket_related_ticket_number);size(255);null"`
	TicketRequire             string `orm:"column(ticket_require);size(255);null" description:"工单归属需求"`
	TicketSolveMember         string `orm:"column(ticket_solve_member);size(20);null" description:"工单实际解决人"`
}
