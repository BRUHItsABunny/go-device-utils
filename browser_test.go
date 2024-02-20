package device_utils

import (
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"io"
	"os"
	"testing"
)

func TestBrowser_FromPEET(t *testing.T) {
	browser := AvailableBrowsers["brave"]["1.50.114"]

	f, err := os.Open("./_resources/samples/peet_brave_120.json")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	err = browser.FromPEETRaw(data)
	if err != nil {
		t.Fatal(err)
	}
	newData, err := protojson.Marshal(browser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(newData))
}

func TestBrowser_FromDLFingerprint(t *testing.T) {
	browser := AvailableBrowsers["brave"]["1.50.114"]

	f, err := os.Open("./_resources/samples/fingerprint_brave_120.json")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	err = browser.FromDLFingerprintRaw(data)
	if err != nil {
		t.Fatal(err)
	}
	newData, err := protojson.Marshal(browser)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(newData))
	os.WriteFile("test.json", newData, 0666)
}

func TestBrowser_FromBoth(t *testing.T) {
	browser := AvailableBrowsers["brave"]["1.50.114"]

	f, err := os.Open("./_resources/samples/peet_brave_120.json")
	if err != nil {
		t.Fatal(err)
	}
	dataPeet, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	f, err = os.Open("./_resources/samples/fingerprint_brave_120.json")
	if err != nil {
		t.Fatal(err)
	}
	dataDLFP, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	err = browser.FromPEETRaw(dataPeet)
	if err != nil {
		t.Fatal(err)
	}
	err = browser.FromDLFingerprintRaw(dataDLFP)
	if err != nil {
		t.Fatal(err)
	}
	newData, err := protojson.Marshal(browser)
	if err != nil {
		t.Fatal(err)
	}
	os.WriteFile("test_full.json", newData, 0666)
}
