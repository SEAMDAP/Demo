package utils

import (
	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"time"
)

type NewSensorRes struct {
	UID          uuid.UUID//RandomId
	Name         string
	Owner        string
	CreationTime time.Time
}

type SensorInstance struct {
	UID 	uuid.UUID
	TD_ID	uuid.UUID
	CreationTime time.Time
	OwnerID	int
}

type InstanceRegistrationRequest struct{
	TDID      uuid.UUID
	UserID    int
	BoardName string
	Area      geojson.Feature
	Server ServerData
}
type ServerData struct {
	Url string
	Period int
}

type InstanceRegistrationResponse struct{
	InstanceID uuid.UUID
	Endpoint string
	CreationTime time.Time
	BoardName	string
	Manufacturer string
}

type SamplingResponse struct{
	Status string
}

func NewGeojsonFeature() geojson.Feature{
	pol := orb.Polygon{orb.Ring{
		orb.Point{11.534409577930553, 44.570339452069078},
		orb.Point{11.533747072303875, 44.569433736543559},
		orb.Point{11.533323283279522, 44.569598065724215},
		orb.Point{11.534042115295513, 44.5704961356024},
		orb.Point{11.534409577930553, 44.570339452069078}}}

	return *(geojson.NewFeature(pol))
}