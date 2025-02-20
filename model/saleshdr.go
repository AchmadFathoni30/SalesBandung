package model

import "gorm.io/gorm"

type SalesHDR struct {
	gorm.Model
	Id           string `json:"id"`
	Reset_no     int    `json:"reset_no"`
	Counter_no   int    `json:"counter_no"`
	Trx_id       string `json:"trx_id"`
	Trx_date     string `json:"trx_date"`
	Branch_id    string `json:"branch_id"`
	Client_id    string `json:"client_id"`
	Session_id   int    `json:"session_id"`
	Pos_no       string `json:"pos_no"`
	Pos_id       string `json:"pos_id"`
	Order_no     string `json:"order_no,omitempty"`
	Ent_date     string `json:"ent_date"`
	Ent_user     string `json:"ent_user"`
	Table_no     string `json:"table_no,omitempty"`
	Tax_rate     string `json:"tax_rate"`
	Svc_charge   string `json:"svc_charge"`
	Sdisc_rate   string `json:"sdisc_rate"`
	Guest        string `json:"guest"`
	Status       string `json:"status"`
	Valid_status string `json:"valid_status"`
	Valid_date   string `json:"valid_date"`
	Prn_rev_no   string `json:"prn_rev_no"`
	Refund_vch   string `json:"refund_vch"`
	Delivery_no  string `json:"delivery_no,omitempty"`
	Dp_status    string `json:"dp_status,omitempty"`
	Kd_pemakai   string `json:"kd_pemakai"`
	Href         string `json:"href"`
	Trx_status   string `json:"trx_status,omitempty"`
	Flag_omni    string `json:"flag_omni,omitempty"`
}
