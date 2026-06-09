package domain

type User struct {
	ID        int64  
	Name      string 
	Email     string 
	AvatarUrl string 
}

type UpdateUserParam struct {
	Name string
	Email string
}