package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func ping(c *gin.Context) {
	fmt.Println("in ping")
	t, err := template.ParseFiles("./templates/index.tmpl", "./templates/ping.tmpl")
	if err != nil {
		fmt.Println("parse templates failed", err)
		return
	}
	t.ExecuteTemplate(c.Writer, "ping.tmpl", "minxu是个大帅b")

}
func pong(c *gin.Context) {
	fmt.Println("in pong")
	t, err := template.ParseFiles("./templates/index.tmpl", "./templates/pong.tmpl")
	if err != nil {
		fmt.Println("parse templates failed", err)
		return
	}
	t.ExecuteTemplate(c.Writer, "pong.tmpl", "minxu是个大帅b")
}
func minxu(c *gin.Context) {
	data := gin.H{
		"name":    "闵续",
		"age":     18,
		"message": "hello world",
	}
	c.JSON(200, data)
}

type people struct {
	Name string `json:"name"`//注意 此处是``漂撇  而不是单引号
	Age int `json:"age"`
	Message string `json:"message"`
}

func jsontest(c *gin.Context) {
	data := people{
		Name:"闵续",
		Age:18,
		Message: "hello liu huan",
	}
	c.JSON(200, data)
}
func formGet(c *gin.Context) {
	name:=c.Query("name")
	fmt.Println(name,"提交表单formGet  重定向至/ping")
	c.Redirect(302,"/ping")//此处状态码不能为300
}
func toLogin(c *gin.Context) {
	fmt.Println("进入tologin方法")
	c.HTML(200,"login.html",nil)
}
func test(c *gin.Context) {
	c.HTML(200,"test.html",nil)
}
func formPost(c *gin.Context) {
	username:=c.PostForm("username")
	password:=c.PostForm("password")
	//c.DefaultPostForm()
	//c.GetPostForm()
	fmt.Println(username,password,"提交表单formPost 重定向至/pong")
	c.HTML(200,"dashboard.html",gin.H{
		"username":username,
		"password":password,
	})
	//c.Redirect(302,"/pong")//此处状态码不能为300
}
func path(c *gin.Context)  {//路径传值  需要注意的是 最好又前缀  不然容易错误匹配报错
	name:=c.Param("name")
	age:=c.Param("age")//返回的都是string类型
	fmt.Println(name,age)
	c.JSON(200,gin.H{
		"name":name,
		"age":age,
	})
}

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
func fengzhuang(c *gin.Context)  {
	//username:=c.Query("username")
	//password:=c.Query("password")
	var user User
	//如果传入user  会被拷贝 所以需要传递 指针/地址
	err:=c.ShouldBind(&user)//无论是get post（json）  都可以绑定数据
	if(err!=nil){
		c.JSON(200,gin.H{
			"message":err.Error(),
		})
	}else {
		c.JSON(200,gin.H{
			"user":user,
		})
	}

}
//文件上传
func upload(c *gin.Context)  {
	file,err:=c.FormFile("file")
	if(err!=nil){
		fmt.Println("FormFile failed")
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
	}else{
		err:=c.SaveUploadedFile(file,"./"+file.Filename)
		if(err!=nil){
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}else{
			c.JSON(200,gin.H{"message":"./"+file.Filename})
		}
	}
}
func toUpload(c *gin.Context)  {
	c.HTML(200,"upload.html",nil)
}
func filter(c *gin.Context)  {
	fmt.Println("filter in------")
	c.Next()//c.Abort()阻止后续函数 不进入后续拦截器 也不进入最后的请求 执行完当前函数就结束
	fmt.Println("filter out------")
}
func loginHander(c *gin.Context){

}

//代码格式化  opt+cmd+L
func main() {
	r := gin.Default()//默认使用了logger（打印日志）和recovery（返回500）中间件
	//如果使用协程  需要传入c.Copy()

	r.LoadHTMLGlob("./templates/*.html")
	//r.LoadHTMLFiles("./templates/login.html","./templates/dashboard.html","./templates/test.html")
	r.GET("/ping", ping)
	r.GET("/pong", pong)
	r.GET("/minxu", minxu)
	r.GET("/jsontest", jsontest)
	r.GET("/formGet", formGet)
	//r.Use(filter)  给下列方法添加过滤起
	//r.Use(filter1,filter2)  给下列方法添加多个过滤起  前一个set放值  后一个可以get取值
	r.GET("/toLogin",filter, toLogin)//给方法添加过滤器
	r.GET("/test", test)
	r.GET("/path/:name/:age", path)
	r.GET("/fengzhuang", fengzhuang)
	r.POST("/login", formPost)
	r.GET("/upload",toUpload)
	r.POST("/upload", upload)
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404,"404.html",nil)
	})

	userGroup:=r.Group("/user")
	{
		userGroup.GET("/add", func(c *gin.Context) {c.JSON(200,gin.H{"message":"/user/add"} )})
		userGroup.GET("/delete", func(c *gin.Context) {c.JSON(200,gin.H{"message":"/user/delete"})})
		userGroup.GET("/find", func(c *gin.Context) {c.JSON(200,gin.H{"message":"/user/find"})})
	}
	articleGroup:=r.Group("/article")
	{
		articleGroup.GET("/add", func(c *gin.Context) {c.JSON(200,gin.H{"message":"/article/add"})})
		articleGroup.GET("/delete", func(c *gin.Context) {c.JSON(200,gin.H{"message":"/article/delete"})})
		articleGroup.GET("/find", func(c *gin.Context) {c.JSON(200,gin.H{"message":"/article/find"})})
	}


	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
}
