package builders

import "github.com/Thomasevano/EasyDocker/src/models"

func BuildDockerComposeNetworks(networks []models.Network) map[string]models.DockerComposeNetwork {
	dockerComposeNetworks := make(map[string]models.DockerComposeNetwork)

	for _, network := range networks {
		dockerComposeNetworks[network.Name] = BuildDockerComposeNetwork(network)
	}

	return dockerComposeNetworks
}

func BuildDockerComposeNetwork(network models.Network) models.DockerComposeNetwork {
	return models.DockerComposeNetwork{
		External: network.IsExternal,
		Driver:   network.Driver,
	}
}
