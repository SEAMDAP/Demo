package configs

import (
	"fmt"
	"github.com/gPenzotti/SEAMDAP/utils"
	"math/rand"
)


//TD := utils.ThingDescription{
//ID:           client.TDId.String(), //MA CHI LO PASSA A CHI??
//Title:        "TD_TITLE_EXAMPLE_" + strconv.Itoa(client.Index),
//Model:        "TD_MODEL_EXAMPLE_" + strconv.Itoa(client.Index),
//Description:  "TD_DESC_EXAMPLE_" + strconv.Itoa(client.Index),
//Manufacturer: "UNIPR",
//Properties:   map[string]utils.DataSchema{
//"temperature" : utils.DataSchema{
//Type:        "number",
//Description: map[string]string{"name":"tem"},
//MinVal:      -20.0,
//MaxVal:      +60.0,
//},
//"humidity" : utils.DataSchema{
//Type:        "number",
//Description: map[string]string{"name":"hum"},
//MinVal:      0.0,
//MaxVal:      +100.0,
//},
//},
//Events:       nil,
//}

var TestMessagesTD []utils.ThingDescription = []utils.ThingDescription {
	utils.ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]utils.DataSchema{
			"Air Temperature" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"tem"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Air Humidity" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"hum"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
		},
		Events:       nil,
	},

	utils.ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]utils.DataSchema{
			"Air Temperature" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"tem"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Air humidity" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"hum"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Luminosity" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"lux"},
				MinVal:      0.0,
				MaxVal:      50000.0,
			},
		},
		Events:       nil,
	},

	utils.ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]utils.DataSchema{
			"Soil Water Content Gravimetric in 30 cm" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"swc_g_30"},
				MinVal:      0.0,
				MaxVal:      100.0,
			},
		},
		Events:       nil,
	},

	utils.ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]utils.DataSchema{
			"Soil Temperature at 10 cm" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_tem_10"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Soil Moisture at 10 cm" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_mois_10"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Soil Temperature at 20 cm" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_tem_20"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Soil Moisture at 20 cm" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_mois_20"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Soil Temperature at 50 cm" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_tem_50"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Soil Moisture at 50 cm" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"s_mois_50"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
		},
		Events:       nil,
	},

	utils.ThingDescription{
		ID:           "",
		Title:        "",
		Model:        "",
		Description:  "",
		Manufacturer: "UNIPR",
		Properties:   map[string]utils.DataSchema{
			"Temperature Celsius" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"tem"},
				MinVal:      -20.0,
				MaxVal:      +60.0,
			},
			"Humidity Relative" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"hum"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Rainfall cm" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"rnf"},
				MinVal:      0.0,
				MaxVal:      +300.0,
			},
			"Wind Speed m/s" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"wnd_sp"},
				MinVal:      0.0,
				MaxVal:      +100.0,
			},
			"Wind Direction" : utils.DataSchema{
				Type:        "string",
				Description: map[string]string{"name":"wnd_dr"},
			},
			"Air Pressure bar" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"air_pr"},
				MinVal:      0.0,
			},
			"Light Intensity cd" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"lux"},
			},
			"UV Index" : utils.DataSchema{
				Type:        "number",
				Description: map[string]string{"name":"uv"},
				MinVal:      0.0,
				MaxVal:      +10.0,
			},
		},
		Events:       nil,
	},
}






func GetSENML(index int) (utils.SenMLPos, error) {
	switch index {
	case 0:
		return utils.SenMLPos{
			TimeRecord: "",
			Name:       "",
			Data: map[string]interface{}{
				"tem": float64(rand.Intn(80) - 20),
				"hum": float64(rand.Intn(100)),
			}}, nil
	case 1:
		return utils.SenMLPos{
				TimeRecord: "",
				Name:       "",
				Data:       map[string]interface{}{
					"tem": float64(rand.Intn(80) - 20),
					"hum" : float64(rand.Intn(100)),
					"lux" : float64(rand.Intn(50000)),
				}}, nil
	case 2:
		return utils.SenMLPos{
			TimeRecord: "",
			Name:       "",
			Data:       map[string]interface{}{
				"swc_g_30": float64(rand.Intn(100)),
			}}, nil
	case 3:
		return utils.SenMLPos{
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
		return utils.SenMLPos{
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

	return utils.SenMLPos{}, fmt.Errorf("Invalid index: %d", index)
}




