package entity

type User struct {
	ID            int
	Name          string
	Token         string
	WalletBalance float64
}

var (
	Users = []User{
		{ID: 1, Name: "Charu", WalletBalance: 1000, Token: "token1"},
		{ID: 2, Name: "Swati", WalletBalance: 1500, Token: "token2"},
		{ID: 3, Name: "Prajkta", WalletBalance: 800, Token: "token3"},
		{ID: 4, Name: "Akash", WalletBalance: 1200, Token: "token4"},
	}
)
