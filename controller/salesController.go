package controller

import (
	"SalesBandung/config"
	"SalesBandung/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSalesHDR(c *gin.Context) {
	branch_id := c.Query("branch_id")
	trx_date := c.Query("trx_date")

	if branch_id == "" || trx_date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Branch id dan Trx date harap di isi!!!"})
		return
	}

	query := `SELECT * FROM sales_hdr_bdg WHERE branch_id=$1 AND trx_date=$2`
	rows, err := config.DB.Raw(query, branch_id, trx_date).Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var results []model.SalesHDR
	for rows.Next() {
		var sales_hdr model.SalesHDR
		if err := rows.Scan(&sales_hdr.Id, &sales_hdr.Reset_no, &sales_hdr.Counter_no, &sales_hdr.Trx_id, &sales_hdr.Trx_date,
			&sales_hdr.Branch_id, &sales_hdr.Client_id, &sales_hdr.Session_id, &sales_hdr.Pos_no, &sales_hdr.Pos_id,
			&sales_hdr.Order_no, &sales_hdr.Ent_date, &sales_hdr.Ent_user, &sales_hdr.Table_no, &sales_hdr.Tax_rate,
			&sales_hdr.Svc_charge, &sales_hdr.Sdisc_rate, &sales_hdr.Guest, &sales_hdr.Status, &sales_hdr.Valid_status,
			&sales_hdr.Valid_date, &sales_hdr.Prn_rev_no, &sales_hdr.Refund_vch, &sales_hdr.Delivery_no, &sales_hdr.Dp_status,
			&sales_hdr.Kd_pemakai, &sales_hdr.Href, &sales_hdr.Trx_status, &sales_hdr.Flag_omni); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, sales_hdr)
	}
	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, results)
}
