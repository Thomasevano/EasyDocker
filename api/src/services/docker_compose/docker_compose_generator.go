package docker_compose

import (
	"github.com/Thomasevano/EasyDocker/src/models"
	"github.com/Thomasevano/EasyDocker/src/services/docker_compose/builders"
	"gopkg.in/yaml.v3"
)

func GenerateDockerCompose(services []models.Service, networks []models.Network, volumes []models.ManagedVolume) string {
	dockerCompose := builders.BuildDockerCompose(services, networks, volumes)

	dockerComposeContent := DockerComposeToYaml(dockerCompose)

	return dockerComposeContent
}

func DockerComposeToYaml(dockerCompose models.DockerCompose) string {
	result, _ := yaml.Marshal(dockerCompose)

	return string(result)
}
