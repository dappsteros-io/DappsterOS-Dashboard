package responses

import "github.com/kamva/mgm/v3"

type UserResponse struct {
	mgm.DefaultModel `bson:",inline"`
	Email            string `json:"email" bson:"email" binding:"required"`
	Username         string `json:"username" bson:"username" binding:"required"`
}

type CommonResponse struct {
	Success bool   `json:"success" bson:"success" binding:"required"`
	Title   string `json:"title" bson:"title" binding:"required"`
	Detail  string `json:"detail" bson:"detail" binding:"required"`
	Status  int16  `json:"status" bson:"status" binding:"required"`
	Data    any    `json:"data" bson:"data"`
	Message string `json:"message" bson:"message" binding:"required"`
}
