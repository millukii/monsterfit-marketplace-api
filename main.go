package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"api/db"
	"api/handler"
	"api/repository"
	"api/service"
)

// init gets called before the main function
func init() {
    // Log error if .env file does not exist
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found")
    }
}

func main() {

	 var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)
 	defer cancel()
	 client,err := db.NewMongoConn("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false", ctx)
	 if err != nil {
		panic(err)
	 }
	productCollection, err := db.GetCollection(client, "monsterfit","products",ctx)
	if err != nil {
		panic(err)
	 }
	 productRepo := repository.NewProductRepo(productCollection)
	 productService := service.NewProductService(productRepo)
	 productHandler := handler.NewProductHandler(productService)



	 	e := echo.New()
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"http://localhost:3000", "*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		}))

		

    // Its great to version your API's
    v1 := e.Group("/api/v1")
    {
        v1.GET("/products/", productHandler.Find)
				v1.GET("/products/:id/", productHandler.FindById)
				v1.POST("/products/", productHandler.Create)
				v1.PATCH("/products/:id/", productHandler.Update)
				v1.DELETE("/products/:id/", productHandler.Delete)
    }
		e.Logger.Fatal(e.Start(":5000"))
}
func CORS(c *gin.Context) {

	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {
		
		c.Next()

	} else {
        
		// Everytime we receive an OPTIONS request, 
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real 
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}
func CORSConfig() cors.Config {
    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{"http://localhost:3000"}
    corsConfig.AllowCredentials = true
    corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
    corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
    return corsConfig
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Add("Content-Type", "application/json")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
    	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS,POST,PATCH,DELETE")
   	c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Next()
	}
}
func OptionMessage(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PATCH, DELETE")
	c.Header("Access-Control-Allow-Headers", "X-Token")
	c.Header("Access-Control-Allow-Credentials", "true")
}
