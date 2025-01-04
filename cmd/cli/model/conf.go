package model

type Config struct {
	AppVersion     string    `yaml:"appVersion"`
	CurrentServer  string    `yaml:"current-server"`
	CurrentProfile string    `yaml:"current-profile"`
	Servers        []Server  `yaml:"servers"`
	Profiles       []Profile `yaml:"profiles"`
}

type Server struct {
	Name        string `yaml:"name"`
	Host        string `yaml:"host"`
	Description string `yaml:"description"`
}

type Profile struct {
	Name        string `yaml:"name"`
	ApiKey      string `yaml:"apiKey"`
	Description string `yaml:"description"`
}

func (conf *Config) FindProfile(name string) *Profile {
	for _, profile := range conf.Profiles {
		if profile.Name == name {
			return &profile
		}
	}
	return nil
}

func (conf *Config) RemoveProfile(name string) {
	for i, profile := range conf.Profiles {
		if profile.Name == name {
			conf.Profiles = append(conf.Profiles[:i], conf.Profiles[i+1:]...)
			return
		}
	}
}

func (conf *Config) FindServer(name string) *Server {
	for _, server := range conf.Servers {
		if server.Name == name {
			return &server
		}
	}
	return nil
}

func (conf *Config) RemoveServer(name string) {
	for i, server := range conf.Servers {
		if server.Name == name {
			conf.Servers = append(conf.Servers[:i], conf.Servers[i+1:]...)
			return
		}
	}
}
