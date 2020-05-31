package conf

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	ClanName  string
}

type Clan struct {
	Name 		string
	Members 	int64
	Health		float64
	Morale  	float64
	Enemy 		string
	Resources 	int64
}