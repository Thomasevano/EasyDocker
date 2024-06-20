package builders

import (
	"github.com/Thomasevano/EasyDocker/src/models"
)

func DockerComposeVolumeBuilder(volumes []models.ManagedVolume) map[string]struct{} {
	dockerComposeVolumes := make(map[string]struct{}, len(volumes))

	for _, volume := range volumes {
		dockerComposeVolumes[volume.Name] = struct{}{}
	}

	return dockerComposeVolumes
}
