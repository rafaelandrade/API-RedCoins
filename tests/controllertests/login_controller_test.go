
package controllertests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignIn(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	user, err := seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}

	samples := []struct {
		email        string
		password     string
		errorMessage string
	}{
		{
			email:        user.Email,
			password:     "senha",
			errorMessage: "",
		},
		{
			email:        user.Email,
			password:     "Senha incorreta",
			errorMessage: "crypto/bcrypt: hashedPassword tipos diferentes de senha",
		},
		{
			email:        "Email incorreto",
			password:     "senha",
			errorMessage: "Não encontrado registro",
		},
	}

	for _, v := range samples {

		token, err := server.SignIn(v.email, v.password)
		if err != nil {
			assert.Equal(t, err, errors.New(v.errorMessage))
		} else {
			assert.NotEqual(t, token, "")
		}
	}
}

func TestLogin(t *testing.T) {

	refreshUserTable()

	_, err := seedOneUser()
	if err != nil {
		fmt.Printf("This is the error %v\n", err)
	}
	samples := []struct {
		inputJSON    string
		statusCode   int
		email        string
		senha     string
		errorMessage string
	}{
		{
			inputJSON:    `{"email": "teste@teste.com", "senha": "senha"}`,
			statusCode:   200,
			errorMessage: "",
		},
		{
			inputJSON:    `{"email": "teste@teste.com", "senha": "Senha errada"}`,
			statusCode:   422,
			errorMessage: "Senha incorreta",
		},
		{
			inputJSON:    `{"email": "vitor@gmail.com", "senha": "senha"}`,
			statusCode:   422,
			errorMessage: "Detalhes incorretos",
		},
		{
			inputJSON:    `{"email": "kangmail.com", "senha": "senha"}`,
			statusCode:   422,
			errorMessage: "Email invalido",
		},
		{
			inputJSON:    `{"email": "", "senha": "senha"}`,
			statusCode:   422,
			errorMessage: "E necessario colocar o email",
		},
		{
			inputJSON:    `{"email": "kan@gmail.com", "password": ""}`,
			statusCode:   422,
			errorMessage: "E necessario colocar a senha",
		},
		{
			inputJSON:    `{"email": "", "senha": "senha"}`,
			statusCode:   422,
			errorMessage: "E necessario colocar o email",
		},
	}

	for _, v := range samples {

		req, err := http.NewRequest("POST", "/login", bytes.NewBufferString(v.inputJSON))
		if err != nil {
			t.Errorf("Erro: %v", err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(server.Login)
		handler.ServeHTTP(rr, req)

		assert.Equal(t, rr.Code, v.statusCode)
		if v.statusCode == 200 {
			assert.NotEqual(t, rr.Body.String(), "")
		}

		if v.statusCode == 422 && v.errorMessage != "" {
			responseMap := make(map[string]interface{})
			err = json.Unmarshal([]byte(rr.Body.String()), &responseMap)
			if err != nil {
				t.Errorf("Nao consegue converter para JSON: %v", err)
			}
			assert.Equal(t, responseMap["error"], v.errorMessage)
		}
	}
}