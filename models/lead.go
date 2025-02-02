package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	FirstName      string `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName       string `json:"last_name" gorm:"type:varchar(100);not null"`
	Company        string `json:"company" gorm:"type:varchar(255)"`
	DepartmentName string `json:"department_name" gorm:"type:varchar(255)"`
	Phone          string `json:"phone" gorm:"type:varchar(20);not null"`
	Email          string `json:"email" gorm:"type:varchar(100);unique;not null"`
	Street         string `json:"street" gorm:"type:varchar(255)"`
	City           string `json:"city" gorm:"type:varchar(100)"`
	Zip            string `json:"zip" gorm:"type:varchar(20)"`
	State          string `json:"state" gorm:"type:varchar(100)"`
	Country        string `json:"country" gorm:"type:varchar(100)"`
	Website        string `json:"website" gorm:"type:varchar(255)"`
	Industry       string `json:"industry" gorm:"type:varchar(255)"`
}
type Estimate struct {
	gorm.Model
	EstimateTitle string          `json:"estimate_title" gorm:"type:varchar(255);not null"`
	EstimateJSON  json.RawMessage `json:"estimate_json" gorm:"type:json"` // Storing JSON array
	SubTotal      float64         `json:"sub_total" gorm:"type:decimal(10,2);not null"`
	TaxRate       float64         `json:"tax_rate" gorm:"type:decimal(5,2);not null"`
	TotalTax      float64         `json:"total_tax" gorm:"type:decimal(10,2);not null"`
	Total         float64         `json:"total" gorm:"type:decimal(10,2);not null"`
}
type SupplierMaterial struct {
	gorm.Model
	Item      string          `json:"item" gorm:"type:varchar(255);not null"`
	Suppliers json.RawMessage `json:"suppliers" gorm:"type:json"`
}

type Quote struct {
	Estimate_Id int `json:"estimate_id" gorm:"type:integer;not null"`
}
