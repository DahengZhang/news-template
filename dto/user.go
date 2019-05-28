package dto

// User 用户 DTO
type User struct{
	UID int `json:"uid"`
	Phone string `json:"phone"`
	Email string `json:"email"`
	Name string `json:"name"`
	Password string `json:"password,omitempty"`
	CreateTime string `json:"create_time"`
}