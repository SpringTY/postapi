package main

import (
	"postapi/handler"
	"postapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LogerMiddleware())
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// model manage 相关
	r.GET("/model/states", handler.GetModelStatesHandler)
	// remove model 卸载模型
	r.DELETE("/model/:modelName/:modelVersion", handler.RemoveModelHandler)
	// update model 更新模型
	r.POST("/model/:modelName/:modelVersion", handler.UpdateModelHandler)
	// 预测模型
	r.PUT("/model/:modelName/:modelVersion", handler.PredictHandler)
	// 上传数据
	r.PUT("/data/collect", handler.UploadDealHanlder)
	// 获取预测顺序
	r.POST("/data/predict", handler.GetPredictDealHandler)
	// 获取原始数据
	r.POST("/data/raw/", handler.GetRawDealHandler)
	// 获取原始数据目录
	r.GET("/data/raw/tree/", handler.GetRawDataTreeHandler)

	r.GET("/task/gen/:taskId", handler.GetGenTaskInfoHandler)

	r.POST("/task/gen/:tag", handler.GenPostPredictDataHandler)
	r.Run("0.0.0.0:7675") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
