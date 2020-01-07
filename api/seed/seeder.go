package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/rafaelandrade/API-RedCoins/api/models"
)

var users = []models.User{
	models.User{
		Email: "teste@teste.com",
		Senha: "senha",
		Nome: "Teste numero um",
		DtNasc: "10/10/10",
	},
	models.User{
		Email: "rafs@rafs.com",
		Senha: "senha",
		Nome: "Teste numero dois",
		DtNasc: "10/10/11",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatalf("Erro drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatalf("Erro migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("Erro emvoar tabela users: %v", err)
		}
	}
}