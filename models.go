package main

import (
	"gorm.io/gorm"
)

type Item struct {
	ID         int            `json:"id" gorm:"column:id;primaryKey;"`
	ProductID  int            `json:"product_id" gorm:"column:product_id"`
	Name       string         `json:"name" gorm:"column:name"`
	Descripton string         `json:"descripton" gorm:"column:descripton"`
	DeletedAt  gorm.DeletedAt `json:"deleteDate" example:"31/12/9999 23:59" swaggertype:"string" gorm:"index;column:deleted_at"`
	Product    *Product       `json:"category" gorm:"references:product_id;foreignKey:id"`
}

func (*Item) TableName() string {
	return "item"
}

type Category struct {
	ID        int            `json:"id" gorm:"column:id;primaryKey;"`
	Name      string         `json:"name" gorm:"column:name"`
	DeletedAt gorm.DeletedAt `json:"deleteDate" example:"31/12/9999 23:59" swaggertype:"string" gorm:"index;column:deleted_at"`
	Products  []*Product     `json:"products" gorm:"references:id;foreignKey:category_id;"`
}

func (*Category) TableName() string {
	return "category"
}

type Product struct {
	ID         int        `json:"id" gorm:"column:id;primaryKey;"`
	CategoryID int        `json:"category_id" gorm:"column:category_id"`
	Name       string     `json:"name" gorm:"column:name"`
	Price      int        `json:"price" gorm:"column:price"`
	Category   *Category  `json:"category" gorm:"references:category_id;foreignKey:id;"` // colume name id, not is field name
	Items      []*Item    `json:"items" gorm:"references:id;foreignKey:product_id"`
	Factories  []*Factory `gorm:"many2many:product_factories;foreignKey:id;joinForeignKey:product_id;References:id;joinReferences:factory_id"`
}

func (*Product) TableName() string {
	return "product"
}

type Factory struct {
	ID        int         `json:"id" gorm:"column:id;primaryKey;"`
	Name      string      `json:"name" gorm:"column:name"`
	Address   string      `json:"address" gorm:"column:address"`
	Workshops []*Workshop `json:"workshop" gorm:"foreignKey:FactoryID"`
	Products  []*Product  `gorm:"many2many:product_factories;foreignKey:id;joinForeignKey:factory_id;References:id;joinReferences:product_id"`
}

func (*Factory) TableName() string {
	return "factory"
}

type Workshop struct {
	ID        int      `json:"id" gorm:"column:id;primaryKey;"`
	FactoryID int      `json:"factory_id" gorm:"column:factory_id"`
	Name      string   `json:"name" gorm:"column:name"`
	Factory   *Factory `json:"factory" gorm:"references:factory_id;foreignKey:id"`
}

func (*Workshop) TableName() string {
	return "workshop"
}

type ProductFactory struct {
	ID        int `json:"id" gorm:"column:id;primaryKey;"`
	ProductID int `json:"product_id" gorm:"column:product_id;"`
	FactoryID int `json:"factory_id" gorm:"column:factory_id;"`
	Price     int `json:"price" gorm:"column:price;"`
}

func (*ProductFactory) TableName() string {
	return "product_factories"
}
