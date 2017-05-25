package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type PaymentPayOrder struct {
	Id              int    `orm:"column(pay_order_id);auto"`
	CorporationCode string `orm:"column(corporation_code);size(32)"`
	PayTypeCode     string `orm:"column(pay_type_code);size(32)"`
	PayTypeStyle    string `orm:"column(pay_type_style);size(32)"`
	PayCurrency     string `orm:"column(pay_currency);size(32)"`
	TotalAmount     string `orm:"column(total_amount);size(32)"`
	PayAmount       string `orm:"column(pay_amount);size(32)"`
	TradeGoodsList  string `orm:"column(trade_goods_list);size(255)"`
	DiscountAmount  string `orm:"column(discount_amount);size(32)"`
	DiscountList    string `orm:"column(discount_list);size(255)"`
	PayBuyerUser    string `orm:"column(pay_buyer_user);size(32)"`
	PaySaleUser     string `orm:"column(pay_sale_user);size(32)"`
	PayState        int    `orm:"column(pay_state)"`
	MerchantNo      string `orm:"column(merchant_no);size(32)"`
	TpTradeNo       string `orm:"column(tp_trade_no);size(32)"`
	TpRefundNo      string `orm:"column(tp_refund_no);size(32)"`
	RefundReason    string `orm:"column(refund_reason);size(255)"`
	CreateTime      string `orm:"column(create_time);size(32)"`
	PayedTime       string `orm:"column(payed_time);size(32)"`
	RefundTime      string `orm:"column(refund_time);size(32)"`
}

func (t *PaymentPayOrder) TableName() string {
	return "payment_pay_order"
}

func init() {
	orm.RegisterModel(new(PaymentPayOrder))
}

// AddPaymentPayOrder insert a new PaymentPayOrder into database and returns
// last inserted Id on success.
func AddPaymentPayOrder(m *PaymentPayOrder) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPaymentPayOrderById retrieves PaymentPayOrder by Id. Returns error if
// Id doesn't exist
func GetPaymentPayOrderById(id int) (v *PaymentPayOrder, err error) {
	o := orm.NewOrm()
	v = &PaymentPayOrder{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPaymentPayOrder retrieves all PaymentPayOrder matches certain condition. Returns empty list if
// no records exist
func GetAllPaymentPayOrder(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PaymentPayOrder))
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

	var l []PaymentPayOrder
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

// UpdatePaymentPayOrder updates PaymentPayOrder by Id and returns error if
// the record to be updated doesn't exist
func UpdatePaymentPayOrderById(m *PaymentPayOrder) (err error) {
	o := orm.NewOrm()
	v := PaymentPayOrder{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePaymentPayOrder deletes PaymentPayOrder by Id and returns error if
// the record to be deleted doesn't exist
func DeletePaymentPayOrder(id int) (err error) {
	o := orm.NewOrm()
	v := PaymentPayOrder{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&PaymentPayOrder{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
