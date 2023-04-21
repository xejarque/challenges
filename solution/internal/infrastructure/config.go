package infrastructure

type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	Driver       string
}

func LocalDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Host:         "localhost",
		Port:         5432,
		User:         "postgres",
		Password:     "mypassword",
		DatabaseName: "mydb",
		Driver:       "postgres",
	}
}

// add any other configs here
