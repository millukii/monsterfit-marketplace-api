package handler

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"

	"api/models"
	"api/service"
)

type ProductHandler interface {
	Create(c *gin.Context) 
	Update( c *gin.Context) 
	Find( c *gin.Context)
	FindById( c *gin.Context)
	Delete(c *gin.Context) 
}
type ProductHandlerImpl struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) ProductHandler {
	return &ProductHandlerImpl{service}
}

	func(h ProductHandlerImpl)	Create(c *gin.Context) {

		var body models.Product

    if err := c.BindJSON(&body); err != nil {
        return
    }

		product, err :=	h.service.Create(&body, c)

		if err != nil {
				c.JSON(500, err)
				return
		}
		c.JSON(200, product)
	}
		func(h ProductHandlerImpl)	Update( c *gin.Context) {
			
			id :=c.Param("id")
		    var body models.Product

    if err := c.BindJSON(&body); err != nil {
        return
    }

		product, err :=	h.service.Update(id,&body, c)

		if err != nil {
			c.JSON(500, err)
			return
		}
			c.JSON(200, product)
		}
		func(h ProductHandlerImpl)		Find( c *gin.Context) {
			pageParam := c.Query("page")
					page,_ :=strconv.ParseInt(pageParam,10,64)
		   	limitParam :=c.Query("limit")
			limit,_ :=strconv.ParseInt(limitParam,10,64)
						products, err := h.service.Find(page, limit, c)

		if err != nil {
				c.JSON(500, err)
				return
		}
		log.Printf("API RESPONSE %+v", products)
		c.JSON(200,products )
		}
		func(h ProductHandlerImpl)	FindById( c *gin.Context) {
			id := c.Param("id")
			product, err := h.service.FindById(id, c)

			if err != nil {
				c.JSON(500, err)
				return
			}
			c.JSON(200, product)
		}
		func(h ProductHandlerImpl)	Delete(c *gin.Context) {}