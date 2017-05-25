package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PaymentPayTypeAlipay struct {
	Id               int    `orm:"column(pay_type_id);auto"`
	PayTypeTitle     string `orm:"column(pay_type_title);size(255)"`
	AlipayGatewayUrl string `orm:"column(alipay_gateway_url);size(255)"`
	AlipayAppId      string `orm:"column(alipay_app_id);size(255)"`
	AlipayRsaKey     string `orm:"column(alipay_rsa_key);size(255)"`
	AlipayAllowStyle string `orm:"column(alipay_allow_style);size(255)"`
	AlipayFormatType string `orm:"column(alipay_format_type);size(32)"`
	AlipayCharset    string `orm:"column(alipay_charset);size(32)"`
	AlipaySignType   string `orm:"column(alipay_sign_type);size(32)"`
	CreateTime       string `orm:"column(create_time);size(32)"`
	LastEditTime     string `orm:"column(last_edit_time);size(32)"`
}

func (t *PaymentPayTypeAlipay) TableName() string {
	return "payment_pay_type_alipay"
}

func init() {
	orm.RegisterModel(new(PaymentPayTypeAlipay))
}

// AddPaymentPayTypeAlipay insert a new PaymentPayTypeAlipay into database and returns
// last inserted Id on success.
func AddPaymentPayTypeAlipay(m *PaymentPayTypeAlipay) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPaymentPayTypeAlipayById retrieves PaymentPayTypeAlipay by Id. Returns error if
// Id doesn't exist
func GetPaymentPayTypeAlipayById(id int) (v *PaymentPayTypeAlipay, err error) {
	o := orm.NewOrm()
	v = &PaymentPayTypeAlipay{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPaymentPayTypeAlipay retrieves all PaymentPayTypeAlipay matches certain condition. Returns empty list if
// no records exist
func GetAllPaymentPayTypeAlipay(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PaymentPayTypeAlipay))
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

	var l []PaymentPayTypeAlipay
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

// UpdatePaymentPayTypeAlipay updates PaymentPayTypeAlipay by Id and returns error if
// the record to be updated doesn't exist
func UpdatePaymentPayTypeAlipayById(m *PaymentPayTypeAlipay) (err error) {
	o := orm.NewOrm()
	v := PaymentPayTypeAlipay{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePaymentPayTypeAlipay deletes PaymentPayTypeAlipay by Id and returns error if
// the record to be deleted doesn't exist
func DeletePaymentPayTypeAlipay(id int) (err error) {
	o := orm.NewOrm()
	v := PaymentPayTypeAlipay{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PaymentPayTypeAlipay{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
