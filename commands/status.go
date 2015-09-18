package commands

import (
	"fmt"

	"strconv"

	"github.com/catalyzeio/catalyze/helpers"
	"github.com/catalyzeio/catalyze/models"
	"github.com/fatih/color"
)

// Status prints out an environment healthcheck. The status of the environment
// and every service in the environment is printed out.
func Status(settings *models.Settings) {
	helpers.SignIn(settings)
	env := helpers.RetrieveEnvironment("pod", settings)
	fmt.Printf("%s (environment ID = %s):\n", env.Data.Name, env.ID)
	for _, service := range *env.Data.Services { // loops through services
		if service.Type != "utility" {
			b := color.New(color.Bold)
			b.Printf("\tService: %s - [ Ram = %s, Storage = %s ]\n", service.Label, strconv.Itoa(service.Ram), strconv.Itoa(service.Storage))
			jobs := helpers.RetrieveRunningJobs("service_id", settings)
			// fmt.Printf("%+v\n", jobs)
			fmt.Printf("\tID\t\tType\t\tStatus\t\tCreated At\n")
			fmt.Printf("\t------------------------------------------------------------\n")
			for jobID, job := range *jobs {
				fmt.Printf("\t%s\t\t%s\t\t%s\t\t%s\n", jobID[0:7], job.Type, job.Status, job.CreatedAt)
			}
			fmt.Printf("\n")
			//first, get information from service: serviceID, ['label'], ['size']['ram'], ['size']['storage']
			//next, pull running jobs:
			// RetrieveRunningJobs(service_id)
			// loop through these jobs:
			// for _, job := range *....
			// get information: jobID, type, status, created_at
			//if unsuccessful, drop a message about unable to find and move to next service

			// if service.Type == "code" {
			// 	switch service.Size.(type) {
			// 	case string:
			// 		printLegacySizing(&service)
			// 	default:
			// 		printNewSizing(&service)
			// 	}
			// } else {
			// 	switch service.Size.(type) {
			// 	case string:
			// 		sizeString := service.Size.(string)
			// 		defer fmt.Printf("\t%s (size = %s, image = %s, status = %s) ID: %s\n", service.Label, sizeString, service.Name, service.DeployStatus, service.ID)
			// 	default:
			// 		serviceSize := service.Size.(map[string]interface{})
			// 		defer fmt.Printf("\t%s (ram = %.0f, storage = %.0f, behavior = %s, type = %s, cpu = %.0f, image = %s, status = %s) ID: %s\n", service.Label, serviceSize["ram"], serviceSize["storage"], serviceSize["behavior"], serviceSize["type"], serviceSize["cpu"], service.Name, service.DeployStatus, service.ID)
			// 	}
			// }
		}
	}
}

func printLegacySizing(service *models.Service) {
	sizeString := service.Size.(string)
	fmt.Printf("\t%s (size = %s, build status = %s, deploy status = %s) ID: %s\n", service.Label, sizeString, service.BuildStatus, service.DeployStatus, service.ID)
}

func printNewSizing(service *models.Service) {
	serviceSize := service.Size.(map[string]interface{})
	fmt.Printf("\t%s (ram = %.0f, storage = %.0f, behavior = %s, type = %s, cpu = %.0f, build status = %s, deploy status = %s) ID: %s\n", service.Label, serviceSize["ram"], serviceSize["storage"], serviceSize["behavior"], serviceSize["type"], serviceSize["cpu"], service.BuildStatus, service.DeployStatus, service.ID)
}
