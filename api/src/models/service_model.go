package models

import "github.com/google/uuid"

type Service struct {
	ID          *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string     `gorm:"type:varchar(255)"`
	DockerImage string     `gorm:"type:varchar(255)"`
	DockerTag   string     `gorm:"type:varchar(255)"`
	Entrypoint  string     `gorm:"type:varchar(255)"`
	Description string     `gorm:"type:text"`
	PositionX   float32    `gorm:"type:decimal(20,8);not null"`
	PositionY   float32    `gorm:"type:decimal(20,8);not null"`
	Context     string     `gorm:"type:varchar(255)"`
	Dockerfile  string     `gorm:"type:varchar(255)"`

	StackID string `gorm:"type:uuid;not null"`
	Stack   Stack

	ServiceVolumes      []ServiceVolume
	ServiceEnvVariables []ServiceEnvVariable
	ServicePorts        []ServicePort
}

type ServiceCreateInput struct {
	Name        string  `json:"name"`
	DockerImage string  `json:"dockerImage"`
	DockerTag   string  `json:"dockerTag"`
	Entrypoint  string  `json:"entrypoint"`
	Description string  `json:"description"`
	PositionX   float32 `json:"positionX" validate:"required"`
	PositionY   float32 `json:"positionY" validate:"required"`
}

type ServiceUpdateInput struct {
	Name        string  `json:"name"`
	DockerImage string  `json:"dockerImage"`
	DockerTag   string  `json:"dockerTag"`
	Context     string  `json:"context"`
	Dockerfile  string  `json:"dockerfile"`
	Entrypoint  string  `json:"entrypoint"`
	Description string  `json:"description"`
	PositionX   float32 `json:"positionX"`
	PositionY   float32 `json:"positionY"`
}

type ServiceResponse struct {
	ID          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	DockerImage string     `json:"dockerImage"`
	DockerTag   string     `json:"dockerTag"`
	Entrypoint  string     `json:"entrypoint"`
	Description string     `json:"description"`
	PositionX   float32    `json:"positionX"`
	PositionY   float32    `json:"positionY"`
}

type ServiceResponseItem struct {
	ID          *uuid.UUID `json:"id"`
	Name        string     `json:"name"`
	DockerImage string     `json:"dockerImage"`
	DockerTag   string     `json:"dockerTag"`
	Entrypoint  string     `json:"entrypoint"`
	Description string     `json:"description"`
	PositionX   float32    `json:"positionX"`
	PositionY   float32    `json:"positionY"`

	Volumes      []ServiceVolumeResponse      `json:"volumes"`
	EnvVariables []ServiceEnvVariableResponse `json:"envVariables"`
	Ports        []ServicePortResponse        `json:"ports"`
}
