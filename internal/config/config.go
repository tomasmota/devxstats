package config

type AppConfig struct {
	Port int       `env:"PORT,default=8080"`
	Db   *DbConfig `env:",prefix=DB_"`
	Git  *GitConfig
	Cd   *CdConfig
}

type DbConfig struct {
	Host string `env:"HOST,required"`
	Port int    `env:"PORT,default=27017"`
}

type GitConfig struct {
	Github struct {
		Token string `env:"GITHUB_TOKEN,required"`
		Url   string `env:"GITHUB_URL"`
	}
	Bitbucket struct {
		Url   string `env:"BITBUCKET_URL,required"`
		Token string `env:"BITBUCKET_TOKEN,required"`
	}
}

type CdConfig struct {
	Octopus struct {
		Token string `env:"OCTOPUS_TOKEN,required"`
		Url   string `env:"OCTOPUS_URL,required"`
	}
}
