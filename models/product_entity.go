package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	Id            primitive.ObjectID `bson:"_id" json:"-"`
	ProductId     string             `json:"product_id,omitempty"`
	UserId        string             `json:"createdby_id,omitempty"`
	Name          string             `json:"name,omitempty" validate:"required"`
	Price         float64            `json:"price,omitempty" validate:"required"`
	TaxPercentage float64            `json:"tax_percentage,omitempty" validate:"required"`
	Mrp           float64            `json:"mrp,omitempty" validate:"required"`
	TaxAmount     float64            `json:"tax_amount,omitempty"`
	CreatedAt     time.Time          `json:"created_at,omitempty"`
	UpdatedAt     time.Time          `json:"updated_at,omitempty"`
}

func (p *Product) ValidateTaxDetail() bool {
	var check bool = false
	if p.Price > 0 && p.TaxPercentage > 0 {
		check = true
	}
	return check
}

func (p *Product) CalculateTax() float64 {
	var taxamount = 0.0
	if p.ValidateTaxDetail() {
		taxamount := (p.Price * p.TaxPercentage) / 100
		return taxamount
	}
	return taxamount
}
