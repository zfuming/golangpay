package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PaymentPayTypeWxpay struct {
	Id              int    `orm:"column(pay_type_id);auto"`
	PayTypeTitle    string `orm:"column(pay_type_title);size(255)"`
	WxpayGatewayUrl string `orm:"column(wxpay_gateway_url);size(255)"`
	WxpayAppId      string `orm:"column(wxpay_app_id);size(255)"`
	WxpayMchId      string `orm:"column(wxpay_mch_id);size(255)"`
	WxpayKey        string `orm:"column(wxpay_key);size(255)"`
	WxpayAppSecret  string `orm:"column(wxpay_app_secret);size(255)"`
	WxpayAllowStyle string `orm:"column(wxpay_allow_style);size(255)"`
	WxpayClientCert string `orm:"column(wxpay_client_cert);size(255)"`
	WxpayClientKey  string `orm:"column(wxpay_client_key);size(255)"`
	CreateTime      string `orm:"column(create_time);size(32)"`
	LastEditTime    string `orm:"column(last_edit_time);size(32)"`
}

func (t *PaymentPayTypeWxpay) TableName() string {
	return "payment_pay_type_wxpay"
}

func init() {
	orm.RegisterModel(new(PaymentPayTypeWxpay))
}

// AddPaymentPayTypeWxpay insert a new PaymentPayTypeWxpay into database and returns
// last inserted Id on success.
func AddPaymentPayTypeWxpay(m *PaymentPayTypeWxpay) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPaymentPayTypeWxpayById retrieves PaymentPayTypeWxpay by Id. Returns error if
// Id doesn't exist
func GetPaymentPayTypeWxpayById(id int) (v *PaymentPayTypeWxpay, err error) {
	o := orm.NewOrm()
	v = &PaymentPayTypeWxpay{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPaymentPayTypeWxpay retrieves all PaymentPayTypeWxpay matches certain condition. Returns empty list if
// no records exist
func GetAllPaymentPayTypeWxpay(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PaymentPayTypeWxpay))
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

	var l []PaymentPayTypeWxpay
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

// UpdatePaymentPayTypeWxpay updates PaymentPayTypeWxpay by Id and returns error if
// the record to be updated doesn't exist
func UpdatePaymentPayTypeWxpayById(m *PaymentPayTypeWxpay) (err error) {
	o := orm.NewOrm()
	v := PaymentPayTypeWxpay{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePaymentPayTypeWxpay deletes PaymentPayTypeWxpay by Id and returns error if
// the record to be deleted doesn't exist
func DeletePaymentPayTypeWxpay(id int) (err error) {
	o := orm.NewOrm()
	v := PaymentPayTypeWxpay{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PaymentPayTypeWxpay{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
