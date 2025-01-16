package data

type LoginResponse struct {
	Token     string
	UserID    int64
	SekolahID int
	Role      string
	Status    string
}
