package handler

import (
	"log"
	"strconv"

	"api/models"
	"api/service"

	"github.com/labstack/echo"
)

type ProductHandler interface {
	Create(c echo.Context)  error
	Update( c echo.Context)  error
	Find( c echo.Context) error
	FindById( c echo.Context) error
	Delete(c echo.Context) error
}
type ProductHandlerImpl struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) ProductHandler {
	return &ProductHandlerImpl{service}
}

	func(h ProductHandlerImpl)	Create(c echo.Context) error {

		var body models.Product

    if err := c.Bind(&body); err != nil {
        return err
    }

		product, err :=	h.service.Create(&body, c)

		if err != nil {
				c.JSON(500, err)
				return err
		}
		c.JSON(200, product)
		return nil
	}
		func(h ProductHandlerImpl)	Update( c echo.Context) error{
			
			id :=c.Param("id")
		    var body models.Product

    if err := c.Bind(&body); err != nil {
        return err
    }

		product, err :=	h.service.Update(id,&body, c)

		if err != nil {
			c.JSON(500, err)
			return err
		}
			c.JSON(200, product)
			return nil
		}
		func(h ProductHandlerImpl)		Find( c echo.Context) error{
			pageParam := c.QueryParam("page")
					page,_ :=strconv.ParseInt(pageParam,10,64)
		   	limitParam :=c.QueryParam("limit")
			limit,_ :=strconv.ParseInt(limitParam,10,64)
						products, err := h.service.Find(page, limit, c)

		if err != nil {
				c.JSON(500, err)
				return err
		}
		log.Printf("API RESPONSE %+v", products)
		c.JSON(200,products )
		return nil
		}
		func(h ProductHandlerImpl)	FindById( c echo.Context) error {
			id := c.QueryParam("id")
			product, err := h.service.FindById(id, c)

			if err != nil {
				c.JSON(500, err)
				return err
			}
			c.JSON(200, product)
			return nil
		}
		func(h ProductHandlerImpl)	Delete(c echo.Context) error{
			return nil
		}