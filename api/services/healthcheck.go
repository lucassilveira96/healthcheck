package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/lucassilveira96/healthcheck/api/models"
	"github.com/lucassilveira96/healthcheck/api/utils"
)

func GetAllDockers() []models.Docker {
	c := &http.Client{
		Timeout: utils.TIME_DELAY * time.Second,
	}

	time.Local, _ = time.LoadLocation("America/Sao_Paulo")

	var dockers []models.Docker

	req, err := c.Get(os.Getenv("DOCKER_API_JWT"))
	if err != nil {
		docker := new(models.Docker)
		docker.Name = utils.DOCKER_JWT_STOPPED
		docker.Running = false
		docker.CheckedAt = time.Now().Format(("02/Jan/2006 15:04:05 "))
		dockers = append(dockers, *docker)
	} else {
		_, err := ioutil.ReadAll(req.Body)
		if err != nil {
			docker := new(models.Docker)
			docker.Name = utils.DOCKER_JWT_STOPPED
			docker.Running = false
			docker.CheckedAt = time.Now()
			dockers = append(dockers, *docker)
		}
		docker := new(models.Docker)
		docker.Name = utils.DOCKER_JWT_RUNNING
		docker.Running = true
		docker.CheckedAt = time.Now()
		dockers = append(dockers, *docker)
	}

	req, err = c.Get(os.Getenv("DOCKER_API_JWT"))
	if err != nil {
		docker := new(models.Docker)
		docker.Name = utils.DOCKER_JWT_STOPPED
		docker.Running = false
		docker.CheckedAt = time.Now()
		dockers = append(dockers, *docker)
	} else {
		_, err := ioutil.ReadAll(req.Body)
		if err != nil {
			docker := new(models.Docker)
			docker.Name = utils.DOCKER_JWT_STOPPED
			docker.Running = false
			docker.CheckedAt = time.Now()
			dockers = append(dockers, *docker)
		}
		docker := new(models.Docker)
		docker.Name = utils.DOCKER_JWT_RUNNING
		docker.Running = true
		docker.CheckedAt = time.Now()
		dockers = append(dockers, *docker)
	}

	req, err = c.Get(os.Getenv("DOCKER_API_JWT"))
	if err != nil {
		docker := new(models.Docker)
		docker.Name = utils.DOCKER_JWT_STOPPED
		docker.Running = false
		docker.CheckedAt = time.Now()
		dockers = append(dockers, *docker)
	} else {
		_, err := ioutil.ReadAll(req.Body)
		if err != nil {
			docker := new(models.Docker)
			docker.Name = utils.DOCKER_JWT_STOPPED
			docker.Running = false
			docker.CheckedAt = time.Now()
			dockers = append(dockers, *docker)
		}
		docker := new(models.Docker)
		docker.Name = utils.DOCKER_JWT_RUNNING
		docker.Running = true
		docker.CheckedAt = time.Now()
		dockers = append(dockers, *docker)
	}

	return dockers
}
