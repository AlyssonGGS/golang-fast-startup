package settings

var Settings struct {
	Server struct {
		Port    string `envconfig:"default=5000"`
		Context string `envconfig:"default=startup"`
	}

	Database struct {
		UserName        string `envconfig:"default=root"`
		Password        string `envconfig:"default=root"`
		Host            string `envconfig:"default=localhost"`
		Port            string `envconfig:"default=3306"`
		DatabaseName    string `envconfig:"default=dbexample"`
		SetMaxOpenConns int    `envconfig:"default=100"`
	}
}
