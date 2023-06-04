package device_utils

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"strings"
)

var LocationDB = map[string]map[string]*GPSLocation{
	"US": {
		"newyorkcity": &GPSLocation{
			Longitude: -74.005973,
			Latitude:  40.712775,
			Altitude:  10.440,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
		"losangeles": &GPSLocation{
			Longitude: -118.243685,
			Latitude:  34.052234,
			Altitude:  86.854,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
		"chicago": &GPSLocation{
			Longitude: -87.629798,
			Latitude:  41.878114,
			Altitude:  181.513,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
		"houston": &GPSLocation{
			Longitude: -95.369803,
			Latitude:  29.760427,
			Altitude:  14.562,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
		"washington": &GPSLocation{
			Longitude: -77.036871,
			Latitude:  38.907192,
			Altitude:  22.015,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
		"philadelphia": &GPSLocation{
			Longitude: -75.165222,
			Latitude:  39.952584,
			Altitude:  14.336,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
		"miami": &GPSLocation{
			Longitude: -80.191790,
			Latitude:  25.761680,
			Altitude:  0.537,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
	},
	"MX": {
		"mexicocity": &GPSLocation{
			Longitude: -99.133208,
			Latitude:  19.432608,
			Altitude:  2229.729,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
	},
	"CA": {
		"toronto": &GPSLocation{
			Longitude: -79.383184,
			Latitude:  43.653226,
			Altitude:  91.723,
			Provider:  GPSLocation_LocationProvider_NONE,
		},
	},
}

var AvailableCountries = []string{
	"US", "MX", "CA",
}

var AvailableCities = map[string][]string{
	"US": {
		"newyorkcity",
		"losangeles",
		"chicago",
		"houston",
		"washington",
		"philadelphia",
		"miami",
	},
	"MX": {
		"mexicocity",
	},
	"CA": {
		"toronto",
	},
}

func GetDBLocation(countryISO, city string) (*GPSLocation, error) {
	countryISO = strings.ToUpper(countryISO)
	city = strings.ReplaceAll(strings.ToLower(city), " ", "")
	result := new(GPSLocation)
	var err error

	_, ok := AvailableCities[countryISO]
	if ok {
		result, ok = LocationDB[countryISO][city]
		if !ok {
			err = errors.New("city not supported")
		}
	} else {
		err = errors.New("country not supported")
	}

	return result, err
}

func GetRandomDBLocation(countryISO string) *GPSLocation {
	_, ok := AvailableCities[countryISO]
	if !ok {
		countryISO = randomStrSlice(AvailableCountries)
	}
	city := randomStrSlice(AvailableCities[countryISO])
	location := LocationDB[countryISO][city]

	return proto.Clone(location).(*GPSLocation)
}
