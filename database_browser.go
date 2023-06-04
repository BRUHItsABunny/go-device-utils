package go_device_utils

var AvailableBrowsers = map[string]map[string]*Browser{
	"brave": {
		"1.50.114": &Browser{
			Version:        "1.50.114",
			Name:           "brave",
			UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
			BrandHeader:    "\"Chromium\";v=\"112\", \"Brave\";v=\"112\", \"Not:A-Brand\";v=\"99\"",
			TlsFingerprint: &Browser_TLSFingerprint{Version: 771, CipherSuites: []Browser_TLSFingerprint_CipherSuite{4865, 4866, 4867, 49195, 49199, 49196, 49200, 52393, 52392, 49171, 49172, 156, 157, 47, 53}, Extensions: []Browser_TLSFingerprint_Extensions{27, 16, 35, 11, 17513, 43, 13, 5, 23, 0, 18, 51, 10, 65281, 45, 21}, EllipticCurves: []Browser_TLSFingerprint_EllipticCurve{29, 23, 24}, EllipticCurvePointFormats: []Browser_TLSFingerprint_EllipticCurvePointFormat{0}},
		},
	},
	"chrome": {
		"111.0.5563.147": &Browser{
			Version:        "111.0.5563.147",
			Name:           "chrome",
			UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
			BrandHeader:    "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"",
			TlsFingerprint: &Browser_TLSFingerprint{Version: 771, CipherSuites: []Browser_TLSFingerprint_CipherSuite{4865, 4866, 4867, 49195, 49199, 49196, 49200, 52393, 52392, 49171, 49172, 156, 157, 47, 53}, Extensions: []Browser_TLSFingerprint_Extensions{27, 16, 35, 11, 17513, 43, 13, 5, 23, 0, 18, 51, 10, 65281, 45, 21}, EllipticCurves: []Browser_TLSFingerprint_EllipticCurve{29, 23, 24}, EllipticCurvePointFormats: []Browser_TLSFingerprint_EllipticCurvePointFormat{0}},
		},
	},
}
