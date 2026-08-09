package model

const savedDockerfileDefaultName = "dockerfile_"

type Dockerfile struct {
	Identifier Identifier
	OriginPath string
}

func CreateDockerfile(identifier Identifier, originPath string) Dockerfile {
	return Dockerfile{
		Identifier: identifier,
		OriginPath: originPath,
	}
}
