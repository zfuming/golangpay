package routers

import (
	"pmt/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	// api接口
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/payment_charges",
			beego.NSInclude(
				&controllers.PaymentChargesController{},
			),
		),

		beego.NSNamespace("/payment_inquiry",
			beego.NSInclude(
				&controllers.PaymentInquiryController{},
			),
		),

		beego.NSNamespace("/payment_refund",
			beego.NSInclude(
				&controllers.PaymentRefundController{},
			),
		),

	)
	beego.AddNamespace(ns)
}
