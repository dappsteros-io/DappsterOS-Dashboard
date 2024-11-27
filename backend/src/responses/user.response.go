package responses

import "github.com/kamva/mgm/v3"

type UserResponse struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email" binding:"required"`
	Username         string `json:"username" bson:"username" binding:"required"`
}
