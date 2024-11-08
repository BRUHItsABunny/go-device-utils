package device_utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
)

var (
	// brandPermutations src: https://source.chromium.org/chromium/chromium/src/+/main:components/embedder_support/user_agent_utils.cc;l=329;drc=2385479e028cfd50420ff8a4406da113d65622c6
	brandPermutations = [][3]int{
		{0, 1, 2}, {0, 2, 1}, {1, 0, 2},
		{1, 2, 0}, {2, 0, 1}, {2, 1, 0},
	}
	// greasyChars src: https://source.chromium.org/chromium/chromium/src/+/main:components/embedder_support/user_agent_utils.cc;l=407;drc=2385479e028cfd50420ff8a4406da113d65622c6;bpv=1;bpt=1
	greasyChars = []string{" ", "(", ":", "-", ".", "/", ")", ";", "=", "?", "_"}
	// greasyCharsLegacy src: https://source.chromium.org/chromium/chromium/src/+/main:components/embedder_support/user_agent_utils.cc;l=421;drc=2385479e028cfd50420ff8a4406da113d65622c6
	greasyCharsLegacy = []string{" ", " ", ";"}
	// greasedVersions src: https://source.chromium.org/chromium/chromium/src/+/main:components/embedder_support/user_agent_utils.cc;l=409;drc=2385479e028cfd50420ff8a4406da113d65622c6;bpv=1;bpt=1
	greasedVersions = []string{"8", "99", "24"}
)

// GenerateBrandHeader Ported GREASing from Chromium
func GenerateBrandHeader(brand string, majorVersion int, useLegacy ...bool) string {
	var (
		grease string
		result = make([]string, 3)
	)

	legacyAlg := false
	if len(useLegacy) > 0 {
		legacyAlg = useLegacy[0]
	}

	postFix := ""
	if len(useLegacy) > 1 && useLegacy[1] {
		postFix = ".0.0.0"
	}

	order := brandPermutations[majorVersion%len(brandPermutations)]

	// https://source.chromium.org/chromium/chromium/src/+/main:components/embedder_support/user_agent_utils.cc;l=392;drc=2385479e028cfd50420ff8a4406da113d65622c6;bpv=1;bpt=1
	if !legacyAlg {
		grease = fmt.Sprintf("\"Not%sA%sBrand\";v=\"%s\"",
			greasyChars[majorVersion%len(greasyChars)],
			greasyChars[(majorVersion+1)%len(greasyChars)],
			greasedVersions[majorVersion%len(greasedVersions)],
		)
	} else {
		grease = fmt.Sprintf("\"%sNot%sA%sBrand\";v=\"%s\"",
			greasyCharsLegacy[order[0]],
			greasyCharsLegacy[order[1]],
			greasyCharsLegacy[order[2]],
			greasedVersions[1],
		)
	}

	// https://source.chromium.org/chromium/chromium/src/+/main:components/embedder_support/user_agent_utils.cc;l=315;drc=2385479e028cfd50420ff8a4406da113d65622c6
	if len(brand) > 0 {
		result[order[0]] = grease
		result[order[1]] = fmt.Sprintf("\"Chromium\";v=\"%d%s\"", majorVersion, postFix)
		result[order[2]] = fmt.Sprintf("\"%s\";v=\"%d%s\"", brand, majorVersion, postFix)
	} else {
		result = make([]string, 2)
		result[majorVersion%2] = grease
		result[(majorVersion+1)%2] = fmt.Sprintf("\"Chromium\";v=\"%d%s\"", majorVersion, postFix)
	}

	return strings.Join(result, ", ")
}

const (
	// src: https://versionhistory.googleapis.com/v1/chrome/platforms/
	PlatformWindows   = "win"
	PlatformWindows64 = "win64"
	PlatformIOS       = "ios"
	PlatformAndroid   = "android"
	PlatformMac       = "mac"
	PlatformMacARM64  = "mac_arm64"
	PlatformLinux     = "linux"

	// src: https://versionhistory.googleapis.com/v1/chrome/platforms/win/channels
	ChannelExtended = "extended"
	ChannelStable   = "stable"
	ChannelBeta     = "beta"
	ChannelDev      = "dev"
	ChannelCanary   = "canary"
)

type ChromiumVersionResponse struct {
	Versions      []*ChromiumVersion `json:"versions"`
	NextPageToken string             `json:"nextPageToken"`
}

type ChromiumVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func (v *ChromiumVersion) GetMajorVersion() int {
	majorStr := strings.Split(v.Version, ".")[0]
	majorInt, err := strconv.Atoi(majorStr)
	if err != nil {
		return 0
	}
	return majorInt
}

func (v *ChromiumVersion) GetUAVersion() string {
	majorStr := strings.Split(v.Version, ".")[0]
	return fmt.Sprintf("%s.0.0.0", majorStr)
}

func GetLatestChromium(index int, args ...string) (*ChromiumVersion, error) {
	platform := PlatformWindows
	if len(args) >= 1 {
		platform = args[0]
	}

	channelId := ChannelStable
	if len(args) >= 2 {
		channelId = args[1]
	}

	reqURL := fmt.Sprintf("https://versionhistory.googleapis.com/v1/chrome/platforms/%s/channels/%s/versions", platform, channelId)
	resp, err := http.Get(reqURL)
	if err != nil {
		return nil, fmt.Errorf("http.Get: %w", err)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}

	respParsed := &ChromiumVersionResponse{}
	err = json.Unmarshal(respBody, respParsed)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	unsigned := index
	if index < 0 {
		unsigned = index * -1
		return respParsed.Versions[len(respParsed.Versions)-(unsigned%len(respParsed.Versions))], nil
	} else {
		return respParsed.Versions[unsigned%len(respParsed.Versions)], nil
	}
}

func MustChromiumHeaders(brand string, defaultMajorVersion int, withFullVersions bool) http.Header {
	if brand == "" {
		brand = "Google Chrome"
	}
	if defaultMajorVersion < 1 {
		defaultMajorVersion = 131
	}

	latest, err := GetLatestChromium(0)
	if err != nil {
		latest = &ChromiumVersion{
			Name:    fmt.Sprintf("chrome/platforms/win/channels/stable/versions/%d.0.0.0", defaultMajorVersion),
			Version: fmt.Sprintf("%d.0.0.0", defaultMajorVersion),
		}
	}

	result := http.Header{
		"user-agent":       {fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%s Safari/537.36", latest.GetUAVersion())},
		"sec-ch-ua":        {GenerateBrandHeader(brand, latest.GetMajorVersion())},
		"sec-ch-ua-mobile": {"?0"},
	}

	if withFullVersions {
		result["sec-ch-ua-platform"] = []string{"\"Windows\""}
		result["sec-ch-ua-arch"] = []string{"\"x86\""}
		result["sec-ch-ua-platform-version"] = []string{"\"19.0.0\""}
		result["sec-ch-ua-model"] = []string{""}
		result["sec-ch-ua-full-version-list"] = []string{GenerateBrandHeader(brand, latest.GetMajorVersion())}
	}

	return result
}
