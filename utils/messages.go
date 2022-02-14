package utils

import (
	"fmt"
	"math/rand"
)


/*
	List of message prototypes used in phase 1 and phase 3. The messages are associated based on the position (index),
	and constitute examples of real messages extrapolated from smart farming contexts.
	The phase 2 messages are generated on the fly.
*/

var TestMessagesTD []ThingDescription = []ThingDescription{
	ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]DataSchema{
			"Air Temperature" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"tem"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Air Humidity" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"hum"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
		},
		Events:       nil,
	},

	ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]DataSchema{
			"Air Temperature" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"tem"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Air humidity" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"hum"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Luminosity" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"lux"},
				MinVal:      0.0,
				MaxVal:      50000.0,
			},
		},
		Events:       nil,
	},

	ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]DataSchema{
			"Soil Water Content Gravimetric in 30 cm" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"swc_g_30"},
				MinVal:      0.0,
				MaxVal:      100.0,
			},
		},
		Events:       nil,
	},

	ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]DataSchema{
			"Soil Temperature at 10 cm" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_tem_10"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Soil Moisture at 10 cm" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_mois_10"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Soil Temperature at 20 cm" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_tem_20"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Soil Moisture at 20 cm" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_mois_20"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Soil Temperature at 50 cm" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_tem_50"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Soil Moisture at 50 cm" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_mois_50"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
		},
		Events:       nil,
	},

	ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]DataSchema{
			"Temperature Celsius" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"tem"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Humidity Relative" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"hum"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Rainfall cm" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"rnf"},
				MinVal:      0.0,
				MaxVal:      +300.0,
			},
			"Wind Speed m/s" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"wnd_sp"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Wind Direction" : DataSchema{
				Type:        "string",
				Description: map[string]string{"name":"wnd_dr"},
			},
			"Air Pressure bar" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"air_pr"},
				MinVal:      0.0,
			},
			"Light Intensity cd" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"lux"},
			},
			"UV Index" : DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"uv"},
				MinVal:      0.0,
				MaxVal:      +10.0,
			},
		},
		Events:       nil,
	},
}

// Using a function to obtain fill the message with random values
func GetSENML(index int) (SenMLPos, error) {
	switch index {
	case 0:
		return SenMLPos{
			TimeRecord: "",
			Name:       "",
			Data: map[string]interface{}{
				"tem": float64(rand.Intn(80) - 20),
				"hum": float64(rand.Intn(100)),
			}}, nil
	case 1:
		return SenMLPos{
				TimeRecord: "",
				Name:       "",
				Data:       map[string]interface{}{
					"tem": float64(rand.Intn(80) - 20),
					"hum" : float64(rand.Intn(100)),
					"lux" : float64(rand.Intn(50000)),
				}}, nil
	case 2:
		return SenMLPos{
			TimeRecord: "",
			Name:       "",
			Data:       map[string]interface{}{
				"swc_g_30": float64(rand.Intn(100)),
			}}, nil
	case 3:
		return SenMLPos{
			TimeRecord: "",
			Name:       "",
			Data:       map[string]interface{}{
				"s_tem_10": float64(rand.Intn(80) - 20),
				"s_mois_10" : float64(rand.Intn(100)),
				"s_tem_20": float64(rand.Intn(80) - 20),
				"s_mois_20" : float64(rand.Intn(100)),
				"s_tem_50": float64(rand.Intn(80) - 20),
				"s_mois_50" : float64(rand.Intn(100)),
			}}, nil
	case 4:
		return SenMLPos{
			TimeRecord: "",
			Name:       "",
			Data:       map[string]interface{}{
				"tem": float64(rand.Intn(80) - 20),
				"hum" : float64(rand.Intn(100)),
				"rnf" : float64(rand.Intn(300)),
				"wnd_sp" : float64(rand.Intn(100)),
				"wnd_dr" : "S/W",
				"air_pr" : float64(rand.Intn(10000)),
				"lux" : float64(rand.Intn(50000)),
				"uv" : float64(rand.Intn(10)),
			}}, nil
		}

	return SenMLPos{}, fmt.Errorf("Invalid index: %d", index)
}



