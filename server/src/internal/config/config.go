package config

type App struct {
	Name string `yaml:"name"`
	Port int    `yaml:"port"`
}

type Postgres struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Db struct {
	Postgres Postgres `yaml:"postgres"`
}

type Logger struct {
	Level int  `yaml:"level"`
	Trace bool `yaml:"trace"`
}

type Metrics struct {
	Collect bool `yaml:"collect"`
}

type Config struct {
	App     App     `yaml:"app"`
	Db      Db      `yaml:"db"`
	Logger  Logger  `yaml:"logger"`
	Metrics Metrics `yaml:"metrics"`
}
