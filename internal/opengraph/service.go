package opengraph

import (
	"net/http"
	"time"

	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/sirupsen/logrus"
)

type OpenGraphService struct {
	Client http.Client
}

func (s *OpenGraphService) GetOpenGraphTags(href string) (opengraph.OpenGraph, error) {
	resp, err := s.Client.Get(href)
	og := opengraph.OpenGraph{}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"time":  time.Now(),
		}).Warn()
		return og, err
	}

	err = og.ProcessHTML(resp.Body)

	return og, err
}
