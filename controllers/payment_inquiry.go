package controllers

import (
	"github.com/astaxie/beego"
	"pmt/models"
	"encoding/json"
)

// PaymentInquiryController operations for PaymentInquiry
type PaymentInquiryController struct {
	beego.Controller
}

// URLMapping ...
func (c *PaymentInquiryController) URLMapping() {
	c.Mapping("Post", c.Post)
}

func (c *PaymentInquiryController) Post() {
	var v models.PaymentPayOrder
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddPaymentPayOrder(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}

	type Student struct {
		Name    string
		Age     int
	}
	st := &Student {
		"Xiao Ming",
		16,
	}
	c.Data["json"] = st
	c.ServeJSON()
}