package main

import (
	"github.com/gin-gonic/gin"
	"goweb-demo/dao"
	"strconv"
)

/**
 * 使用 go mod vendor 可以将当前项目中所有使用到的到依赖移到vendor目录中.
 * @author ZhenXinMa
 * @slogan 浮生若梦,若梦非梦,浮生何梦? 如梦之梦.
 * @create 2020-09-21 22:27 周一
 **/
func main() {

	// 使用gin创建一个默认的路由.
	r := gin.Default()
	// GET：请求方式；/hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	r.GET("/json", getQuery)
	r.GET("/add", addToDb)
	r.GET("/get", queryById)
	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run()

	// 解决跨域 没有请求过,不知道.

	/**
	w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json") //返回数据格式是json
	*/

}

func queryById(c *gin.Context) {

	// 从url中获取的值是string类型.
	id := c.Query("id")
	// 将字符串转换为数值类型
	idInt,_ := strconv.ParseInt(id, 10, 64)
	// 查询数据.
	result := dao.Query(int(idInt))
	// 现在要将该结果显示到页面上.
	c.JSON(200,result)

}

func addToDb(c *gin.Context) {

	dao.Add()

}

func getQuery(c *gin.Context) {

	param := c.Query("query")

	// 直接以c返回.
	c.JSON(200, gin.H{
		"queryParam": param,
	})

}

func postQuery(c *gin.Context) {

	// 获取表单中的数据.
	// 这个key应该是HTML中的name属性.
	c.PostForm("username")

}

func paramQuery(c *gin.Context) {

	// 获取路径上的参数进行显示.

}
