package repo

import (
	"github.com/dappster-io/DappsterOS-Dashboard/backend/src/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := mgm.Coll(&models.User{}).FindOne(mgm.Ctx(), bson.M{"email": email}).Decode(user)
	return user, err
}

func FindUserByID(id string) (*models.User, error) {
	user := &models.User{}
	old, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = mgm.Coll(&models.User{}).FindByID(old, user)
	return user, err
}
