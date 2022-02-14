package utils

import (
	"github.com/google/uuid"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"math/rand"
	"time"
)

/*
	models.go contains the models, in the form of structs, of the main entities that are used in the Demo, and utility
	functions used to deal with them.
*/

///////////////////////////////////////////////PHASE 1
///////////////////////////////////////////////PHASE 1

// The Request is a TD object
type NewInterfaceResponse struct {
	UID          uuid.UUID
	Name         string
	Owner        string
	CreationTime time.Time
}


///////////////////////////////////////////////PHASE 2
///////////////////////////////////////////////PHASE 2

type SensorInstance struct {
	UID 	int// short UID instead of uuid.UUID
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
	InstanceID int// short UID instead of uuid.UUID
	Endpoint string
	CreationTime time.Time
	BoardName	string
	Manufacturer string
}

// NewGeojsonFeature function randomly returns a valid and consistent Geojson feature.
// It is used in the generation of the Area field of InstanceRegistrationRequest
func NewGeojsonFeature() geojson.Feature {
	pol := orb.Polygon{}
	switch rand.Intn(3){
	case 0:
		pol = orb.Polygon{orb.Ring{
			orb.Point{11.534409577930553, 44.570339452069078},
			orb.Point{11.533747072303875, 44.569433736543559},
			orb.Point{11.533323283279522, 44.569598065724215},
			orb.Point{11.534042115295513, 44.5704961356024},
			orb.Point{11.534409577930553, 44.570339452069078}}}
	case 1:
		pol = orb.Polygon{orb.Ring{
			orb.Point{10.764557719230652, 44.70813322227338},
			orb.Point{10.767009258270264, 44.70764522151563},
			orb.Point{10.767889022827148, 44.7095971998612},
			orb.Point{10.768682956695555, 44.7113813724823},
			orb.Point{10.767652988433836, 44.71157961048961},
			orb.Point{10.76686978340149, 44.7097153839366},
			orb.Point{10.765276551246643, 44.70998606268386},
			orb.Point{10.764557719230652, 44.70813322227338}}}
	case 2:
		pol = orb.Polygon{orb.Ring{
			orb.Point{ 10.765142440795898, 44.724295922018655},
			orb.Point{ 10.771965980529785, 44.718105836474514},
			orb.Point{ 10.778017044067383, 44.72972313679335},
			orb.Point{ 10.765142440795898, 44.724295922018655}}}
	}

	return *(geojson.NewFeature(pol))
}

///////////////////////////////////////////////PHASE 3
///////////////////////////////////////////////PHASE 3

// The request is a SenML object
type SamplingResponse struct{
	Status string
}

