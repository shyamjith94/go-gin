package models

type Product struct {
	UserId        string  `json:"createdbyid,omitempty"`
	Name          string  `json:"name,omitempty" validate:"required"`
	Price         float64 `json:"price,omitempty" validate:"required"`
	TaxPercentage float64 `json:"taxpercentage,omitempty" validate:"required"`
	Mrp           float64 `json:"mrp,omitempty" validate:"required"`
	TaxAmount     float64 `json:"taxamount"`
}

func (p *Product) CalculateTax() float64 {
	taxamount := (p.Price * p.TaxPercentage) / 100
	return taxamount
}
