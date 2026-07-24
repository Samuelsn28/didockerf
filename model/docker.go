package model

const savedDockerfileDefaultName = "dockerfile-"

type Dockerfile struct {
	Name string 
	Version string
	OriginPath string
}

func (d Dockerfile) GetFileName() string {
	return savedDockerfileDefaultName + d.Name + "-" + d.Version
}


