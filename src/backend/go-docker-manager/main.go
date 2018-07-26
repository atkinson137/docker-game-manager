package main

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"git.apache.org/thrift.git/lib/go/thrift"

	"github.com/hhkbp2/go-logging"
)

// https://www.melvinvivas.com/learning-go-with-docker-sdk/

func main() {

	var containers []dockerContainer

	config_file := "./config/logging.yml"
	if err := logging.ApplyConfigFile(config_file); err != nil {
		panic(err.Error())
	}

	mainLog := logging.GetLogger("RNSM")

	//ensure logs are flushed on shutdown
	defer logging.Shutdown()

	mainLog.Info("Program startup")

	cli := getCli()

	cont := makeContainer(cli, "hello-world", &mainLog)

	containers = append(containers, *cont)

	containerList := getContainers(cli)

	for _, container := range containerList {
		str := fmt.Sprintf("%s %s\n", container.ID[:10], container.Image)

		mainLog.Info(str)
	}

	//mainLog.Info("Program end")

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	addr := "localhost:9090"
	secure := false //TODO: secure

	if err := runServer(transportFactory, protocolFactory, addr, secure); err != nil {
		mainLog.Errorf("Error running server: ", err)
	}
}

// infoLog takes a message and writes it to the given log
func infoLog(message io.Reader, logger *logging.Logger) {
	scanner := bufio.NewScanner(message)
	for scanner.Scan() {
		text := scanner.Text()
		(*logger).Info(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
