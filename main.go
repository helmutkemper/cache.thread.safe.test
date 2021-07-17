package main

import (
	docker "github.com/helmutkemper/iotmaker.docker.builder"
	dockerNetwork "github.com/helmutkemper/iotmaker.docker.builder.network"
	"strconv"
)

func createNetwork() (netDocker *dockerNetwork.ContainerBuilderNetwork, err error) {
	netDocker = &dockerNetwork.ContainerBuilderNetwork{}
	err = netDocker.Init()
	if err != nil {
		return
	}

	// create a network named cache_delete_after_test, subnet 10.0.0.0/16 e gatway 10.0.0.1
	err = netDocker.NetworkCreate("cache_delete_after_test", "10.0.0.0/28", "10.0.0.1")
	if err != nil {
		return
	}

	return
}

func createProject() (err error) {

	var container *docker.ContainerBuilder
	container = &docker.ContainerBuilder{}

	// new image name delete:latest
	container.SetImageName("delete:latest")
	// set a folder path to make a new image
	container.SetBuildFolderPath("./server")
	//container.MakeDefaultDockerfileForMe()
	container.SetPrintBuildOnStrOut()
	// inicialize container object
	err = container.Init()
	if err != nil {
		return
	}

	// builder new image from folder
	err = container.ImageBuildFromFolder()
	if err != nil {
		return
	}

	return
}

func buildProject(containerName string, netDocker *dockerNetwork.ContainerBuilderNetwork) (container *docker.ContainerBuilder, err error) {

	container = &docker.ContainerBuilder{}

	container.SetNetworkDocker(netDocker)
	container.SetEnvironmentVar([]string{"DEBUG_NAME=" + containerName, "IP_SERVICE=10.0.0.2", "GATEWAY_SERVICE=10.0.0.0/28"})

	// new image name delete:latest
	container.SetImageName("delete:latest")
	// set a folder path to make a new image
	//container.MakeDefaultDockerfileForMe()
	container.SetPrintBuildOnStrOut()
	// container name container_delete_server_after_test
	container.SetContainerName(containerName)
	// set a waits for the text to appear in the standard container output to proceed [optional]
	//container.SetWaitStringWithTimeout("starting server", 30*time.Second)
	// inicialize container object
	err = container.Init()
	if err != nil {
		return
	}

	// build a new container from image
	err = container.ContainerBuildFromImage()
	if err != nil {
		return
	}

	return
}

func buildCache() (err error) {
	var imageCacheName = "cache:latest"
	var imageId string
	var container = &docker.ContainerBuilder{}

	imageId, err = container.ImageFindIdByName(imageCacheName)
	if err != nil && err.Error() != "image name not found" {
		return
	}

	if imageId != "" {
		return
	}

	// new image name delete:latest
	container.SetImageName(imageCacheName)
	// set a folder path to make a new image
	//container.MakeDefaultDockerfileForMe()
	container.SetPrintBuildOnStrOut()
	// container name container_delete_server_after_test
	container.SetContainerName(imageCacheName)

	// build an image used as cache
	container.SetBuildFolderPath("./cache")

	// inicialize container object
	err = container.Init()
	if err != nil {
		return
	}

	// build a new container from image
	err = container.ImageBuildFromFolder()
	if err != nil {
		return
	}

	return
}

func main() {
	var err error
	var netDocker *dockerNetwork.ContainerBuilderNetwork
	var container *docker.ContainerBuilder
	var idList = make(map[int]*docker.ContainerBuilder)

	docker.GarbageCollector()

	err = buildCache()
	if err != nil {
		panic(err)
	}

	netDocker, err = createNetwork()
	if err != nil {
		panic(err)
	}

	err = createProject()
	if err != nil {
		panic(err)
	}

	for i := 0; i != 2; i += 1 {
		id := strconv.FormatInt(int64(i), 10)
		container, err = buildProject("container_delete_after_test_"+id, netDocker)
		if err != nil {
			panic(err)
		}

		idList[i] = container
	}
}
