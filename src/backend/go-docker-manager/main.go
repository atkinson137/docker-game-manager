package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"
)

// https://www.melvinvivas.com/learning-go-with-docker-sdk/

func main() {

	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	bootContainer(cli, "hello-world")

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}

func bootContainer(cli *docker.Client, image string) {
	closer, imgPullErr := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	io.Copy(os.Stdout, closer)

	if imgPullErr != nil {
		panic(imgPullErr)
	}

	config := &container.Config{
		Image: image,
	}
	body, err1 := cli.ContainerCreate(context.Background(), config, nil, nil, "")
	if err1 != nil {
		panic(err1)
	}

	err := cli.ContainerStart(context.Background(), body.ID, types.ContainerStartOptions{})
	if err != nil {
		//fmt.Println(err)
		panic(err)
	}
}
