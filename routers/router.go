package routers

import (
	"pmt/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/payment_corporation",
			beego.NSInclude(
				&controllers.PaymentCorporationController{},
			),
		),

		beego.NSNamespace("/payment_pay_order",
			beego.NSInclude(
				&controllers.PaymentPayOrderController{},
			),
		),

		beego.NSNamespace("/payment_pay_type",
			beego.NSInclude(
				&controllers.PaymentPayTypeController{},
			),
		),

		beego.NSNamespace("/payment_pay_type_alipay",
			beego.NSInclude(
				&controllers.PaymentPayTypeAlipayController{},
			),
		),

		beego.NSNamespace("/payment_pay_type_unipay",
			beego.NSInclude(
				&controllers.PaymentPayTypeUnipayController{},
			),
		),

		beego.NSNamespace("/payment_pay_type_wxpay",
			beego.NSInclude(
				&controllers.PaymentPayTypeWxpayController{},
			),
		),

		beego.NSNamespace("/payment_charges",
			beego.NSInclude(
				&controllers.PaymentChargesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
