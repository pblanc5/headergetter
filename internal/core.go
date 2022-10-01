package internal

import (
	"github.com/dyatlov/go-opengraph/opengraph"
)

type Service interface {
	GetOpenGraphTags(href string) (opengraph.OpenGraph, error)
}
