package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/lucassilveira96/healthcheck/api/utils"
)

func GetAllDockers() string {
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := c.Get(os.Getenv("DOCKER_API_JWT"))
	if err != nil {
		return utils.DOCKER_JWT_STOPPED
	} else {
		_, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return utils.DOCKER_JWT_STOPPED
		}

		return utils.DOCKER_JWT_RUNNING
	}
}
