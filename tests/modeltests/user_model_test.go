package modeltests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rafaelandrade/API-RedCoins/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("Erro pegando usuario: %v\n", err)
		return
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser := models.User{
		ID:       1,
		Email: "teste@teste.com",
		Senha: "senha",
		Nome: "Teste numero um",
		DtNasc: "10/10/10",
	}
	savedUser, err := newUser.SaverUser(server.DB)
	if err != nil {
		t.Errorf("Este e o erro em salvar o usuario: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Nome, savedUser.Nome)
	assert.Equal(t, newUser.DtNasc, savedUser.DtNasc)
}

func TestGetUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Não pode enviar usuarios para tabela: %v", err)
	}
	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("Este e o erro ao tentar pegar um usuario: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nome, user.Nome)
	assert.Equal(t, foundUser.DtNasc, user.DtNasc)
}


func TestDeleteAUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()

	if err != nil {
		log.Fatalf("Não pode enviar usuario: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("Este e o erro atualizando o usuario: %v\n", err)
		return
	}

	assert.Equal(t, isDeleted, int64(1))
}