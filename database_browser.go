package device_utils

var AvailableBrowsers = map[string]map[string]*Browser{
	"brave": {
		"1.50.114": &Browser{
			Version:        "1.50.114",
			Name:           "brave",
			UserAgent:      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36",
			BrandHeader:    "\"Chromium\";v=\"112\", \"Brave\";v=\"112\", \"Not:A-Brand\";v=\"99\"",
			TlsFingerprint: &Browser_TLSFingerprint{Version: 771, CipherSuites: []Browser_TLSFingerprint_CipherSuite{4865, 4866, 4867, 49195, 49199, 49196, 49200, 52393, 52392, 49171, 49172, 156, 157, 47, 53}, Extensions: []Browser_TLSFingerprint_Extension{27, 16, 35, 11, 17513, 43, 13, 5, 23, 0, 18, 51, 10, 65281, 45, 21}, EllipticCurves: []Browser_TLSFingerprint_EllipticCurve{29, 23, 24}, EllipticCurvePointFormats: []Browser_TLSFingerprint_EllipticCurvePointFormat{0}},
		},
	},
	"chrome": {
		"111.0.5563.147": &Browser{
			Version:     "111.0.5563.147",
			Name:        "chrome",
			UserAgent:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",
			BrandHeader: "\"Chromium\";v=\"112\", \"Google Chrome\";v=\"112\", \"Not:A-Brand\";v=\"99\"",
			TlsFingerprint: &Browser_TLSFingerprint{
				Version:                   771,
				CipherSuites:              []Browser_TLSFingerprint_CipherSuite{4865, 4866, 4867, 49195, 49199, 49196, 49200, 52393, 52392, 49171, 49172, 156, 157, 47, 53},
				Extensions:                []Browser_TLSFingerprint_Extension{27, 16, 35, 11, 17513, 43, 13, 5, 23, 0, 18, 51, 10, 65281, 45, 21},
				EllipticCurves:            []Browser_TLSFingerprint_EllipticCurve{29, 23, 24},
				EllipticCurvePointFormats: []Browser_TLSFingerprint_EllipticCurvePointFormat{0},
				ExtensionData: []*Browser_TLSFingerprint_ExtensionData{
					{
						ExtensionId: Browser_TLSFingerprint_SIGNATURE_ALGORITHMS,
						SignatureAlgorithms: &Browser_TLSFingerprint_ExtensionData_SignatureAlgorithms{
							SupportedSignatureAlgorithms: []Browser_TLSFingerprint_SignatureScheme{
								Browser_TLSFingerprint_ECDSA_SECP256R1_SHA256,
								Browser_TLSFingerprint_RSA_PSS_RSAE_SHA256,
								Browser_TLSFingerprint_RSA_PKCS1_SHA256,
								Browser_TLSFingerprint_ECDSA_SECP384R1_SHA384,
								Browser_TLSFingerprint_RSA_PSS_RSAE_SHA384,
								Browser_TLSFingerprint_RSA_PKCS1_SHA384,
								Browser_TLSFingerprint_RSA_PSS_RSAE_SHA512,
								Browser_TLSFingerprint_RSA_PKCS1_SHA512,
							},
						},
					},
					{
						ExtensionId: Browser_TLSFingerprint_COMPRESS_CERTIFICATE,
						CompressCertificate: &Browser_TLSFingerprint_ExtensionData_CompressCertificate{
							Algorithms: []Browser_TLSFingerprint_ExtensionData_CompressCertificate_CertificateCompression{
								Browser_TLSFingerprint_ExtensionData_CompressCertificate_BROTLI,
							},
						},
					},
					{
						ExtensionId: Browser_TLSFingerprint_SUPPORTED_VERSIONS,
						SupportedVersions: &Browser_TLSFingerprint_ExtensionData_SupportedVersions{
							Versions: []Browser_TLSFingerprint_ProtocolVersion{
								Browser_TLSFingerprint_TLS1_3,
								Browser_TLSFingerprint_TLS1_2,
							},
						},
					},
					{
						ExtensionId: Browser_TLSFingerprint_EXTENSION_ENCRYPTED_CLIENT_HELLO,
						ExtensionEncryptedClientHello: &Browser_TLSFingerprint_ExtensionData_ExtensionEncryptedClientHello{
							CandidateCipherSuites: []*Browser_TLSFingerprint_ExtensionData_ExtensionEncryptedClientHello_HPKESymmetricCipherSuite{
								{
									KdfId:  Browser_TLSFingerprint_ExtensionData_ExtensionEncryptedClientHello_HKDF_SHA256,
									AeadId: Browser_TLSFingerprint_ExtensionData_ExtensionEncryptedClientHello_HPKEAEAD_AES_128_GCM,
								},
								{
									KdfId:  Browser_TLSFingerprint_ExtensionData_ExtensionEncryptedClientHello_HKDF_SHA256,
									AeadId: Browser_TLSFingerprint_ExtensionData_ExtensionEncryptedClientHello_HPKEAEAD_CHACHA20POLY1305,
								},
							},
							CandidatePayloadLens: []uint32{
								128, 160, 192, 224,
							},
						},
					},
					{
						ExtensionId: Browser_TLSFingerprint_EXTENSION_APPLICATIONS_SETTINGS,
						ExtensionApplicationsSettings: &Browser_TLSFingerprint_ExtensionData_ExtensionApplicationsSettings{
							Protocols: []string{
								"h2",
							},
						},
					},
					{
						ExtensionId: Browser_TLSFingerprint_APPLICATION_LAYER_PROTOCOL_NEGOTIATION,
						ApplicationLayerProtocolNegotiation: &Browser_TLSFingerprint_ExtensionData_ApplicationLayerProtocolNegotiation{
							Protocols: []string{
								"h2",
								"http/1.1",
							},
						},
					},
				},
			},
		},
	},
}
