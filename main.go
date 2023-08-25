package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// album代表关于唱片专辑的数据。
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// 专辑切片以种子记录专辑数据。
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// 从album结构切片创建 JSON，并将 JSON 写入响应。
func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "访问首页成功",
	}) // 从包中传递常量net/http来指示HTTP 状态代码
}

// 从album结构切片创建 JSON，并将 JSON 写入响应。
func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums) // 从包中传递常量net/http来指示HTTP 状态代码
}

// 查找与客户端id相匹配的专辑，然后返回该专辑作为响应。
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// 遍历专辑列表，寻找 ID 值与参数匹配的专辑。
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// 添加来自请求正文中接收到的 JSON 的专辑。
func postAlbums(c *gin.Context) {
	var newAlbum album

	// 调用 BindJSON 将接收到的 参数 绑定到newAlbum。
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// 将新专辑添加到切片中。
	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

// 删除客户端id相匹配的专辑，然后返回该专辑作为响应
func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// 遍历专辑列表，寻找 ID 值与参数匹配的专辑。
	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	println("http://localhost:8080")
	// 初始化 Gin 路由器 Default
	router := gin.Default()
	router.GET("/", home)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	// router.DELETE("/albums/:id", deleteAlbums)

	// 将路由器连接到http.Server并启动服务器
	router.Run("localhost:8080")
}
