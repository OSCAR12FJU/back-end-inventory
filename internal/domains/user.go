package domains

type Users struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Age         int    `json:"age"`
	Nacionality string `json:"nacionality"`
	Image       string `json:"image"`
}

type UsersResponse struct {
	Token string `json:"token"`
	User  *Users `json:"user"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
