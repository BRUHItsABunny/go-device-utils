package device_utils

import (
	"encoding/json"
	"fmt"
	"testing"
)

// TestGenerateBrandHeader src: // https://source.chromium.org/chromium/chromium/src/+/main:components/embedder_support/user_agent_utils_unittest.cc;l=774-776;drc=2385479e028cfd50420ff8a4406da113d65622c6
func TestGenerateBrandHeader(t *testing.T) {
	type testArgs struct {
		brand        string
		majorVersion int
		result       string
	}

	testCases := []testArgs{
		{
			brand:        "Brave",
			majorVersion: 126,
			result:       `"Not/A)Brand";v="8", "Chromium";v="126", "Brave";v="126"`,
		},
	}

	for i, testCase := range testCases {
		result := GenerateBrandHeader(testCase.brand, testCase.majorVersion)
		fmt.Println(fmt.Sprintf("Test %d generated: %s", i, result))
		if result != testCase.result {
			t.Error(fmt.Sprintf("Test %d failed got: %s, want: %s", i, result, testCase.result))
		}
	}
}

func TestGetLatestChrome(t *testing.T) {
	latest, err := GetLatestChromium(0)
	if err != nil {
		t.Fatal(err)
	}

	brandHeader := GenerateBrandHeader("Google Chrome", latest.GetMajorVersion())
	uaHeader := fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", latest.GetUAVersion())

	result := map[string]string{
		"User-Agent": uaHeader,
		"Sec-Ch-Ua":  brandHeader,
	}
	resultBytes, err := json.MarshalIndent(&result, "", "    ")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(resultBytes))
}
