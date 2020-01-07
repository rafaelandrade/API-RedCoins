package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "email") {
		return errors.New("Já existe email parecido!")
	}

	if strings.Contains(err, "hashedSenha") {
		return errors.New("Senha incorreta!")
	}

	return errors.New("Detalhes incorretos!")
}