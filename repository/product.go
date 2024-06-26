package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"

	"api/entities"
	"api/models"
)

type ProductRepo interface {
	Create(product *models.Product, c echo.Context) (*entities.MProduct, error)
	Update(id string, product *models.Product, c echo.Context) (*entities.MProduct, error)
	FindById(id string, c echo.Context) (*entities.MProduct, error)
	Find(page int64, limit int64, c echo.Context) ([]*entities.MProduct, error)
	Delete(id string, c echo.Context) error
}

type ProductRepoImpl struct {
	collection *mongo.Collection
}

func NewProductRepo(collection *mongo.Collection) ProductRepo {
	return &ProductRepoImpl{collection}
}
func (p ProductRepoImpl ) 	Create(product *models.Product, ctx echo.Context) (*entities.MProduct, error){
	c := context.Background()
result, err := p.collection.InsertOne(c, product)
// check for errors in the insertion
	if err != nil {
					log.Print(err)
					return nil, err
	}
	log.Print(result.InsertedID)
	createdProduct, err := p.FindById(fmt.Sprint(result.InsertedID), ctx)
		if err != nil {
					log.Print(err)
					return nil, err
	}
	return createdProduct, nil
}
	func (p ProductRepoImpl ) 	Update(id string, product *models.Product, ctx echo.Context) (*entities.MProduct, error){
		
filter := bson.D{{Name: "internalCode", Value: id}}
	c := context.Background()
updateResult, err := p.collection.UpdateOne(c, filter, bson.D{
	{Name: "$set", Value: product},
})
		if err != nil {
					log.Print(err)
					return nil, err
	}

fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	updatedProduct, err := p.FindById(fmt.Sprint(updateResult.UpsertedID), ctx)
		if err != nil {
					log.Print(err)
					return nil, err
	}
	return updatedProduct, nil
	}
func (p ProductRepoImpl ) 		FindById(id string, ctx echo.Context) (*entities.MProduct, error){
	var result entities.MProduct
//filter := bson.M{"internalCode": id}
//filter := bson.D{{"quantity", bson.D{{"$eq", "mf-0000001"}}}}
/* project := bson.D{{ "internalCode", 1 }}
opts := options.FindOne().SetProjection(project) */
	c := context.Background()
res := p.collection.FindOne(c, bson.M{"internalCode": id})
if res.Err() != nil{
		log.Println("FindOne repository result error ", res.Err())
		return nil, res.Err()
}
err := res.Decode(&result)
if err != nil {
	log.Println("FindOne repository decode error ", err)
	return nil, err
}
fmt.Printf("product found: %v\n", result)
		return &result, nil
}
func (p ProductRepoImpl ) 		Find(page int64, limit int64, ctx echo.Context) ([]*entities.MProduct, error){
	   result := make([]*entities.MProduct, 0)

	c := context.Background()
   curr, err := p.collection.Find(c, bson.M{})

   if err != nil {
      return nil, err
   }
	 err = curr.All(c, &result) 
   if err != nil {
      return nil, err
   }
   return result, nil
}
func (p ProductRepoImpl ) 		Delete(id string, ctx echo.Context) error{
		return nil
}



type mongoPaechoate struct {
   limit int64
   page int64
}
func newMongoPaechoate(limit, page int64) *mongoPaechoate {
   return &mongoPaechoate{
      limit: int64(limit),
      page:  int64(page),
   }
}

func (mp *mongoPaechoate) getPaechoatedOpts() *options.FindOptions {
   l := mp.limit
   skip := mp.page*mp.limit - mp.limit
   fOpt := options.FindOptions{Limit: &l, Skip: &skip}

   return &fOpt
}