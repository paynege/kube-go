package main

type Node struct {
	NodeName string `yaml:"name"`
	NodeIP   string `yaml:"IP"`
}

type Host struct {
	Master []*Node `yaml:"masters"`
	Nodes  []*Node `yaml:"nodes"`
}

type Config struct {
	HostInfo Host         `yaml:"hosts"`
	AppInfo  []*App       `yaml:"apps"`
	ParaInfo []*Parameter `yaml:"parameters"`
}

type App struct {
	AppName       string `yaml:"name"`
	AppSourcePath string `yaml:"source"`
	AppTargetPath string `yaml:"target"`
}

type Parameter struct {
	ParaName  string `yaml:"name"`
	ParaValue string `yaml:"value"`
}

var (
	hostConfig Config
)
