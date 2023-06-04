package device_utils

import "strings"

func (location *GPSLocation) Accuracy() int {
	if location.Provider == 0 {
		return randomInt(1, 3)
	}
	return int(location.Provider)
}

func (location *GPSLocation) ProviderString() string {
	provider := location.Provider
	if location.Provider == 0 {
		provider = GPSLocation_LocationProvider(randomInt(1, 3))
	}
	return strings.ToLower(GPSLocation_LocationProvider_name[int32(provider)])
}
