package requests

type UserRequest struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}
