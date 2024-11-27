package models

import (
	"log"

	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/utils"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`

	Username string `valid:"stringlength(5|20)" json:"username" bson:"username"`
	Email    string `valid:"stringlength(5|20)" json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Token    string `json:"token" bson:"token"`
}

func (u *User) CheckPassword(password string) error {
	return utils.VerifyPassword(u.Password, password)
}

func (p *User) Register() (*User, error) {
	pwd, err := utils.HashPassword(p.Password)
	if err != nil {
		return nil, err
	}
	p.Password = pwd

	err = mgm.Coll(&User{}).Create(p)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return p, nil
}

func (p *User) Get() ([]User, error) {
	result := []User{}
	err := mgm.Coll(&User{}).SimpleFind(&result, bson.M{})
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return result, nil
}

func (p *User) Post() (*User, error) {
	err := mgm.Coll(&User{}).Create(p)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return p, nil
}

func (p *User) Put() (*User, error) {
	err := mgm.Coll(p).Update(p)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return p, nil
}
