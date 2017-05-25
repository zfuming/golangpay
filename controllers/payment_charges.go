package controllers

import (
	"github.com/astaxie/beego"
	"pmt/models"
	"encoding/json"
)

// PaymentPayTypeController operations for PaymentPayType
type PaymentChargesController struct {
	beego.Controller
}

// URLMapping ...
func (c *PaymentChargesController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Post
// @Description create PaymentCharges
// @Param	body		body 	models.PaymentCharges	true		"body for PaymentCharges content"
// @Success 201 {int} models.PaymentCharges
// @Failure 403 body is empty
// @router / [post]
func (c *PaymentChargesController) Post() {
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
