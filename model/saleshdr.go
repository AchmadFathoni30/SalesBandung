package model

type SalesHDR struct {
	Id           string  `json:"id" gorm:"primaryKey;column:id"`
	Reset_no     int     `json:"reset_no" gorm:"column:reset_no"`
	Counter_no   int     `json:"counter_no" gorm:"column:counter_no"`
	Trx_id       string  `json:"trx_id" gorm:"column:trx_id"`
	Trx_date     string  `json:"trx_date" gorm:"column:trx_date"`
	Branch_id    string  `json:"branch_id" gorm:"column:branch_id"`
	Client_id    *string `json:"client_id" gorm:"column:client_id"`
	Session_id   int     `json:"session_id" gorm:"column:session_id"`
	Pos_no       *string `json:"pos_no" gorm:"column:pos_no"`
	Pos_id       string  `json:"pos_id" gorm:"column:pos_id"`
	Order_no     *string `json:"order_no" gorm:"column:order_no"`
	Ent_date     string  `json:"ent_date" gorm:"column:ent_date"`
	Ent_user     string  `json:"ent_user" gorm:"column:ent_user"`
	Table_no     *string `json:"table_no" gorm:"column:table_no"`
	Tax_rate     string  `json:"tax_rate" gorm:"column:tax_rate"`
	Svc_charge   string  `json:"svc_charge" gorm:"column:svc_charge"`
	Sdisc_rate   string  `json:"sdisc_rate" gorm:"column:sdisc_rate"`
	Guest        string  `json:"guest" gorm:"column:guest"`
	Status       string  `json:"status" gorm:"column:status"`
	Valid_status string  `json:"valid_status" gorm:"column:valid_status"`
	Valid_date   string  `json:"valid_date" gorm:"column:valid_date"`
	Prn_rev_no   string  `json:"prn_rev_no" gorm:"column:prn_rev_no"`
	Refund_vch   string  `json:"refund_vch" gorm:"column:refund_vch"`
	Delivery_no  *string `json:"delivery_no" gorm:"column:delivery_no"`
	Dp_status    *string `json:"dp_status" gorm:"column:dp_status"`
	Kd_pemakai   string  `json:"kd_pemakai" gorm:"column:kd_pemakai"`
	Href         *string `json:"href" gorm:"column:href"`
	Trx_status   *string `json:"trx_status" gorm:"column:trx_status"`
	Flag_omni    *string `json:"flag_omni" gorm:"column:flag_omni"`
}

func (SalesHDR) TableName() string {
	return "sales_hdr_bdg"
}
