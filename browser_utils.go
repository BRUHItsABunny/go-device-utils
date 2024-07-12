package device_utils

import (
	"fmt"
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

	order := brandPermutations[majorVersion%len(brandPermutations)]

	// https://source.chromium.org/chromium/chromium/src/+/main:components/embedder_support/user_agent_utils.cc;l=392;drc=2385479e028cfd50420ff8a4406da113d65622c6;bpv=1;bpt=1
	if !legacyAlg {
		grease = fmt.Sprintf("\"Not%sA%sBrand\";v=\"%s\"",
			greasyChars[majorVersion%len(greasyChars)],
			greasyChars[(majorVersion+1)%len(greasyChars)],
			greasedVersions[majorVersion%len(brandPermutations)],
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
		result[order[1]] = fmt.Sprintf("\"Chromium\";v=\"%d\"", majorVersion)
		result[order[2]] = fmt.Sprintf("\"%s\";v=\"%d\"", brand, majorVersion)
	} else {
		result = make([]string, 2)
		result[majorVersion%2] = grease
		result[(majorVersion+1)%2] = fmt.Sprintf("\"Chromium\";v=\"%d\"", majorVersion)
	}

	return strings.Join(result, ", ")
}
