package utils

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Struct for reading Thing Description to set new instance of sensor node
// Ref: https://www.w3.org/TR/wot-thing-description/

////////////////////////////////////////

// Simply TD message structure
type ThingDescription struct {
	ID           string  `json:"id"`
	Title        string `json:"title"`
	Model        string  `json:"model"`
	Description string `json:"description"`
	Manufacturer string  `json:"manufacturer"`
	Properties map[string]DataSchema            `json:"properties"`
	Events     map[string]InteractionAffordance `json:"events"`
}

type DataSchema struct {
	Type        string            `json:"type"`
	Description map[string]string `json:"description"`
	MinVal      float64          `json:"minimum"`
	MaxVal      float64          `json:"maximum"`
}

type InteractionAffordance struct {
	Forms []Form `json:"forms"`
}

type Form struct {
	Href        string `json:"href"`
	ContentType string `json:"contentType"`
}


////////////////////////////////////////
// Utility functions

func (td *ThingDescription) ToSaveSensor() []byte {

	tdCpy := *td
	tdCpy.ID = ""

	stc, e := json.Marshal(tdCpy)
	if e != nil {
		fmt.Println(e)
	}

	return stc

}

/*
Check if given Thing Description descibe initialized sensor
*/
func (td *ThingDescription) Equals(o ThingDescription) (bool, error) {
	eqProp := true

	if o.Model == "" || td.Model == "" || o.Manufacturer == "" || td.Manufacturer == "" || o.Title == "" || td.Title == "" {
		return false, errors.New("Check Thing Description format. Some values must not be empty.")
	}

	eqProp = eqProp && (len(o.Properties) == len(td.Properties))

	if eqProp {
		count := 0
		for k, _ := range o.Properties {

			if _, ok := td.Properties[k]; ok {
				count += 1
			}
		}

		eqProp = eqProp && (count == len(td.Properties))
	}

	return eqProp && (o.Model == td.Model) && (o.Title == td.Title) && (o.Manufacturer == td.Manufacturer), nil
}
