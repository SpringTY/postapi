package handler

import (
	"net/http"
	view_model "postapi/model/view"
	"postapi/rpc"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetDealsHandler(c *gin.Context) {
	// define io
	dataSource := c.Param("dataSource")
	dealId := c.Param("dealId")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageNum := c.DefaultPostForm("pageNum", "01")
	resp := &view_model.CommonResult{
		Message: "success",
		Status:  0,
		Data:    nil,
	}
	// call handler
	GetDeals(dataSource, dealId, pageSize, pageNum, resp)
	// give resp
	c.JSON(http.StatusOK, resp)
}

func GetDeals(dataSource, dealId, pageSize, pageNum string, resp *view_model.CommonResult) {
	data := view_model.GetDealsData{}
	if dataSource == "static" {
		data.Deals, _ = rpc.GetStaticDealData()
		resp.Data = data
		logrus.WithFields(logrus.Fields{
			"data": data,
		}).Info()
		return
	}
	logrus.Warn("Error")
}
