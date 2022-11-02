package service

import (
	"fmt"

	"api/models"
	"api/repository"

	"github.com/labstack/echo"
)
type ProductService interface {
	Create(product *models.Product, c echo.Context) (*models.Product, error)
	Update(id string,product *models.Product, c echo.Context) (*models.Product, error)
	FindById(id string, c echo.Context) (*models.Product, error)
	Find(page int64, limit int64, c echo.Context) ([]*models.Product, error)
	Delete(id string, c echo.Context) error
}

type ProductServiceImpl struct {
	productRepo repository.ProductRepo
}

func NewProductService(productRepo repository.ProductRepo) ProductService {
	return &ProductServiceImpl{productRepo}
}
func (p ProductServiceImpl ) 	Create(product *models.Product, ctx echo.Context) (*models.Product, error){

	dbProduct, err := p.productRepo.Create(product, ctx)
	if err != nil {
		return nil, err
	}
	return &models.Product{
		Image: dbProduct.Image,
		Id: dbProduct.Id.String(),
		InternalCode: dbProduct.InternalCode,
		Sku: dbProduct.Sku,
		VendorCode: dbProduct.VendorCode,
		VendorProductId: dbProduct.VendorProductId,
		Name: dbProduct.Name,
		Description: dbProduct.Description,
		Price: fmt.Sprint(dbProduct.Price),
		Cost: fmt.Sprint(dbProduct.Cost),
	}, nil
}
	func (p ProductServiceImpl ) 	Update(id string, product *models.Product, ctx echo.Context) (*models.Product, error){
		
	dbProduct, err := p.productRepo.Update(id,product, ctx)
	if err != nil {
		return nil, err
	}
	return &models.Product{
		Image: dbProduct.Image,
		Id: dbProduct.Id.String(),
		InternalCode: dbProduct.InternalCode,
		Sku: dbProduct.Sku,
		VendorCode: dbProduct.VendorCode,
		VendorProductId: dbProduct.VendorProductId,
		Name: dbProduct.Name,
		Description: dbProduct.Description,
		Price: fmt.Sprint(dbProduct.Price),
		Cost: fmt.Sprint(dbProduct.Cost),
	}, nil
	}
func (p ProductServiceImpl ) 		FindById(id string, ctx echo.Context) (*models.Product, error){
	dbProduct, err := p.productRepo.FindById(id, ctx)
	if err != nil {
		return nil, err
	}
	return &models.Product{
		Image: dbProduct.Image,
		Id: dbProduct.Id.String(),
		InternalCode: dbProduct.InternalCode,
		Sku: dbProduct.Sku,
		VendorCode: dbProduct.VendorCode,
		VendorProductId: dbProduct.VendorProductId,
		Name: dbProduct.Name,
		Description: dbProduct.Description,
		Price: fmt.Sprint(dbProduct.Price),
		Cost: fmt.Sprint(dbProduct.Cost),
	}, nil
}
func (p ProductServiceImpl ) 		Find(page int64, limit int64, ctx echo.Context) ([]*models.Product, error){
	dbProducts, err := p.productRepo.Find(page,limit, ctx)
	if err != nil {
		return nil, err
	}
	var products []*models.Product
 for _, j :=range dbProducts{

	products = append(products,	&models.Product{
		Image: j.Image,
		Id: j.Id.String(),
		InternalCode: j.InternalCode,
		Sku: j.Sku,
		VendorCode: j.VendorCode,
		VendorProductId: j.VendorProductId,
		Name: j.Name,
		Description: j.Description,
		Price: fmt.Sprint(j.Price),
		Cost: fmt.Sprint(j.Cost),
	})
 }
	return products , nil
}
func (p ProductServiceImpl ) 		Delete(id string, ctx echo.Context) error{
		return nil
}


