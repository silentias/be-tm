package user

type User struct {
	Id       string `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password" type:"varchar(100)"`
}

type RegistrationCode struct {
	Id    string `gorm:"primaryKey"`
	Email string
	Code  int32
}
