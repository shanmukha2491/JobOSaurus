package models



type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
