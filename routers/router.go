// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"testpes/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/pes_backend_category",
			beego.NSInclude(
				&controllers.PesBackendCategoryController{},
			),
		),

		beego.NSNamespace("/pes_category",
			beego.NSInclude(
				&controllers.PesCategoryController{},
			),
		),

		beego.NSNamespace("/pes_department",
			beego.NSInclude(
				&controllers.PesDepartmentController{},
			),
		),

		beego.NSNamespace("/pes_field",
			beego.NSInclude(
				&controllers.PesFieldController{},
			),
		),

		beego.NSNamespace("/pes_findpassword",
			beego.NSInclude(
				&controllers.PesFindpasswordController{},
			),
		),

		beego.NSNamespace("/pes_mail_template",
			beego.NSInclude(
				&controllers.PesMailTemplateController{},
			),
		),

		beego.NSNamespace("/pes_member",
			beego.NSInclude(
				&controllers.PesMemberController{},
			),
		),

		beego.NSNamespace("/pes_menu",
			beego.NSInclude(
				&controllers.PesMenuController{},
			),
		),

		beego.NSNamespace("/pes_model",
			beego.NSInclude(
				&controllers.PesModelController{},
			),
		),

		beego.NSNamespace("/pes_node",
			beego.NSInclude(
				&controllers.PesNodeController{},
			),
		),

		beego.NSNamespace("/pes_node_group",
			beego.NSInclude(
				&controllers.PesNodeGroupController{},
			),
		),

		beego.NSNamespace("/pes_option",
			beego.NSInclude(
				&controllers.PesOptionController{},
			),
		),

		beego.NSNamespace("/pes_route",
			beego.NSInclude(
				&controllers.PesRouteController{},
			),
		),

		beego.NSNamespace("/pes_send",
			beego.NSInclude(
				&controllers.PesSendController{},
			),
		),

		beego.NSNamespace("/pes_ticket",
			beego.NSInclude(
				&controllers.PesTicketController{},
			),
		),

		beego.NSNamespace("/pes_ticket_chat",
			beego.NSInclude(
				&controllers.PesTicketChatController{},
			),
		),

		beego.NSNamespace("/pes_ticket_content",
			beego.NSInclude(
				&controllers.PesTicketContentController{},
			),
		),

		beego.NSNamespace("/pes_ticket_form",
			beego.NSInclude(
				&controllers.PesTicketFormController{},
			),
		),

		beego.NSNamespace("/pes_ticket_model",
			beego.NSInclude(
				&controllers.PesTicketModelController{},
			),
		),

		beego.NSNamespace("/pes_user",
			beego.NSInclude(
				&controllers.PesUserController{},
			),
		),

		beego.NSNamespace("/pes_user_group",
			beego.NSInclude(
				&controllers.PesUserGroupController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//自定义路由
	beego.Router("/myrt/:id",&controllers.TestController{},"GET:Get")
	//没有指定方法的路由
	//beego.Router("/myrt/:id",&controllers.TestController{})
	beego.Router("/upload/*.*",&controllers.UploadController{})

}
