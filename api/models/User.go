package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"ID"`
	Email  string    `gorm:"size:255;not null;unique" json:"email"`
	Senha  string    `gorm:"size:100;not null;" json:"senha"`
	Nome  string    `gorm:"size:255;not null;" json:"nome"`
	DtNasc  string    `gorm:"not null;" json:"dtNasc"`
	vReal        uint32    `gorm:"not null;" json:"vReal"`
	vCoin        uint32    `gorm:"not null;" json:"vCoin"`
}

func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

func VerifyPassword(hashedSenha, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedSenha), []byte(senha))
}

func (u *User) BeforeSave() error {
	hashedSenha, err := Hash(u.Senha)
	if err != nil {
		return err
	}

	u.Senha = string(hashedSenha)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Nome = html.EscapeString(strings.TrimSpace(u.Nome))
	u.DtNasc = html.EscapeString(strings.TrimSpace(u.DtNasc))
	u.vReal = 0
	u.vCoin = 0
}

func (u *User) Validate (action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			return errors.New("Necessario colocar o email!!!")
		}

		if u.Senha == "" {
			return errors.New("Necessario colocar senha!!!")
		}
		if u.Nome == "" {
			return errors.New("Necessario colocar o nome!!!")
		}
		if u.DtNasc == "" {
			return errors.New("Necessario colocar a data de nascimento!!!")
		}

		return nil

	case "login":
		if u.Email == "" {
			return errors.New("Necessario colocar o email!!!")
		}
		if u.Senha == "" {
			return errors.New("Necessario colocar senha!!!")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email invalido!")
		}
		return nil

	default:
		if u.Email == "" {
			return errors.New("Necessario colocar o email!!!")
		}

		if u.Senha == "" {
			return errors.New("Necessario colocar senha!!!")
		}
		if u.Nome == "" {
			return errors.New("Necessario colocar o nome!!!")
		}
		if u.DtNasc == "" {
			return errors.New("Necessario colocar a data de nascimento!!!")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email invalido!")
		}
		return nil
	}
}

func (u *User) SaverUser(db *gorm.DB) (*User, error) {

	var err error

	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(50).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID( db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, err
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
