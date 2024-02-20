package device_utils

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type PeetResponse struct {
	IP          string `json:"ip"`
	HTTPVersion string `json:"http_version"`
	Method      string `json:"method"`
	UserAgent   string `json:"user_agent"`
	TLS         TLS    `json:"tls"`
	Http2       *Http2 `json:"http2,omitempty"`
	Tcpip       Tcpip  `json:"tcpip"`
}

type Http2 struct {
	AkamaiFingerprint     string      `json:"akamai_fingerprint"`
	AkamaiFingerprintHash string      `json:"akamai_fingerprint_hash"`
	SentFrames            []SentFrame `json:"sent_frames"`
}

type SentFrame struct {
	FrameType string    `json:"frame_type"`
	Length    int64     `json:"length"`
	Settings  []string  `json:"settings,omitempty"`
	Increment *int64    `json:"increment,omitempty"`
	StreamID  *int64    `json:"stream_id,omitempty"`
	Headers   []string  `json:"headers,omitempty"`
	Flags     []string  `json:"flags,omitempty"`
	Priority  *Priority `json:"priority,omitempty"`
}

type Priority struct {
	Weight    int64 `json:"weight"`
	DependsOn int64 `json:"depends_on"`
	Exclusive int64 `json:"exclusive"`
}

type TLS struct {
	Ciphers              []string    `json:"ciphers"`
	Extensions           []Extension `json:"extensions"`
	TLSVersionRecord     string      `json:"tls_version_record"`
	TLSVersionNegotiated string      `json:"tls_version_negotiated"`
	Ja3                  string      `json:"ja3"`
	Ja3Hash              string      `json:"ja3_hash"`
	Ja4                  string      `json:"ja4"`
	Peetprint            string      `json:"peetprint"`
	PeetprintHash        string      `json:"peetprint_hash"`
	ClientRandom         string      `json:"client_random"`
	SessionID            string      `json:"session_id"`
}

type Extension struct {
	Name                       string              `json:"name"`
	Protocols                  []string            `json:"protocols,omitempty"`
	SharedKeys                 []map[string]string `json:"shared_keys,omitempty"`
	Data                       *string             `json:"data,omitempty"`
	SignatureAlgorithms        []string            `json:"signature_algorithms,omitempty"`
	MasterSecretData           *string             `json:"master_secret_data,omitempty"`
	ExtendedMasterSecretData   *string             `json:"extended_master_secret_data,omitempty"`
	StatusRequest              *StatusRequest      `json:"status_request,omitempty"`
	SupportedGroups            []string            `json:"supported_groups,omitempty"`
	Algorithms                 []string            `json:"algorithms,omitempty"`
	EllipticCurvesPointFormats []string            `json:"elliptic_curves_point_formats,omitempty"`
	Versions                   []string            `json:"versions,omitempty"`
	PSKKeyExchangeMode         *string             `json:"PSK_Key_Exchange_Mode,omitempty"`
	ServerName                 *string             `json:"server_name,omitempty"`
	SignatureHashAlgorithms    []string            `json:"signature_hash_algorithms,omitempty"`
}

type SharedKey struct {
	TLSGREASE0Xdada *string `json:"TLS_GREASE (0xdada),omitempty"`
	X2551929        *string `json:"X25519 (29),omitempty"`
	P25623          *string `json:"P-256 (23),omitempty"`
}

type StatusRequest struct {
	CertificateStatusType   string `json:"certificate_status_type"`
	ResponderIDListLength   int64  `json:"responder_id_list_length"`
	RequestExtensionsLength int64  `json:"request_extensions_length"`
}

type Tcpip struct {
	CapLength int64 `json:"cap_length"`
	DstPort   int64 `json:"dst_port"`
	SrcPort   int64 `json:"src_port"`
	IP        IP    `json:"ip"`
	TCP       TCP   `json:"tcp"`
}

type IP struct {
	ID        int64  `json:"id"`
	TTL       int64  `json:"ttl"`
	IPVersion int64  `json:"ip_version"`
	DstIP     string `json:"dst_ip"`
	SrcIP     string `json:"src_ip"`
}

type TCP struct {
	ACK      int64 `json:"ack"`
	Checksum int64 `json:"checksum"`
	Seq      int64 `json:"seq"`
	Window   int64 `json:"window"`
}

func MustUint(in string) uint64 {
	result, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func MustInt(in string) int64 {
	result, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}

func MustUintBetween(in string) uint64 {
	openingIndex := strings.LastIndex(in, "(")
	closingIndex := strings.LastIndex(in, ")")

	if openingIndex == -1 || closingIndex == -1 {
		panic(fmt.Errorf("parentheses not found"))
	}

	numberStr := in[openingIndex+1 : closingIndex]
	return MustUint(numberStr)
}

// FromPEET Imports and normalizes output from https://tls.peet.ws/api/all
func (b *Browser) FromPEET(response *PeetResponse) error {
	b.TlsFingerprint = &Browser_TLSFingerprint{
		CipherSuites:              make([]Browser_TLSFingerprint_CipherSuite, 0),
		Extensions:                make([]Browser_TLSFingerprint_Extension, 0),
		EllipticCurves:            make([]Browser_TLSFingerprint_EllipticCurve, 0),
		EllipticCurvePointFormats: make([]Browser_TLSFingerprint_EllipticCurvePointFormat, 0),
		ExtensionData:             make([]*Browser_TLSFingerprint_ExtensionData, 0),
	}
	// Set JA3 stuff
	for sectionIdx, section := range strings.Split(response.TLS.Ja3, ",") {
		switch sectionIdx {
		case 0:
			// TLS version
			b.TlsFingerprint.Version = Browser_TLSFingerprint_ProtocolVersion(MustUint(section))
			break
		case 1:
			// Ciphers
			for _, element := range strings.Split(section, "-") {
				b.TlsFingerprint.CipherSuites = append(b.TlsFingerprint.CipherSuites, Browser_TLSFingerprint_CipherSuite(MustUint(element)))
			}
			break
		case 2:
			// Extensions
			for _, element := range strings.Split(section, "-") {
				b.TlsFingerprint.Extensions = append(b.TlsFingerprint.Extensions, Browser_TLSFingerprint_Extension(MustUint(element)))
			}
			break
		case 3:
			// Elliptic Curves
			for _, element := range strings.Split(section, "-") {
				b.TlsFingerprint.EllipticCurves = append(b.TlsFingerprint.EllipticCurves, Browser_TLSFingerprint_EllipticCurve(MustUint(element)))
			}
			break
		case 4:
			// Elliptic Curve Point Formats
			for _, element := range strings.Split(section, "-") {
				b.TlsFingerprint.EllipticCurvePointFormats = append(b.TlsFingerprint.EllipticCurvePointFormats, Browser_TLSFingerprint_EllipticCurvePointFormat(MustUint(element)))
			}
			break
		}
	}

	// Set extension data
	for _, extensionRawData := range response.TLS.Extensions {
		if strings.HasPrefix(extensionRawData.Name, "TLS_GREASE") {
			continue
		}
		extensionId := MustUintBetween(extensionRawData.Name)

		switch extensionId {
		case 10:
			// Supported groups
			// extensionData := &Browser_TLSFingerprint_ExtensionData{ExtensionId: Browser_TLSFingerprint_Extension(extensionId)}
			// for _, e := range extensionRawData.SupportedGroups {
			// 	if strings.HasPrefix(e, "TLS_GREASE") {
			// 		continue
			// 	}
			// }
			break
		case 13:
			// Signature Algos
			extensionData := &Browser_TLSFingerprint_ExtensionData{ExtensionId: Browser_TLSFingerprint_Extension(extensionId)}
			extensionData.SignatureAlgorithms = &Browser_TLSFingerprint_ExtensionData_SignatureAlgorithms{
				SupportedSignatureAlgorithms: []Browser_TLSFingerprint_SignatureScheme{},
			}
			versionMap := map[string]Browser_TLSFingerprint_SignatureScheme{
				"rsa_pkcs1_sha256":       Browser_TLSFingerprint_SignatureScheme(0x0401),
				"rsa_pkcs1_sha384":       Browser_TLSFingerprint_SignatureScheme(0x0501),
				"rsa_pkcs1_sha512":       Browser_TLSFingerprint_SignatureScheme(0x0601),
				"ecdsa_secp256r1_sha256": Browser_TLSFingerprint_SignatureScheme(0x0403),
				"ecdsa_secp384r1_sha384": Browser_TLSFingerprint_SignatureScheme(0x0503),
				"ecdsa_secp521r1_sha512": Browser_TLSFingerprint_SignatureScheme(0x0603),
				"rsa_pss_rsae_sha256":    Browser_TLSFingerprint_SignatureScheme(0x0804),
				"rsa_pss_rsae_sha384":    Browser_TLSFingerprint_SignatureScheme(0x0805),
				"rsa_pss_rsae_sha512":    Browser_TLSFingerprint_SignatureScheme(0x0806),
				"ed25519":                Browser_TLSFingerprint_SignatureScheme(0x0807),
				"ed448":                  Browser_TLSFingerprint_SignatureScheme(0x0808),
				"rsa_pss_pss_sha256":     Browser_TLSFingerprint_SignatureScheme(0x0809),
				"rsa_pss_pss_sha384":     Browser_TLSFingerprint_SignatureScheme(0x080a),
				"rsa_pss_pss_sha512":     Browser_TLSFingerprint_SignatureScheme(0x080b),
				"rsa_pkcs1_sha1":         Browser_TLSFingerprint_SignatureScheme(0x0201),
				"ecdsa_sha1":             Browser_TLSFingerprint_SignatureScheme(0x0203),
			}
			for _, e := range extensionRawData.SignatureAlgorithms {
				eResult, ok := versionMap[e]
				if ok {
					extensionData.SignatureAlgorithms.SupportedSignatureAlgorithms = append(extensionData.SignatureAlgorithms.SupportedSignatureAlgorithms, eResult)
				}
			}
			b.TlsFingerprint.ExtensionData = append(b.TlsFingerprint.ExtensionData, extensionData)
		case 16:
			// ALPN
			extensionData := &Browser_TLSFingerprint_ExtensionData{ExtensionId: Browser_TLSFingerprint_Extension(extensionId)}
			extensionData.ApplicationLayerProtocolNegotiation = &Browser_TLSFingerprint_ExtensionData_ApplicationLayerProtocolNegotiation{
				Protocols: extensionRawData.Protocols,
			}
			b.TlsFingerprint.ExtensionData = append(b.TlsFingerprint.ExtensionData, extensionData)
			break
		case 27:
			// Compress Certificates
			extensionData := &Browser_TLSFingerprint_ExtensionData{ExtensionId: Browser_TLSFingerprint_Extension(extensionId)}
			extensionData.CompressCertificate = &Browser_TLSFingerprint_ExtensionData_CompressCertificate{
				Algorithms: []Browser_TLSFingerprint_ExtensionData_CompressCertificate_CertificateCompression{},
			}
			for _, algo := range extensionRawData.Algorithms {
				extensionData.CompressCertificate.Algorithms = append(extensionData.CompressCertificate.Algorithms, Browser_TLSFingerprint_ExtensionData_CompressCertificate_CertificateCompression(MustUintBetween(algo)))
			}
			b.TlsFingerprint.ExtensionData = append(b.TlsFingerprint.ExtensionData, extensionData)
			break
		case 43:
			// Supported TLS versions
			extensionData := &Browser_TLSFingerprint_ExtensionData{ExtensionId: Browser_TLSFingerprint_Extension(extensionId)}
			extensionData.SupportedVersions = &Browser_TLSFingerprint_ExtensionData_SupportedVersions{
				Versions: []Browser_TLSFingerprint_ProtocolVersion{},
			}
			versionMap := map[string]Browser_TLSFingerprint_ProtocolVersion{
				"TLS 1.3": Browser_TLSFingerprint_TLS1_3,
				"TLS 1.2": Browser_TLSFingerprint_TLS1_2,
				"TLS 1.1": Browser_TLSFingerprint_TLS1_1,
				"TLS 1.0": Browser_TLSFingerprint_TLS1,
			}
			for _, e := range extensionRawData.Versions {
				if strings.HasPrefix(e, "TLS_GREASE") {
					continue
				}
				eResult, ok := versionMap[e]
				if ok {
					extensionData.SupportedVersions.Versions = append(extensionData.SupportedVersions.Versions, eResult)
				}
			}
			b.TlsFingerprint.ExtensionData = append(b.TlsFingerprint.ExtensionData, extensionData)
			break
		case 45:
			// PSK exchange modes
			extensionData := &Browser_TLSFingerprint_ExtensionData{ExtensionId: Browser_TLSFingerprint_Extension(extensionId)}
			extensionData.PskKeyExchangeModes = &Browser_TLSFingerprint_ExtensionData_PSKKeyExchangeModes{
				Modes: []Browser_TLSFingerprint_ExtensionData_PSKKeyExchangeModes_Mode{},
			}
			if extensionRawData.PSKKeyExchangeMode != nil {
				extensionData.PskKeyExchangeModes.Modes = append(extensionData.PskKeyExchangeModes.Modes, Browser_TLSFingerprint_ExtensionData_PSKKeyExchangeModes_Mode(MustUintBetween(*extensionRawData.PSKKeyExchangeMode)))
				b.TlsFingerprint.ExtensionData = append(b.TlsFingerprint.ExtensionData, extensionData)
			}
			break
		case 51:
			// Shared Keys
			extensionData := &Browser_TLSFingerprint_ExtensionData{ExtensionId: Browser_TLSFingerprint_Extension(extensionId)}
			extensionData.KeyShareExtension = &Browser_TLSFingerprint_ExtensionData_KeyShareExtension{
				KeyShares: []*Browser_TLSFingerprint_ExtensionData_KeyShareExtension_KeyShare{},
			}
			for _, keyShare := range extensionRawData.SharedKeys {
				for name, _ := range keyShare { // _ = hex data, TODO: allow importing that?
					if strings.HasPrefix(name, "TLS_GREASE") {
						continue
					}
					extensionData.KeyShareExtension.KeyShares = append(extensionData.KeyShareExtension.KeyShares, &Browser_TLSFingerprint_ExtensionData_KeyShareExtension_KeyShare{
						Group: Browser_TLSFingerprint_EllipticCurve(MustUintBetween(name)),
					})
				}
			}
			b.TlsFingerprint.ExtensionData = append(b.TlsFingerprint.ExtensionData, extensionData)
			break
		case 17513:
			// Application Settings
			extensionData := &Browser_TLSFingerprint_ExtensionData{ExtensionId: Browser_TLSFingerprint_Extension(extensionId)}
			extensionData.ExtensionApplicationsSettings = &Browser_TLSFingerprint_ExtensionData_ExtensionApplicationsSettings{
				Protocols: extensionRawData.Protocols,
			}
			b.TlsFingerprint.ExtensionData = append(b.TlsFingerprint.ExtensionData, extensionData)
			break
		case 65037:
			// ECH
			// TODO: don't hardcode this like this
			if extensionRawData.Data != nil {
				extensionData := &Browser_TLSFingerprint_ExtensionData{
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
				}
				b.TlsFingerprint.ExtensionData = append(b.TlsFingerprint.ExtensionData, extensionData)
			}
			break
		}
	}

	// Set AKAMAI stuff
	if response.Http2 != nil && response.Http2.AkamaiFingerprint != "" {
		b.HttpFingerprint = &Browser_HTTPFingerprint{
			SettingsFrame: &Browser_HTTPFingerprint_SettingsFrameOpts{
				// -1 assumes not present
				HeaderTableSize:      -1,
				EnablePush:           -1,
				MaxConcurrentStreams: -1,
				InitialWindowSize:    -1,
				MaxFrameSize:         -1,
				MaxHeaderListSize:    -1,
			},
			PseudoHeaderOrder: []string{},
			PriorityFrames:    []*Browser_HTTPFingerprint_PriorityFrameOpts{},
		}
		for sectionId, section := range strings.Split(response.Http2.AkamaiFingerprint, "|") {
			switch sectionId {
			case 0:
				// SETTINGS frame
				for _, parameterData := range strings.Split(section, ",") {
					parameterDataSplit := strings.Split(parameterData, ":")
					if len(parameterDataSplit) == 2 {
						parameterValue := MustInt(parameterDataSplit[1])
						switch parameterDataSplit[0] {
						case "1":
							b.HttpFingerprint.SettingsFrame.HeaderTableSize = parameterValue
							break
						case "2":
							b.HttpFingerprint.SettingsFrame.EnablePush = parameterValue
							break
						case "3":
							b.HttpFingerprint.SettingsFrame.MaxConcurrentStreams = parameterValue
							break
						case "4":
							b.HttpFingerprint.SettingsFrame.InitialWindowSize = parameterValue
							break
						case "5":
							b.HttpFingerprint.SettingsFrame.MaxFrameSize = parameterValue
							break
						case "6":
							b.HttpFingerprint.SettingsFrame.MaxHeaderListSize = parameterValue
							break
						}
					}
				}
				break
			case 1:
				// WINDOW_UPDATE frame
				b.HttpFingerprint.WindowUpdateIncrement = MustInt(section)
				break
			case 2:
				// PRIORITY frames
				if section == "0" {
					continue
				}
				for _, priorityData := range strings.Split(section, ",") {
					priorityDataSplit := strings.Split(priorityData, ":")
					if len(priorityDataSplit) == 4 {
						priorityFrame := &Browser_HTTPFingerprint_PriorityFrameOpts{
							StreamId:  MustInt(priorityDataSplit[0]),
							StreamDep: MustInt(priorityDataSplit[2]),
							Exclusive: priorityDataSplit[2] == "1",
							Weight:    int32(MustInt(priorityDataSplit[3])),
						}
						b.HttpFingerprint.PriorityFrames = append(b.HttpFingerprint.PriorityFrames, priorityFrame)
					}
				}
				break
			case 3:
				// Pseudo Header Order
				headers := map[string]string{
					"m": ":method",
					"p": ":path",
					"a": ":authority",
					"s": ":scheme",
				}
				for _, pseudoHeader := range strings.Split(section, ",") {
					b.HttpFingerprint.PseudoHeaderOrder = append(b.HttpFingerprint.PseudoHeaderOrder, headers[pseudoHeader])
				}
				break
			}
		}
	}

	// Set priority for HEADERS frame
	for _, frameRawData := range response.Http2.SentFrames {
		if frameRawData.FrameType != "SETTINGS" {
			continue
		}

		if frameRawData.Priority == nil {
			break
		}

		b.HttpFingerprint.HeaderFramePriority = &Browser_HTTPFingerprint_PriorityFrameOpts{
			StreamId:  *frameRawData.StreamID,
			StreamDep: frameRawData.Priority.DependsOn,
			Exclusive: frameRawData.Priority.Exclusive == 1,
			Weight:    int32(frameRawData.Priority.Weight),
		}
	}
	return nil
}

func (b *Browser) FromPEETRaw(data []byte) error {
	response := &PeetResponse{}
	err := json.Unmarshal(data, response)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	return b.FromPEET(response)
}
