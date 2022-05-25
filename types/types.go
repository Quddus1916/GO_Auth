package types

type Config struct {
	Port     string
	MongoUrl string
}

type LoginUser struct {
	Email    string
	Password string
}
