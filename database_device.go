package go_device_utils

import (
	"google.golang.org/protobuf/proto"
	"math/rand"
	"time"
)

// Here we store a few devices and way to get them, just easy access in case you want to prototype a few devices in a library fast
// TODO: Actually add a few devices here...
// TODO: Turn this into an interface to customize with custom backend
var DeviceDB = map[string]*AndroidDevice{
	// "oneplus3": "",
	"oneplus5": {
		Locale: &Locale{
			Language:   "en",
			CountryISO: "US",
		},
		Version: AndroidDevice_V9_0,
		Build: &AndroidDevice_BuildData{
			Device:             "OnePlus5",
			Manufacturer:       "OnePlus",
			Model:              "ONEPLUS A5000",
			Product:            "OnePlus5",
			Id:                 "PKQ1.180716.001",
			Type:               "user",
			Tags:               "release-keys",
			IncrementalVersion: "2002242003",
		},
		Screen: &ScreenData{
			Density:              420,
			ResolutionHorizontal: 1080,
			ResolutionVertical:   1920,
		},
		SimSlots: []*SIMCard{
			{
				Imei: &SIMCard_IMEI{
					TAC: "86463003",
				},
			},
			{
				Imei: &SIMCard_IMEI{
					TAC: "86463003",
				},
			},
		},
		MacAddress: &MAC{
			OUI:     "A091A2",
			Address: "",
		},
		Cpu: &CPUData{
			Arch:    CPUData_ARM64,
			AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
		},
	},
	"oneplus7t": {
		Locale: &Locale{
			Language:   "en",
			CountryISO: "US",
		},
		Version: AndroidDevice_V10_0,
		Build: &AndroidDevice_BuildData{
			Device:             "OnePlus7T",
			Manufacturer:       "OnePlus",
			Model:              "HD1905",
			Product:            "OnePlus7T",
			Id:                 "QKQ1.190716.003",
			Type:               "user",
			Tags:               "release-keys",
			IncrementalVersion: "2101212100",
		},
		Screen: &ScreenData{
			Density:              420,
			ResolutionHorizontal: 1080,
			ResolutionVertical:   2400,
		},
		SimSlots: []*SIMCard{
			{
				Imei: &SIMCard_IMEI{
					TAC: "86789104",
				},
			},
			{
				Imei: &SIMCard_IMEI{
					TAC: "86789104",
				},
			},
		},
		MacAddress: &MAC{
			OUI:     "A091A2",
			Address: "",
		},
		Cpu: &CPUData{
			Arch:    CPUData_ARM64,
			AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
		},
	},
	// "oneplus9": "",
	"oneplus9pro": {
		Id: nil,
		Locale: &Locale{
			Language:   "en",
			CountryISO: "US",
		},
		Version: AndroidDevice_V11_0,
		Build: &AndroidDevice_BuildData{
			Device:             "OnePlus9Pro",
			Manufacturer:       "OnePlus",
			Model:              "LE2125",
			Product:            "OnePlus9Pro",
			Id:                 "RKQ1.201105.002",
			Type:               "user",
			Tags:               "release-keys",
			IncrementalVersion: "2105290043",
		},
		Screen: &ScreenData{
			Density:              600,
			ResolutionHorizontal: 1440,
			ResolutionVertical:   3216,
		},
		SimSlots: []*SIMCard{
			{
				Imei: &SIMCard_IMEI{
					TAC: "86381505",
				},
			},
		},
		MacAddress: &MAC{
			OUI:     "A091A2",
			Address: "",
		},
		Cpu: &CPUData{
			Arch:    CPUData_ARM64,
			AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
		},
	},
}

var DeviceDBKeys = []string{
	// "oneplus3",
	"oneplus5",
	"oneplus7t",
	// "oneplus9",
	"oneplus9pro",
}

func GetDBDevice(key string) (*AndroidDevice, bool) {
	device := new(AndroidDevice)
	device.Build = new(AndroidDevice_BuildData)
	device.Screen = new(ScreenData)
	device.Cpu = &CPUData{
		Arch:    CPUData_ARM64,
		AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
	}

	val, found := DeviceDB[key]
	if found {
		device = proto.Clone(val).(*AndroidDevice)
	}
	// Device from DB needs to be random ID
	device.Id = NewAndroidID()
	device.Location = GetRandomDBLocation(device.Locale.GetCountryISO())
	if device.SimSlots == nil || len(device.SimSlots) == 0 {
		device.SimSlots = []*SIMCard{GetRandomDBSIMCard(device.Locale.GetCountryISO())}
	}
	for _, sim := range device.SimSlots {
		sim.Randomize(device.Locale.GetCountryISO())
		if sim.Imei == nil {
			sim.Imei = &SIMCard_IMEI{}
		}
		sim.Imei.Generate("", "")
	}
	if device.MacAddress == nil {
		device.MacAddress = new(MAC)
	}
	device.MacAddress.Generate("", false, true)
	return device, found
}

func GetRandomDevice() *AndroidDevice {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	val := DeviceDB[DeviceDBKeys[r.Intn(len(DeviceDBKeys))]]
	device := proto.Clone(val).(*AndroidDevice)
	// Device from DB needs to be random ID
	device.Id = NewAndroidID()
	device.Location = GetRandomDBLocation(device.Locale.GetCountryISO())
	if device.SimSlots == nil || len(device.SimSlots) == 0 {
		device.SimSlots = []*SIMCard{GetRandomDBSIMCard(device.Locale.GetCountryISO())}
	}
	for _, sim := range device.SimSlots {
		sim.Randomize(device.Locale.GetCountryISO())
		if sim.Imei == nil {
			sim.Imei = &SIMCard_IMEI{}
		}
		sim.Imei.Generate("", "")
	}

	if device.MacAddress == nil {
		device.MacAddress = new(MAC)
	}
	device.MacAddress.Generate("", false, true)
	return device
}
