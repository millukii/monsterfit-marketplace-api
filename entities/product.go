package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

)

type MProduct struct {
	Id        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	InternalCode     string             `json:"internalCode,omitempty" bson:"internalCode,omitempty"`
	Sku   string             `json:"sku,omitempty" bson:"sku,omitempty"`
	Image     string             `json:"image,omitempty" bson:"image,omitempty"`
	VendorCode      string             `json:"vendorCode,omitempty" bson:"vendorCode,omitempty"`
		VendorProductId      string             `json:"vendorProductId,omitempty" bson:"vendorProductId,omitempty"`
			Name      string             `json:"name,omitempty" bson:"name,omitempty"`
				Description      string             `json:"description,omitempty" bson:"description,omitempty"`
					Price      float64             `json:"price,omitempty" bson:"price,omitempty"`
						Cost      float64             `json:"cost,omitempty" bson:"cost,omitempty"`
	CreateAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}