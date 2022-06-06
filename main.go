package main

import (
	"github.com/adityarizkyramadhan/tes_intern_delos/infrastructure/database"
	db_con "github.com/adityarizkyramadhan/tes_intern_delos/infrastructure/db_conn"
	"github.com/adityarizkyramadhan/tes_intern_delos/middleware"
	"github.com/adityarizkyramadhan/tes_intern_delos/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORS())
	envDb, err := database.NewDriverMysql()
	if err != nil {
		panic(err)
	}
	db, err := db_con.InitMYSql(envDb)
	if err != nil {
		panic(err)
	}
	route.Route(r, db)
	if err := r.Run(); err != nil {
		panic(err)
	}
}

//type test struct {
//}
//
//func (t test) Test(c *gin.Context) {
//	c.JSON(200, gin.H{
//		"tes": "Hai",
//	})
//}
//func main() {
//	env, _ := app.NewDriverApp()
//	fmt.Println(env.SecretKey)
//	r := gin.Default()
//
//	r.GET("/", func(c *gin.Context) {
//		token, _ := middleware.GenerateToken(100)
//		c.JSON(200, gin.H{
//			"text":  "Hello World",
//			"token": token,
//		})
//	})
//	a := r.Group("/a", middleware.ValidateJWToken())
//	a.GET("", func(c *gin.Context) {
//		c.JSON(200, gin.H{
//			"text": "Hello World",
//			"id":   c.MustGet("login"),
//		})
//	})
//	a.GET("/v", test{}.Test)
//	err := r.Run()
//	if err != nil {
//		panic(err)
//	}
//}
