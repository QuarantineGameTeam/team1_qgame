package conf

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	ClanName  string
}
