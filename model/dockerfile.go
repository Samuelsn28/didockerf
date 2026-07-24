package model

const savedDockerfileDefaultName = "dockerfile_"

type Dockerfile struct {
	Name string 
	Version string
	OriginPath string
}

func (d Dockerfile) GetFileName() string {
	return savedDockerfileDefaultName + d.Name + "_" + d.Version
}


