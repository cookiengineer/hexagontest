package types

import "strconv"
import "strings"

type Geolocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewGeolocation() Geolocation {

	var geolocation Geolocation

	// Antarctica
	geolocation.Latitude = -90.0
	geolocation.Longitude = 0.0

	return geolocation

}

func ParseGeolocation(value string) *Geolocation {

	var result *Geolocation = nil

	if strings.Contains(value, ",") {

		tmp1 := strings.Split(value, ",")

		if len(tmp1) == 2 {

			num1, err1 := strconv.ParseFloat(tmp1[0], 64)
			num2, err2 := strconv.ParseFloat(tmp1[1], 64)

			if err1 == nil && err2 == nil {

				geolocation := Geolocation{
					Latitude:  num1,
					Longitude: num2,
				}

				result = &geolocation

			}

		}

	}

	return result

}

func ToGeolocation(latitude float64, longitude float64) Geolocation {

	var geolocation Geolocation

	geolocation.Latitude = latitude
	geolocation.Longitude = longitude

	return geolocation

}

func (geolocation Geolocation) String() string {

	lat_encoded := strconv.FormatFloat(geolocation.Latitude, 'g', -1, 64)
	long_encoded := strconv.FormatFloat(geolocation.Longitude, 'g', -1, 64)

	return lat_encoded + "," + long_encoded

}
