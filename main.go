package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Document struct {
	ID    string
	URL   string
	Title string
}

func main() {
	if err := initDB(); err != nil {
		fmt.Println("init error")
		return
	}
	/*
		c := init_crawl()
		err := c.Visit("https://wiki.osdev.org")
		if err != nil {
			panic(err)
		}
		fmt.Println("finish")

		err2 := c.Visit("https://www.nesdev.org/wiki/")
		if err2 != nil {
			panic(err2)
		}
		fmt.Println("finish")
	*/
	// 初始化一个新的 Gin 路由器实例
	router := gin.Default()

	// 从 "templates" 目录加载 HTML 模板
	router.LoadHTMLGlob("templates/*")

	// 定义一个 GET 路由来提供搜索页面
	router.GET("/", func(c *gin.Context) {
		// 当访问根 URL 时，渲染 "search.html" 模板
		c.HTML(http.StatusOK, "search.html", nil)
	})

	// 定义一个 POST 路由来处理搜索请求
	router.POST("/search", func(c *gin.Context) {
		// 定义一个结构体来绑定 JSON 请求负载
		var json struct {
			Query string `json:"query"`
		}

		// 绑定 JSON 负载到结构体并检查错误
		if err := c.ShouldBindJSON(&json); err != nil {
			// 如果有错误，响应一个 Bad Request 状态和错误信息
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 使用请求负载中的查询执行搜索
		//ans := merge_list(str)

		str := strings.Split(clean_str(json.Query), " ")
		ans := merge_list_dc(str, 0, len(str)-1)
		results := return_web(ans)

		// 以 JSON 格式响应搜索结果
		c.JSON(http.StatusOK, results)
	})

	// 在端口 50086 上启动服务器
	router.Run(":50086")
}
