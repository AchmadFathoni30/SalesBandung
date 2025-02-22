package model

type SalesVOID struct {
	Id          string `json:"id" gorm:"primaryKey;column:id"`
	Ent_date    string `json:"ent_date" gorm:"column:ent_date"`
	Ent_user    string `json:"ent_user" gorm:"column:ent_user"`
	Reason      int    `json:"reason" gorm:"column:reason"`
	Description string `json:"description" gorm:"column:description"`
	Branch_id   string `json:"branch_id" gorm:"column:branch_id"`
	Trx_date    string `json:"trx_date" gorm:"column:trx_date"`
}

func (SalesVOID) TableName() string {
	return "sales_void_bdg"
}
