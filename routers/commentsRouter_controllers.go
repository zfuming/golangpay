package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["pmt/controllers:PaymentChargesController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentChargesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentCorporationController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayOrderController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeAlipayController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeUnipayController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"] = append(beego.GlobalControllerRouter["pmt/controllers:PaymentPayTypeWxpayController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
