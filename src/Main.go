package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//自定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("before middleware")
	}
}

func main() {

	//gin.Default() 默认开启 logger 和 recovery 两个中间件，从源码中可以看到，相当于是调用 New 函数之后使用 Use 开启两个中间件
	r := gin.Default()
	//分组
	routerGroup := r.Group("v1")

	//使用中间件
	routerGroup.Use(MiddleWare())
	{
		routerGroup.GET("/midd", func(c *gin.Context) {

			request := c.MustGet("request").(string)

			c.JSON(http.StatusOK, gin.H{
				"middile_request": request,
			})
		})

	}

	r.GET("/test", func(c *gin.Context) {
		// 重定向
		// c.Redirect(http.StatusMovedPermanently,"https://baidu.com")

		param := c.Query("name")
		sex := c.PostForm("sex")

		data := map[string]interface{}{
			"lang":  "go lang",
			"param": param,
			"sex":   sex,
		}

		//响应参数组装  json  xml  string
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   data,
		})

	})

	r.Run(":8000")
}
