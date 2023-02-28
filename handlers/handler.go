package handlers

import (
	"busapp/models"

	"github.com/couchbase/gocb/v2"
)

type handler struct {
	DB *gocb.Cluster
}

func New(db *gocb.Cluster) handler {
	return handler{db}
}

type User models.User
type Brand models.Brand
type Property models.Property
type Type models.Type
type BusModel models.BusModel
type Bus models.Bus
type BusProperty models.BusProperty
type Voyage models.Voyage
type Province models.Province
type BusSeat models.BusSeat
type VoyageResponse models.VoyageResponse
type BusResponse models.BusResponse
