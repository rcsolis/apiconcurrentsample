package internal

import "encoding/xml"

type City struct {
	XMLName xml.Name `xml:"city"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
}
type Coordinates struct {
	XMLName xml.Name `xml:"coord"`
	Lon     float64  `xml:"lon,attr"`
	Lat     float64  `xml:"lat,attr"`
}
type Temperature struct {
	XMLName xml.Name `xml:"temperature"`
	Value   float64  `xml:"value,attr"`
	Min     float64  `xml:"min,attr"`
	Max     float64  `xml:"max,attr"`
	Unit    string   `xml:"unit,attr"`
}

type Weather struct {
	XMLName     xml.Name    `xml:"current"`
	City        City        `xml:"city"`
	Coordinates Coordinates `xml:"coord"`
	Country     string      `xml:"country,attr"`
	Temperature Temperature `xml:"temperature"`
	Humidity    string      `xml:"humidity,attr"`
	Pressure    string      `xml:"pressure,attr"`
	Wind        struct {
		Speed     string `xml:"speed,attr"`
		Gusts     string `xml:"gusts,attr"`
		Direction struct {
			Value string `xml:"value,attr"`
			Code  string `xml:"code,attr"`
			Name  string `xml:"name,attr"`
		} `xml:"direction"`
	} `xml:"wind"`
	Clouds struct {
		Value string `xml:"value,attr"`
		Name  string `xml:"name,attr"`
	} `xml:"clouds"`
	Visibility    string `xml:"visibility,attr"`
	Precipitation struct {
		Value string `xml:"value,attr"`
		Mode  string `xml:"mode,attr"`
		Unit  string `xml:"unit,attr"`
	} `xml:"precipitation"`
	WeatherDetails struct {
		Number string `xml:"number,attr"`
		Value  string `xml:"value,attr"`
		Icon   string `xml:"icon,attr"`
	} `xml:"weather"`
	LastUpdate string `xml:"lastupdate,attr"`
}
