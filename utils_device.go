package device_utils

import (
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
)

func DeviceFromFingerprint(fingerprint string) (*AndroidDevice, error) {
	// "OnePlus/OnePlus5/OnePlus5:9/PKQ1.180716.001/2002242003:user/release-keys"
	var (
		err    error
		device = new(AndroidDevice)
	)
	device.Build = new(AndroidDevice_BuildData)
	device.Screen = new(ScreenData)
	device.Cpu = &CPUData{
		Arch:    CPUData_ARM64,
		AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
	}

	mainParts := strings.Split(fingerprint, "/")
	device.Build.Manufacturer = mainParts[0]
	device.Build.Product = mainParts[1]
	subParts := strings.Split(mainParts[2], ":")
	device.Build.Device = subParts[0]
	device.Version, err = AndroidVersionFromVersionString(subParts[1])
	if err == nil {
		device.Build.Id = mainParts[3]
		subParts = strings.Split(mainParts[4], ":")
		device.Build.IncrementalVersion = subParts[0]
		device.Build.Type = subParts[1]
		device.Build.Tags = mainParts[5]
	}
	return device, err
}

func DeviceFromUserAgent(userAgent string) (*AndroidDevice, error) {
	// Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 5 Build/JOP40D)
	var (
		err    error
		device = new(AndroidDevice)
	)
	device.Build = new(AndroidDevice_BuildData)
	device.Screen = new(ScreenData)
	device.Cpu = &CPUData{
		Arch:    CPUData_ARM64,
		AbiList: []string{"arm64-v8a", "armeabi-v7a", "armeabi"},
	}

	indexStart := strings.Index(userAgent, "(")
	indexStop := strings.Index(userAgent, ")")
	deviceStr := userAgent[indexStart:indexStop]
	for _, elem := range strings.Split(deviceStr, "; ") {
		if strings.Contains(elem, "Android ") {
			device.Version, err = AndroidVersionFromVersionString(strings.Split(elem, " ")[1])
		} else if strings.Contains(elem, "-") {
			device.Locale, err = LocaleFromLocaleString(elem)
		} else if strings.Contains(elem, "Build/") {
			parts := strings.Split(elem, " Build/")
			device.Build.Model = parts[0]
			device.Build.Id = parts[1]
		}

		if err != nil {
			break
		}
	}

	return device, err
}

func (device *AndroidDevice) GetUserAgent() string {
	// Returns the device string part of useragent, manually need to prefix and postfix it with data like what software/browser the user agent is supposed to be
	result := "(Linux; Android " + device.Version.ToAndroidVersion() + "; " + strings.ToLower(device.Locale.ToLocale("-", true)) + "; " + device.Build.Model + " Build/" + device.Build.Id + ")"
	return result
}

const (
	DeviceFormatKeyAndroidVersion  = ":andVers"
	DeviceFormatKeyAndroidSDKLevel = ":andSDK"
	DeviceFormatKeyLocale          = ":locale"
	DeviceFormatKeyModel           = ":model"
	DeviceFormatKeyBuild           = ":build"
	DeviceFormatKeyDPI             = ":dpi"
	DeviceFormatKeyDevice          = ":device"
	DeviceFormatKeyManufacturer    = ":manufacturer"
)

func (device *AndroidDevice) FormatUserAgent(format string) string {
	// TODO: Cache this? replace is inefficient everytime?
	format = strings.ReplaceAll(format, DeviceFormatKeyAndroidVersion, device.Version.ToAndroidVersion())
	format = strings.ReplaceAll(format, DeviceFormatKeyAndroidSDKLevel, device.Version.ToAndroidSDK())
	format = strings.ReplaceAll(format, DeviceFormatKeyLocale, strings.ToLower(device.Locale.ToLocale("-", true)))
	format = strings.ReplaceAll(format, DeviceFormatKeyModel, device.Build.Model)
	format = strings.ReplaceAll(format, DeviceFormatKeyBuild, device.Build.Id)
	format = strings.ReplaceAll(format, DeviceFormatKeyDPI, strconv.Itoa(int(device.Screen.Density)))
	format = strings.ReplaceAll(format, DeviceFormatKeyDevice, device.Build.Device)
	format = strings.ReplaceAll(format, DeviceFormatKeyManufacturer, device.Build.Manufacturer)
	return format
}

func (device *AndroidDevice) GetFingerprint() string {
	result := device.Build.Manufacturer + "/" + device.Build.Product + "/" + device.Build.Device + ":" + device.Version.ToAndroidVersion() + "/" + device.Build.Id + "/" + device.Build.IncrementalVersion + ":" + device.Build.Type + "/" + device.Build.Tags
	return result
}

func (device *AndroidDevice) Randomize() { // I do recommend setting to Locale field of the device though
	// Allow device randomization with existing Device instance - useful if you store devices in database and want to randomize upon retrieval
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
}

func (m *MAC) PrettyFormat(separator string) string {
	macChunks := groupSubString(m.Address, "f", 2)
	return strings.Join(macChunks, separator)
}

func (m *MAC) Generate(oui string, multiCast, uua bool) (string, error) {
	if len(oui) < 1 {
		oui = m.OUI
	}
	oui = strings.Map(removeAllNONHex, strings.ToLower(oui))
	ouiChunks := groupSubString(oui, "f", 2)

	macBytes := make([]byte, 6)

	// Randomization and settings, to be overwritten by prefix if set
	rand.Read(macBytes[:])
	if multiCast {
		macBytes[0] |= 0
	} else {
		macBytes[0] ^= 1
	}
	if uua {
		macBytes[0] ^= 1 << 1
	} else {
		macBytes[0] |= 1 << 1
	}

	var (
		err    error
		hexInt int64
	)
	limit := 6
	if len(ouiChunks) < limit {
		limit = len(ouiChunks)
	}
	for i := 0; i < limit; i++ {
		hexInt, err = strconv.ParseInt(ouiChunks[i], 16, 64)
		if err != nil {
			return "", err
		}
		macBytes[i] = byte(hexInt)
	}

	m.Address = hex.EncodeToString(macBytes)

	return m.Address, nil

}
