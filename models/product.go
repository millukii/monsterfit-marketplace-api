package models

type Product struct {
	Id              string `json:"id,omitempty"`
	InternalCode    string `json:"internalCode,omitempty" `
	Sku             string `json:"sku,omitempty" `
	Image           string `json:"image,omitempty" `
	VendorCode      string `json:"vendorCode,omitempty"`
	VendorProductId string `json:"vendorProducId,omitempty" `
	Name            string `json:"name,omitempty" `
	Description     string `json:"description,omitempty" `
	Price           string `json:"price,omitempty"`
	Cost            string `json:"cost,omitempty"`
}
