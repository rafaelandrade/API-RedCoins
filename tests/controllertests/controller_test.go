package controllertests

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
		log.Fatalf("Erro ao conseguir o env %v\n", err)
	}
	Database()

	os.Exit(m.Run())

}

func Database() {
	var err error

	TestDbDriver := os.Getenv("TestDbDriver")

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
	log.Printf("Sucesso em atualizar a tabela!")
	return nil
}

func seedOneUser() (models.User, error) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user := models.User{
			Email: "tulio@gmail.com",
			Senha: "senha",
			Nome: "Tulio",
			DtNasc: "11/11/1122",
		}

	err = server.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}


func seedUsers() ([]models.User, error) {

	var err error
	if err != nil {
		return nil, err
	}
	users := []models.User{
		models.User{
			Email: "vitor@gmail.com",
			Senha: "senha",
			Nome: "Vitor",
			DtNasc: "11/11/1293",
		},
		models.User{
			Email: "Josue@gmail.com",
			Senha: "senha",
			Nome: "Josue",
			DtNasc: "15/15/1512",
		},
	}
	for i, _ := range users {
		err := server.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return []models.User{}, err
		}
	}
	return users, nil
}


