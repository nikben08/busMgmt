package database

import (
	"busapp/seeds"
	"fmt"

	"github.com/couchbase/gocb/v2"
)

func Init() *gocb.Cluster {
	opts := gocb.ClusterOptions{
		Username: "nik",
		Password: "08112001",
	}

	cluster, err := gocb.Connect("couchbase-server", opts)
	if err != nil {
		fmt.Println("connection error")
		panic(err)
	}

	_, err = cluster.Bucket("busmgmt").Collection("brands").Get("2e7a0240-8168-11ed-a1eb-0242ac120002", nil)

	if err != nil {
		users_col := cluster.Bucket("busmgmt").Collection("users")
		user := seeds.SuperUser()
		users_col.Insert(user.ID, user, nil)

		brands_col := cluster.Bucket("busmgmt").Collection("brands")
		brands := seeds.Brands()
		for i := 0; i < len(brands); i++ {
			brands_col.Insert(brands[i].ID, brands[i], nil)
		}

		properties_col := cluster.Bucket("busmgmt").Collection("properties")
		properties := seeds.Properties()
		for i := 0; i < len(properties); i++ {
			properties_col.Insert(properties[i].ID, properties[i], nil)
		}

		types_col := cluster.Bucket("busmgmt").Collection("types")
		types := seeds.Types()
		for i := 0; i < len(types); i++ {
			types_col.Insert(types[i].ID, types[i], nil)
		}

		bus_models_col := cluster.Bucket("busmgmt").Collection("bus_models")
		bus_models := seeds.BusModels()
		for i := 0; i < len(bus_models); i++ {
			bus_models_col.Insert(bus_models[i].ID, bus_models[i], nil)
		}

		provinces_col := cluster.Bucket("busmgmt").Collection("provinces")
		provinces := seeds.Provinces()
		for i := 0; i < len(provinces); i++ {
			provinces_col.Insert(provinces[i].ID, provinces[i], nil)
		}
	}

	return cluster
}
