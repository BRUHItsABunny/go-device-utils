package go_device_utils

import (
	"encoding/base64"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"testing"
	"time"
)

func TestDeviceFromFingerprint(t *testing.T) {
	fingerPrint := "OnePlus/OnePlus5/OnePlus5:9/PKQ1.180716.001/2002242003:user/release-keys"
	device, err := DeviceFromFingerprint(fingerPrint)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(spew.Sdump(device))
}

func TestDeviceFromUserAgent(t *testing.T) {
	userAgent := "Mozilla/5.0 (Linux; Android 4.2.1; en-us; Nexus 5 Build/JOP40D)"
	device, err := DeviceFromUserAgent(userAgent)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(spew.Sdump(device))
}

func TestRandomDevices(t *testing.T) {
	for i := 0; i < 10; i++ {
		device := GetRandomDevice()
		device.MacAddress.Generate("", false, true)
		for _, sim := range device.SimSlots {
			if sim.Imei != nil {
				sim.Imei.Generate("", "")
			}
		}
		fmt.Println(spew.Sdump(device))
		time.Sleep(5 * time.Second)
	}
}

func TestMACGeneration(t *testing.T) {
	// When looking up the result of this MAC it should give us "OnePlus Electronics (Shenzhen) Co., Ltd." for OUI "A091A2"
	mac := &MAC{
		OUI:     "A091A2",
		Address: "",
	}
	fmt.Println(mac.Generate("", true, true))
	fmt.Println(mac.PrettyFormat(":"))
}

func MustInt(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func TestParseJA3(t *testing.T) {
	// proto version, cipher suites in order of priority, TLS extensions, elliptic curves in order or priority, 0
	ja3 := "771,4865-4866-4867-49195-49199-49196-49200-52393-52392-49171-49172-156-157-47-53,27-16-35-11-17513-43-13-5-23-0-18-51-10-65281-45-21,29-23-24,0"
	fingerprint, _ := ParseTLSFingerprint(ja3)
	fmt.Println(spew.Sdump(fingerprint))

	ja3Again := fingerprint.FormatTLSFingerprint(true)

	if ja3 != ja3Again {
		t.Error("fingerprints not matching")
	}
	fmt.Println(ja3)
	fmt.Println(ja3Again)
	fmt.Println(fmt.Sprintf("%#v", fingerprint))

}

func TestDeserializeTSBytes(t *testing.T) {
	tsBytes := "CggxLjUwLjExNBIFYnJhdmUab01vemlsbGEvNS4wIChXaW5kb3dzIE5UIDEwLjA7IFdpbjY0OyB4NjQpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS8xMTIuMC4wLjAgU2FmYXJpLzUzNy4zNiI5IkNocm9taXVtIjt2PSIxMTIiLCAiQnJhdmUiO3Y9IjExMiIsICJOb3Q6QS1CcmFuZCI7dj0iOTkiKkcIgwYSJIEmgiaDJquAA6+AA6yAA7CAA6mZA6iZA5OAA5SAA5wBnQEvNRoUGxAjC+mIASsNBRcAEjMKgf4DLRUiAx0XGCoBAA=="
	result := &Browser{}

	decoded, err := base64.StdEncoding.DecodeString(tsBytes)
	if err != nil {
		t.Error(err)
	}
	err = result.UnmarshalVT(decoded)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(spew.Sdump(result))
}
