package model

type SalesDTL struct {
	Seq_no      string `json:"seq_no" gorm:"column:seq_no"`
	Id          string `json:"id" gorm:"column:id"`
	Prodcode    string `json:"prodcode" gorm:"column:prodcode"`
	Prodmain    string `json:"prodmain" gorm:"column:prodmain"`
	Qty         string `json:"qty" gorm:"column:qty"`
	Price       string `json:"price" gorm:"column:price"`
	Disc_item   string `json:"disc_item" gorm:"column:disc_item"`
	Amount      string `json:"amount" gorm:"column:amount"`
	Item_type   string `json:"item_type" gorm:"column:item_type"`
	Gross       string `json:"gross" gorm:"column:gross"`
	Disc_amount string `json:"disc_amount" gorm:"column:disc_amount"`
	Id_index    string `json:"id_index" gorm:"column:id_index"`
	Prodname    string `json:"prodname" gorm:"column:prodname"`
	Parentid    string `json:"parentid" gorm:"column:parentid"`
	Linedesc    string `json:"linedesc" gorm:"column:linedesc"`
	Promoid     string `json:"promoid" gorm:"column:promoid"`
	Promodesc   string `json:"promodesc" gorm:"column:promodesc"`
	Discreason  string `json:"discreason" gorm:"discreason"`
	Disctype    string `json:"disctype" gorm:"column:disctype"`
	Discvalue   string `json:"discvalue" gorm:"column:discvalue"`
	Ket         string `json:"ket" gorm:"column:ket"`
	Discyes     string `json:"discyes" gorm:"column:discyes"`
	Transrc     string `json:"transrc" gorm:"column:transrc"`
	Flag_omni   string `json:"flag_omni" gorm:"column:flag_omni"`
	Branch_id   string `json:"branch_id" gorm:"column:branch_id"`
	Trx_date    string `json:"trx_date" gorm:"trx_date"`
}

func (SalesDTL) TableName() string {
	return "sales_dtl_bdg"
}
