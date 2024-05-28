package main

import (
	"database/sql"
	"fmt"
	"os"
	"sanber-golang-56-paw/controllers"
	"sanber-golang-56-paw/database"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")

	host := os.Getenv("DB_HOST")
	strPort := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	port, _ := strconv.Atoi(strPort)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {

		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(db)

	defer db.Close()

	router := gin.Default()

	auth := router.Group("/auth")
	auth.POST("register", controllers.InsertCanvasser)
	auth.POST("login", controllers.GetLoginCanvasser)

	// Group route dengan middleware BasicAuth
	authorized := router.Group("/", Auth())
	{
		canvasser := authorized.Group("/canvasser")
		{
			canvasser.GET("/", controllers.GetAllCanvasser)
			canvasser.POST("/", controllers.InsertCanvasser)
			canvasser.PUT("/:id", controllers.UpdateCanvasser)
			canvasser.DELETE("/:id", controllers.DeleteCanvasser)
		}

		item := authorized.Group("/item")
		{
			item.GET("/", controllers.GetAllItem)
			item.POST("/", controllers.InsertItem)
			item.PUT("/:id", controllers.UpdateItem)
			item.DELETE("/:id", controllers.DeleteItem)
		}

		customer := authorized.Group("/customer")
		{

			customer.GET("/", controllers.GetAllCustomer)
			customer.POST("/", controllers.InsertCustomer)
			customer.PUT("/:id", controllers.UpdateCustomer)
			customer.DELETE("/:id", controllers.DeleteCustomer)
		}

		stock := authorized.Group("/stock")
		{

			stock.GET("/", controllers.GetAllStock)
			stock.POST("/", controllers.InsertStock)
			stock.PUT("/:item_id/:canvasser_id", controllers.UpdateStock)
			stock.DELETE("/:item_id/:canvasser_id", controllers.DeleteStock)
		}

		trnsales := authorized.Group("/trnsales")
		{
			trnsales.GET("/", controllers.GetAllTrnSales)
			trnsales.POST("/", controllers.InsertTrnSales)
			trnsales.PUT("/:id", controllers.UpdateTrnSales)
			trnsales.DELETE("/:id", controllers.DeleteTrnSales)
		}

		trnsalesdetail := authorized.Group("/trnsalesdetail")
		{
			trnsalesdetail.GET("/", controllers.GetAllTrnSalesDetail)
			trnsalesdetail.POST("/", controllers.InsertTrnSalesDetail)
			trnsalesdetail.PUT("/:id", controllers.UpdateTrnSalesDetail)
			trnsalesdetail.DELETE("/:id", controllers.DeleteTrnSalesDetail)
		}

		report := authorized.Group("/report")
		{

			report.GET("/stock", controllers.GetFormattedStock)
			report.GET("/sales", controllers.GetFormattedSales)
		}
	}

	router.Run(envPortOr("8080"))
}

func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}

// Fungi Log yang berguna sebagai middleware
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// uname, pwd, ok := c.Request.BasicAuth()
		// if !ok {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau Password tidak boleh kosong"})
		// 	c.Abort()
		// 	return
		// }

		// if (uname == "admin" && pwd == "password") || (uname == "editor" && pwd == "secret") {
		// 	c.Next()
		// 	return
		// }

		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Anda tidak Tidak Memiliki Hak Akses"})
		// c.Abort()
		// return

		c.Next()
		// return
	}
}
