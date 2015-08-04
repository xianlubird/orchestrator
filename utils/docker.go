package utils

import (
	"log"
	"github.com/samalba/dockerclient"
)

// Callback used to listen to Docker's events
func eventCallback(event *dockerclient.Event, ec chan error, args ...interface{}) {
    log.Printf("Received event: %#v\n", *event)
}

func CreateContainer() string {
	docker := GetDockerClient()
	// Create a container
    containerConfig := &dockerclient.ContainerConfig{
        Image: "tomcat:7",
        AttachStdin: true,
        Tty:   true}
    containerId, err := docker.CreateContainer(containerConfig, "bird_tomcat")
    if err != nil {
        log.Fatal(err)
    }
	log.Println("container create successful", containerId)
	return containerId
}

func StartContainer(containerID string) error {
	docker := GetDockerClient()
	// Start the container
    hostConfig := &dockerclient.HostConfig{}
	result := make(map[string][]dockerclient.PortBinding)
	var binding dockerclient.PortBinding
	binding.HostIp = "0.0.0.0"
	binding.HostPort = "9999"
	value := make([]dockerclient.PortBinding, 1)
	value[0] = binding
	result["8080/tcp"] = value
	hostConfig.PortBindings = result
    err := docker.StartContainer(containerID, hostConfig)
    if err != nil {
        log.Fatal(err)
		return err
    }
	return nil
}
