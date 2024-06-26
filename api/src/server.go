package main

import (
	_ "github.com/Thomasevano/EasyDocker/docs"
	"github.com/Thomasevano/EasyDocker/src/controllers"
	"github.com/Thomasevano/EasyDocker/src/initializers"
	middleware "github.com/Thomasevano/EasyDocker/src/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/spf13/viper"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

// @title EasyDocker API
// @BasePath /
func main() {
	app := fiber.New()
	micro := fiber.New()
	config, _ := initializers.LoadConfig(".")
	frontUrl := viper.GetString("FRONT_URL")
	frontPort := viper.GetString("FRONT_PORT")

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     frontUrl + ", http://localhost, http://localhost:" + frontPort + ", " + config.CORSUrl,
		AllowHeaders:     "Origin, Content-Type, Accept, X-Requested-With",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))
	app.Mount("/", micro)

	micro.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controllers.SignUpUser)
		router.Post("/login", controllers.SignInUser)
		router.Post("/logout", controllers.LogoutUser)
	})

	users := micro.Group("/users", middleware.DeserializeUser)
	users.Get("/me", controllers.GetMe)

	stacks := micro.Group("/stacks", middleware.DeserializeUser)
	stacks.Get("/", controllers.GetStacks)
	stacks.Get("/:id", controllers.GetStack)
	stacks.Post("/", controllers.CreateStack)
	stacks.Put("/:id", controllers.UpdateStack)
	stacks.Delete("/:id", controllers.DeleteStack)
	stacks.Post("/:id/duplicate", controllers.DuplicateStack)

	stacks.Post("/:stackId/services", controllers.CreateService)
	stacks.Get("/:stackId/services", controllers.GetServices)

	stacks.Post("/:stackId/networks", controllers.CreateNetwork)
	stacks.Get("/:stackId/networks", controllers.GetNetworks)

	network := micro.Group("/networks", middleware.DeserializeUser)
	network.Get("/:id", controllers.GetNetwork)
	network.Put("/:id", controllers.UpdateNetwork)
	network.Delete("/:id", controllers.DeleteNetwork)

	stacks.Post("/:stackId/managed_volumes", controllers.CreateManagedVolume)

	managedVolume := micro.Group("/managed_volumes", middleware.DeserializeUser)
	managedVolume.Get("/:id", controllers.GetManagedVolume)
	managedVolume.Put("/:id", controllers.UpdateManagedVolume)
	managedVolume.Delete("/:id", controllers.DeleteManagedVolume)

	service := micro.Group("/services", middleware.DeserializeUser)
	service.Get("/:id", controllers.GetService)
	service.Put("/:id", controllers.UpdateService)
	service.Delete("/:id", controllers.DeleteService)

	service.Post("/:serviceId/ports", controllers.CreateServicePort)
	service.Get("/:serviceId/ports", controllers.GetServicePorts)

	servicePort := micro.Group("/ports", middleware.DeserializeUser)
	servicePort.Get("/:id", controllers.GetServicePort)
	servicePort.Put("/:id", controllers.UpdateServicePort)
	servicePort.Delete("/:id", controllers.DeleteServicePort)

	service.Post("/:serviceId/env_variables", controllers.CreateServiceEnvVariable)
	service.Get("/:serviceId/env_variables", controllers.GetServiceEnvVariables)

	serviceEnvVariable := micro.Group("/env_variables", middleware.DeserializeUser)
	serviceEnvVariable.Get("/:id", controllers.GetServiceEnvVariable)
	serviceEnvVariable.Put("/:id", controllers.UpdateServiceEnvVariable)
	serviceEnvVariable.Delete("/:id", controllers.DeleteServiceEnvVariable)

	service.Post("/:serviceId/volumes", controllers.CreateServiceVolume)
	service.Get("/:serviceId/volumes", controllers.GetServiceVolumes)

	serviceVolume := micro.Group("/volumes", middleware.DeserializeUser)
	serviceVolume.Get("/:id", controllers.GetServiceVolume)
	serviceVolume.Put("/:id", controllers.UpdateServiceVolume)
	serviceVolume.Delete("/:id", controllers.DeleteServiceVolume)

	stacks.Get("/:stackId/board", controllers.GetBoard)

	serviceNetworkLink := micro.Group("/service_network_links", middleware.DeserializeUser)
	serviceNetworkLink.Post("/", controllers.CreateServiceNetworkLink)
	serviceNetworkLink.Delete("/:id", controllers.DeleteServiceNetworkLink)

	serviceManagedVolumeLink := micro.Group("/service_managed_volume_links", middleware.DeserializeUser)
	serviceManagedVolumeLink.Post("/", controllers.CreateServiceManagedVolumeLink)
	serviceManagedVolumeLink.Put("/:id", controllers.UpdateServiceManagedVolumeLink)
	serviceManagedVolumeLink.Delete("/:id", controllers.DeleteServiceManagedVolumeLink)

	stacks.Get("/:stackId/docker_compose", controllers.GenerateDockerComposeFile)

	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Fatal(app.Listen(":3000"))
}
