package go_device_utils

import (
	"math/rand"
	"sort"
	"strconv"
	"strings"
)

func mustInt(input string) int {
	res, _ := strconv.Atoi(input)
	return res
}

func mustString(input int) string {
	return strconv.Itoa(input)
}

func ParseTLSFingerprint(fingerprint string) (*Browser_TLSFingerprint, error) {
	result := &Browser_TLSFingerprint{}
	ja3Split := strings.Split(fingerprint, ",")
	for i, content := range ja3Split {
		elements := strings.Split(content, "-")
		switch i {
		case 0:
			result.Version = Browser_TLSFingerprint_ProtocolVersion(mustInt(content))
			break
		case 1:
			result.CipherSuites = make([]Browser_TLSFingerprint_CipherSuite, len(elements))
			for elementI, element := range elements {
				result.CipherSuites[elementI] = Browser_TLSFingerprint_CipherSuite(mustInt(element))
			}
			break
		case 2:
			result.Extensions = make([]Browser_TLSFingerprint_Extensions, len(elements))
			for elementI, element := range elements {
				result.Extensions[elementI] = Browser_TLSFingerprint_Extensions(mustInt(element))
			}
			break
		case 3:
			result.EllipticCurves = make([]Browser_TLSFingerprint_EllipticCurve, len(elements))
			for elementI, element := range elements {
				result.EllipticCurves[elementI] = Browser_TLSFingerprint_EllipticCurve(mustInt(element))
			}
			break
		case 4:
			result.EllipticCurvePointFormats = make([]Browser_TLSFingerprint_EllipticCurvePointFormat, len(elements))
			for elementI, element := range elements {
				result.EllipticCurvePointFormats[elementI] = Browser_TLSFingerprint_EllipticCurvePointFormat(mustInt(element))
			}
			break
		}
	}
	return result, nil
}

func (fp *Browser_TLSFingerprint) FormatTLSFingerprint(strict ...bool) string {
	fmtStrict := false
	if len(strict) > 0 {
		fmtStrict = strict[0]
	}

	extensions := make([]Browser_TLSFingerprint_Extensions, len(fp.Extensions))
	copy(extensions, fp.Extensions)
	if !fmtStrict {
		rand.Shuffle(len(extensions), func(i, j int) {
			extensions[i], extensions[j] = extensions[j], extensions[i]
		})
	}

	result := []string{
		mustString(int(fp.Version)),
		"",
		"",
		"",
		"",
	}

	cipherSuites := make([]string, len(fp.CipherSuites))
	for cipherSuiteI, cipherSuite := range fp.CipherSuites {
		cipherSuites[cipherSuiteI] = mustString(int(cipherSuite))
	}
	result[1] = strings.Join(cipherSuites, "-")

	// Can be randomized except for the ending, needs to end on 21 or 41 or both (21-41)
	extensionsTmp := make([]string, len(extensions))
	extensionsEnd := []int{}
	for extensionI, extension := range extensions {
		if extension == 21 || extension == 41 {
			extensionsEnd = append(extensionsEnd, int(extension))
		} else {
			extensionsTmp[extensionI-len(extensionsEnd)] = mustString(int(extension))
		}
	}
	sort.Ints(extensionsEnd)
	for extensionEndI, extensionEnd := range extensionsEnd {
		extensionsTmp[len(extensionsTmp)-len(extensionsEnd)+extensionEndI] = mustString(extensionEnd)
	}

	result[2] = strings.Join(extensionsTmp, "-")

	ellipticCurves := make([]string, len(fp.EllipticCurves))
	for ellipticCurveI, ellipticCurve := range fp.EllipticCurves {
		ellipticCurves[ellipticCurveI] = mustString(int(ellipticCurve))
	}
	result[3] = strings.Join(ellipticCurves, "-")

	ellipticCurvesFormats := make([]string, len(fp.EllipticCurvePointFormats))
	for ellipticCurvesFormatI, ellipticCurvesFormat := range fp.EllipticCurvePointFormats {
		ellipticCurvesFormats[ellipticCurvesFormatI] = mustString(int(ellipticCurvesFormat))
	}
	result[4] = strings.Join(ellipticCurvesFormats, "-")

	return strings.Join(result, ",")
}
