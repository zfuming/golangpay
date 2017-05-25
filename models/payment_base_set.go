package models

type PaymentBaseSet struct {
	SetSignUse      int    `orm:"column(set_sign_use)"`
	SetSignKey      string `orm:"column(set_sign_key);size(255)"`
	SetPayState     string `orm:"column(set_pay_state);size(255)"`
	SetCurrency     string `orm:"column(set_currency);size(255)"`
	SetCurrencyInfo string `orm:"column(set_currency_info);size(255)"`
}
