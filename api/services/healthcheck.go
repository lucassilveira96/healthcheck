package services

import (
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/lucassilveira96/healthcheck/api/models"
	"github.com/lucassilveira96/healthcheck/api/utils"
)

func GetAllDockers() models.Docker {
	c := &http.Client{
		Timeout: utils.TIME_DELAY * time.Second,
	}
	req, err := c.Get(os.Getenv("DOCKER_API_JWT"))
	if err != nil {
		return models.Docker{
			Name:      utils.DOCKER_JWT_STOPPED,
			Status:    false,
			CheckedAt: time.Now(),
		}
	} else {
		_, err := ioutil.ReadAll(req.Body)
		if err != nil {
			return models.Docker{
				Name:      utils.DOCKER_JWT_STOPPED,
				Status:    false,
				CheckedAt: time.Now(),
			}
		}

		return models.Docker{
			Name:      utils.DOCKER_JWT_RUNNING,
			Status:    true,
			CheckedAt: time.Now(),
		}
	}
}
