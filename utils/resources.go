package utils

import (
	"sync"
	"github.com/samalba/dockerclient"
)

var (
	// Init the client
    docker *dockerclient.DockerClient
	once sync.Once
)

func GetDockerClient() *dockerclient.DockerClient {
	once.Do(func() {
		docker, _ = dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)
	})
	return docker
}
