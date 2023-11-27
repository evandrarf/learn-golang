package database

var connection string = "PostgreSQL"

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}