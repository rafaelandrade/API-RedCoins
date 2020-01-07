package modeltests

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rafaelandrade/API-RedCoins/api/controllers"
	"github.com/rafaelandrade/API-RedCoins/api/models"
)

var server = controllers.Server{}
var userInstance = models.User{}

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("../../.env"))
	if err != nil {
		log.Fatalf("Erro sem encontrar env %v\n", err)
	}
	Database()

	os.Exit(m.Run())
}

func Database() {

	var err error

	TestDbDriver := os.Getenv("TesteDB_DRIVER")

	if TestDbDriver == "mysql" {
		TesteDBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("TesteDB_USER"), os.Getenv("TesteDB_PASSWORD"), os.Getenv("TesteDB_HOST"),  os.Getenv("TesteDB_PORT"), os.Getenv("TesteDB_NAME"))
		server.DB, err = gorm.Open(TestDbDriver, TesteDBURL)
		if err != nil {
			fmt.Printf("Não consegue conectar na  base de dados %s \n", TestDbDriver)
			log.Fatal("Erro:", err)
		} else {
			fmt.Printf("Conectamos na %s base de dados\n", TestDbDriver)
		}
	}
}

func refreshUserTable() error {
	err := server.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = server.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Tabela atualizada com sucesso")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	user := models.User{

		Email: "pet@gmail.com",
		Senha: "senha",
		Nome: "Pet",
		DtNasc: "10/10/10",
	}

	err := server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("Não pode enviar o usuario para tabela: %v", err)
	}
	return user, nil
}

func seedUsers() error {

	users := []models.User{
		models.User{
			Email: "jorge@gmail.com",
			Senha: "senha",
			Nome: "Jorge",
			DtNasc: "11/11/11",
		},
		models.User{
			Email: "carlos@gmail.com",
			Senha: "senha",
			Nome: "Carlos",
			DtNasc: "15/15/15",
		},
	}

	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}