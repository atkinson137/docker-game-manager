package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"

	"github.com/hhkbp2/go-logging"
)

// https://www.melvinvivas.com/learning-go-with-docker-sdk/

func main() {

	config_file := "./config/logging.yml"
	if err := logging.ApplyConfigFile(config_file); err != nil {
		panic(err.Error())
	}

	mainLog := logging.GetLogger("RNSM")

	//ensure logs arre flushed on shutdown
	defer logging.Shutdown()

	mainLog.Info("Program startup")
	cli, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	dockerContainer(cli, &mainLog)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}
}

// infoLog takes a message and writes it to the given log
func infoLog(message *io.Reader, logger *logging.Logger) {
	scanner := bufio.NewScanner(*message)
	for scanner.Scan() {
		text := scanner.Text()
		(*logger).Infof("Docker message: %s", text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
