package seeds

import (
	"busapp/models"

	guuid "github.com/google/uuid"
)

type Brand models.Brand
type Property models.Property
type Type models.Type
type BusModel models.BusModel
type Province models.Province
type User models.User

func SuperUser() User {
	var user = User{
		ID:          guuid.New().String(),
		Username:    "admin",
		Hash:        "CPRtk51_9-ux353hzCRhNceoaUyr8OHBzGfFDhKDLwg=",
		AccessLevel: "1",
	}
	return user
}

func Brands() []Brand {
	var brands = []Brand{
		Brand{
			ID:   "2e7a0240-8168-11ed-a1eb-0242ac120002",
			Name: "Mercedes-Benz",
		},
		Brand{
			ID:   "2e7a06d2-8168-11ed-a1eb-0242ac120002",
			Name: "Volvo",
		},
		Brand{
			ID:   "2e7a083a-8168-11ed-a1eb-0242ac120002",
			Name: "Toyota",
		},
		Brand{
			ID:   "2e7a0a60-8168-11ed-a1eb-0242ac120002",
			Name: "Scania",
		},
	}
	return brands
}

func Types() []Type {
	var types = []Type{
		Type{
			ID:   guuid.New().String(),
			Name: "2+1",
		},
		Type{
			ID:   guuid.New().String(),
			Name: "2+2",
		},
	}
	return types
}

func Properties() []Property {
	var properties = []Property{
		Property{
			ID:   guuid.New().String(),
			Name: "Rahat",
		},
		Property{
			ID:   guuid.New().String(),
			Name: "Internet",
		},
		Property{
			ID:   guuid.New().String(),
			Name: "220V Elektrik",
		},
		Property{
			ID:   guuid.New().String(),
			Name: "Klima",
		},
		Property{
			ID:   guuid.New().String(),
			Name: "Okuma Lambasi",
		},
		Property{
			ID:   guuid.New().String(),
			Name: "Ikram",
		},
	}
	return properties
}

func BusModels() []BusModel {
	var bus_models = []BusModel{
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Mercedes-Benz A",
			BrandId: "2e7a0240-8168-11ed-a1eb-0242ac120002",
		},
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Mercedes-Benz B",
			BrandId: "2e7a0240-8168-11ed-a1eb-0242ac120002",
		},
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Mercedes-Benz C",
			BrandId: "2e7a0240-8168-11ed-a1eb-0242ac120002",
		},
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Volvo A",
			BrandId: "2e7a06d2-8168-11ed-a1eb-0242ac120002",
		},
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Volvo B",
			BrandId: "2e7a06d2-8168-11ed-a1eb-0242ac120002",
		},
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Toyota A",
			BrandId: "2e7a083a-8168-11ed-a1eb-0242ac120002",
		},
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Toyota B",
			BrandId: "2e7a083a-8168-11ed-a1eb-0242ac120002",
		},
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Scania A",
			BrandId: "2e7a0a60-8168-11ed-a1eb-0242ac120002",
		},
		BusModel{
			ID:      guuid.New().String(),
			Name:    "Scania B",
			BrandId: "2e7a0a60-8168-11ed-a1eb-0242ac120002",
		},
	}
	return bus_models
}

func Provinces() []Province {
	var provinces = []Province{
		Province{
			ID:   guuid.New().String(),
			Name: "Istanbul",
		},
		Province{
			ID:   guuid.New().String(),
			Name: "Ankara",
		},
		Province{
			ID:   guuid.New().String(),
			Name: "Izmir",
		},
		Province{
			ID:   guuid.New().String(),
			Name: "Bursa",
		},
		Province{
			ID:   guuid.New().String(),
			Name: "Adana",
		},
		Province{
			ID:   guuid.New().String(),
			Name: "Konya",
		},
		Province{
			ID:   guuid.New().String(),
			Name: "Samsun",
		},
	}
	return provinces
}
