package datastore

import (
	"orm_selasa/entities"

	"gorm.io/gorm"
)

type BarangDB struct {
	DB *gorm.DB
}

func (b BarangDB) InsertBarang(nama string, harga int, stok int) {
	barang := entities.Product{Name: nama, Price: harga, Stock: stok}

	result := b.DB.Create(&barang)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b BarangDB) GetAllProduct() ([]entities.Product, error) {
	var products []entities.Product

	if err := b.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}