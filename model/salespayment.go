package model

type SalesPAYMENT struct {
	Id               string  `json:"id" gorm:"primaryKey;column:id"`
	Gross_sales      string  `json:"gross_sales" gorm:"column:gross_sales"`
	Gross            string  `json:"gross" gorm:"column:gross"`
	Sales_discount   string  `json:"sales_discount" gorm:"column:sales_discount"`
	Total_cost       string  `json:"total_cost" gorm:"column:total_cost"`
	Tax              string  `json:"tax" gorm:"column:tax"`
	Sub_total        string  `json:"sub_total" gorm:"column:sub_total"`
	Net              string  `json:"net" gorm:"column:net"`
	Cash             string  `json:"cash" gorm:"column:cash"`
	Non_cash         string  `json:"non_cash" gorm:"column:non_cash"`
	Rounding         string  `json:"rounding" gorm:"column:rounding"`
	Refund           string  `json:"refund" gorm:"column:refund"`
	Refund_vch       string  `json:"refund_vch" gorm:"column:refund_vch"`
	Item_discount    string  `json:"item_discount" gorm:"column:item_discount"`
	Id_dp            *string `json:"id_dp" gorm:"column:id_dp"`
	Dp               string  `json:"dp" gorm:"column:dp"`
	Net_dp           string  `json:"net_dp" gorm:"column:net_dp"`
	Fbca             string  `json:"fbca" gorm:"column:fbca"`
	Card_amt         string  `json:"card_amt" gorm:"column:card_amt"`
	Card_no          *string `json:"card_no" gorm:"column:card_no"`
	Edc_vdr          *string `json:"edc_vdr" gorm:"column:edc_vdr"`
	Promo_discount   string  `json:"promo_discount" gorm:"column:promo_discount"`
	Note             *string `json:"note" gorm:"column:note"`
	Promoid          *string `json:"promoid" gorm:"column:promoid"`
	Promodesc        *string `json:"promodesc" gorm:"column:promodesc"`
	Discreason       *string `json:"discreason" gorm:"column:discreason"`
	Disctype         *string `json:"disctype" gorm:"column:disctype"`
	Discvalue        *string `json:"discvalue" gorm:"column:discvalue"`
	Item_type        *string `json:"item_type" gorm:"column:item_type"`
	Discmax          *string `json:"discmax" gorm:"column:discmax"`
	Gross_sales_disc *string `json:"gross_sales_disc" gorm:"column:gross_sales_disc"`
	Transrc          *string `json:"transrc" gorm:"column:transrc"`
	Tacharge         string  `json:"tacharge" gorm:"column:tacharge"`
	Flag_omni        *string `json:"flag_omni" gorm:"column:flag_omni"`
	Branch_id        string  `json:"branch_id" gorm:"column:branch_id"`
	Trx_date         string  `json:"trx_date" gorm:"column:trx_date"`
}

func (SalesPAYMENT) TableName() string {
	return "sales_payment_bdg"
}
