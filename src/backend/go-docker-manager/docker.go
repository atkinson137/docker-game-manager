package main

import (
	"context"
	"io"
	"os"
	"time"

	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"docker.io/go-docker/api/types/container"

	"github.com/hhkbp2/go-logging"
)

func dockerContainer(cli *docker.Client, mainlog *logging.Logger) {
	imageName := "hello-world"
	containerId, reader := bootContainer(cli, imageName, mainlog)

	filePath := "./logs/docker/" + containerId + ".log"
	fileMode := os.O_APPEND
	bufferSize := 0
	bufferFlushTime := 30 * time.Second
	inputChanSize := 1
	// set the maximum size of every file to 100 M bytes
	fileMaxBytes := uint64(100 * 1024 * 1024)
	// keep 9 backup at most(including the current using one,
	// there could be 10 log file at most)
	backupCount := uint32(9)
	// create a handler(which represents a log message destination)
	handler := logging.MustNewRotatingFileHandler(
		filePath, fileMode, bufferSize, bufferFlushTime, inputChanSize,
		fileMaxBytes, backupCount)

	// the format for the whole log message
	format := "%(asctime)s %(levelname)s (%(filename)s:%(lineno)d) " +
		"%(name)s %(message)s"
	// the format for the time part
	dateFormat := "%Y-%m-%d %H:%M:%S.%3n"
	// create a formatter(which controls how log messages are formatted)
	formatter := logging.NewStandardFormatter(format, dateFormat)
	// set formatter for handler
	handler.SetFormatter(formatter)

	// create a logger(which represents a log message source)
	dockerLog := logging.GetLogger(imageName)
	dockerLog.SetLevel(logging.LevelInfo)
	dockerLog.AddHandler(handler)

	infoLog(&reader, &dockerLog)

	// ensure all log messages are flushed to disk before program exits.
	defer logging.Shutdown()
}

func bootContainer(cli *docker.Client, image string, mainlog *logging.Logger) (string, io.Reader) {
	(*mainlog).Infof("Pulling container: %s", image)

	closer, imgPullErr := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})

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

	return body.ID, closer
}
