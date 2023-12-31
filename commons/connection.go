package commons

import (
	"log"

	"github.com/eslopezm/Agenda_Personal/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8")

	if err != nil {
		log.Fatal(err)

	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("..Migrando....")

	db.AutoMigrate(&models.Persona{})
}
