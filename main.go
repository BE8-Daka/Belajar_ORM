package main

import (
	"fmt"
	"orm_selasa/datastore"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

func connectDB(c Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	return db
}

func readEnv() Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("ERROR LOAD FILE", err)
	}

	res := Config{}
	res.Host = os.Getenv("DB_HOST")
	res.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	res.Database = os.Getenv("DB_DATABASE")
	res.Username = os.Getenv("DB_USERNAME")
	res.Password = os.Getenv("DB_PASSWORD")

	return res
}

func main() {
	
	isi := readEnv()
	db := connectDB(isi)

	// db.AutoMigrate(&User{})
	// db.AutoMigrate(&Product{})
	
	// User
	AmbilUser := datastore.UserDB{DB: db}
	
	semuaUser, err := AmbilUser.GetAllUser()
	if err != nil {
		panic(err)
	}

	fmt.Println(semuaUser)

	// Barang
	AmbilBarang := datastore.BarangDB{DB: db}
	
	semuaBarang, err := AmbilBarang.GetAllProduct()
	if err != nil {
		panic(err)
	}

	fmt.Println(semuaBarang)

	// Buat User
	// user := entities.User{Name: "Mahmuda Karima", Username: "daka", Password: "123", Email: "dakasakti.id@gmail.com"}

	// result := db.Create(&user)
	// if result.Error != nil {
	// 	panic(result.Error)
	// }

	// Buat Barang
	// InsertBarang(db, "Buku", 10000, 10)
}