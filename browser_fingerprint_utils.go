package device_utils

import (
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"strconv"
	"strings"
)

type DLFingerprint struct {
	AppCodeName         string             `json:"appCodeName"`
	AppName             string             `json:"appName"`
	AppVersion          string             `json:"appVersion"`
	CookieEnabled       bool               `json:"cookieEnabled"`
	DeviceMemory        int64              `json:"deviceMemory"`
	DoNotTrack          interface{}        `json:"doNotTrack"`
	HardwareConcurrency int64              `json:"hardwareConcurrency"`
	Language            string             `json:"language"`
	Languages           []string           `json:"languages"`
	MaxTouchPoints      int64              `json:"maxTouchPoints"`
	PDFViewerEnabled    bool               `json:"pdfViewerEnabled"`
	Platform            string             `json:"platform"`
	Product             string             `json:"product"`
	ProductSub          string             `json:"productSub"`
	UserAgent           string             `json:"userAgent"`
	Vendor              string             `json:"vendor"`
	VendorSub           string             `json:"vendorSub"`
	Webdiver            bool               `json:"webdiver"`
	DevicePixelRatio    float64            `json:"devicePixelRatio"`
	InnerWidth          int64              `json:"innerWidth"`
	InnerHeight         int64              `json:"innerHeight"`
	OuterWidth          int64              `json:"outerWidth"`
	OuterHeight         int64              `json:"outerHeight"`
	Screen              Screen             `json:"screen"`
	Plugins             []Plugin           `json:"plugins"`
	UserActivation      UserActivation     `json:"userActivation"`
	ChromeApp           ChromeApp          `json:"chrome.app"`
	Wow64               bool               `json:"wow64"`
	HighEntropyValues   *HighEntropyValues `json:"HighEntropyValues,omitempty"`
	Darkmode            bool               `json:"darkmode"`
	AvailabeFonts       []string           `json:"availabeFonts"`
	StackNative         []float64          `json:"stack_native"`
	TimingNative        int64              `json:"timing_native"`
	Permissions         Permissions        `json:"permissions"`
	Navigator           []string           `json:"navigator"`
	Window              []string           `json:"window"`
	Document            []string           `json:"document"`
	DocumentElement     []string           `json:"documentElement"`
	SpeechSynthesis     []SpeechSynthesis  `json:"speechSynthesis"`
	CSS                 map[string]any     `json:"css"`
	AudioTypes          map[string]string  `json:"audioTypes"`
	VideoTypes          map[string]string  `json:"videoTypes"`
	AudioContext        AudioContext       `json:"audioContext"`
	Webrtc              Webrtc             `json:"webrtc"`
	Webgpu              Webgpu             `json:"webgpu"`
	MediaDevices        []MediaDevice      `json:"mediaDevices"`
	IsBot               bool               `json:"is_bot"`
	Status              string             `json:"status"`
	StackWorker         []float64          `json:"stack_worker"`
	TimingWorker        int64              `json:"timing_worker"`
	Gl                  map[string][]any   `json:"gl"`
	Gl2                 map[string][]any   `json:"gl2"`
	GlExperimental      map[string][]any   `json:"gl_experimental"`
}

type AudioContext struct {
	BaseLatency   float64             `json:"baseLatency"`
	OutputLatency int64               `json:"outputLatency"`
	SinkID        string              `json:"sinkId"`
	Destination   Destination         `json:"destination"`
	CurrentTime   int64               `json:"currentTime"`
	SampleRate    int64               `json:"sampleRate"`
	Listener      map[string]Listener `json:"listener"`
	State         string              `json:"state"`
	AudioWorklet  AudioWorklet        `json:"audioWorklet"`
}

type AudioWorklet struct {
}

type Context struct {
	BaseLatency   float64      `json:"baseLatency"`
	OutputLatency int64        `json:"outputLatency"`
	SinkID        string       `json:"sinkId"`
	Destination   Destination  `json:"destination"`
	CurrentTime   int64        `json:"currentTime"`
	SampleRate    int64        `json:"sampleRate"`
	Listener      AudioWorklet `json:"listener"`
	State         string       `json:"state"`
	AudioWorklet  AudioWorklet `json:"audioWorklet"`
}

type Destination struct {
	MaxChannelCount       int64    `json:"maxChannelCount"`
	Context               *Context `json:"context,omitempty"`
	NumberOfInputs        int64    `json:"numberOfInputs"`
	NumberOfOutputs       int64    `json:"numberOfOutputs"`
	ChannelCount          int64    `json:"channelCount"`
	ChannelCountMode      string   `json:"channelCountMode"`
	ChannelInterpretation string   `json:"channelInterpretation"`
}

type Listener struct {
	Value          int64   `json:"value"`
	AutomationRate string  `json:"automationRate"`
	DefaultValue   int64   `json:"defaultValue"`
	MinValue       float64 `json:"minValue"`
	MaxValue       float64 `json:"maxValue"`
}

type AudioTypes struct {
	Audio3Gpp              string `json:"audio/3gpp"`
	AudioAAC               string `json:"audio/aac"`
	AudioFLAC              string `json:"audio/flac"`
	AudioMPEG              string `json:"audio/mpeg"`
	AudioMp4CodecsAC3      string `json:"audio/mp4; codecs=\"ac-3\""`
	AudioMp4CodecsEc3      string `json:"audio/mp4; codecs=\"ec-3\""`
	AudioOggCodecsFLAC     string `json:"audio/ogg; codecs=\"flac\""`
	AudioOggCodecsVorbis   string `json:"audio/ogg; codecs=\"vorbis\""`
	AudioOggCodecsOpus     string `json:"audio/ogg; codecs=\"opus\""`
	AudioWavCodecs1        string `json:"audio/wav; codecs=\"1\""`
	AudioWebmCodecsVorbis  string `json:"audio/webm; codecs=\"vorbis\""`
	AudioWebmCodecsOpus    string `json:"audio/webm; codecs=\"opus\""`
	AudioMp4CodecsMp4A40_2 string `json:"audio/mp4; codecs=\"mp4a_40_2\""`
}

type ChromeApp struct {
	IsInstalled  bool         `json:"isInstalled"`
	InstallState InstallState `json:"InstallState"`
	RunningState RunningState `json:"RunningState"`
}

type InstallState struct {
	Disabled     string `json:"DISABLED"`
	Installed    string `json:"INSTALLED"`
	NotInstalled string `json:"NOT_INSTALLED"`
}

type RunningState struct {
	CannotRun  string `json:"CANNOT_RUN"`
	ReadyToRun string `json:"READY_TO_RUN"`
	Running    string `json:"RUNNING"`
}

type HighEntropyValues struct {
	Architecture    string  `json:"architecture"`
	Bitness         string  `json:"bitness"`
	Brands          []Brand `json:"brands,omitempty"`
	FullVersionList []Brand `json:"fullVersionList,omitempty"`
	Mobile          bool    `json:"mobile"`
	Model           string  `json:"model"`
	Platform        string  `json:"platform"`
	PlatformVersion string  `json:"platformVersion"`
	UaFullVersion   string  `json:"uaFullVersion"`
}

type Brand struct {
	Brand   string `json:"brand"`
	Version string `json:"version"`
}

type MediaDevice struct {
	DeviceID string `json:"deviceId"`
	Kind     string `json:"kind"`
	Label    string `json:"label"`
	GroupID  string `json:"groupId"`
}

type Permissions struct {
	BackgroundSync    string `json:"background-sync"`
	ClipboardRead     string `json:"clipboard-read"`
	ClipboardWrite    string `json:"clipboard-write"`
	Geolocation       string `json:"geolocation"`
	Microphone        string `json:"microphone"`
	Camera            string `json:"camera"`
	Notifications     string `json:"notifications"`
	PaymentHandler    string `json:"payment-handler"`
	Accelerometer     string `json:"accelerometer"`
	Gyroscope         string `json:"gyroscope"`
	Magnetometer      string `json:"magnetometer"`
	StorageAccess     string `json:"storage-access"`
	PersistentStorage string `json:"persistent-storage"`
	MIDI              string `json:"midi"`
	BackgroundFetch   string `json:"background-fetch"`
	ScreenWakeLock    string `json:"screen-wake-lock"`
	DisplayCapture    string `json:"display-capture"`
}

type Plugin struct {
	The0        Empty  `json:"0"`
	Name        string `json:"name"`
	Filename    string `json:"filename"`
	Description string `json:"description"`
	Length      int64  `json:"length"`
}

type Empty struct {
	Type        string `json:"type"`
	Suffixes    string `json:"suffixes"`
	Description string `json:"description"`
}

type Screen struct {
	AvailWidth  int64       `json:"availWidth"`
	AvailHeight int64       `json:"availHeight"`
	Width       int64       `json:"width"`
	Height      int64       `json:"height"`
	ColorDepth  int64       `json:"colorDepth"`
	PixelDepth  int64       `json:"pixelDepth"`
	AvailLeft   int64       `json:"availLeft"`
	AvailTop    int64       `json:"availTop"`
	Orientation Orientation `json:"orientation"`
	IsExtended  bool        `json:"isExtended"`
}

type Orientation struct {
	Angle int64  `json:"angle"`
	Type  string `json:"type"`
}

type SpeechSynthesis struct {
	VoiceURI     string `json:"voiceURI"`
	Name         string `json:"name"`
	Lang         string `json:"lang"`
	LocalService bool   `json:"localService"`
	Default      bool   `json:"default"`
}

type UserActivation struct {
	HasBeenActive bool `json:"hasBeenActive"`
	IsActive      bool `json:"isActive"`
}

type VideoTypes struct {
	VideoMp4CodecsFLAC       string `json:"video/mp4; codecs=\"flac\""`
	VideoOggCodecsTheora     string `json:"video/ogg; codecs=\"theora\""`
	VideoOggCodecsOpus       string `json:"video/ogg; codecs=\"opus\""`
	VideoWebmCodecsVp9Opus   string `json:"video/webm; codecs=\"vp9, opus\""`
	VideoWebmCodecsVp8Vorbis string `json:"video/webm; codecs=\"vp8, vorbis\" "`
}

type Webgpu struct {
	Features          Features         `json:"features"`
	Limits            map[string]int64 `json:"limits"`
	IsFallbackAdapter bool             `json:"isFallbackAdapter"`
	Vendor            string           `json:"vendor"`
	Architecture      string           `json:"architecture"`
	Device            string           `json:"device"`
	Description       string           `json:"description"`
}

type Features struct {
	Size int64 `json:"size"`
}

type Webrtc struct {
	Video Audio `json:"video"`
	Audio Audio `json:"audio"`
}

type Audio struct {
	Codecs           []Codec           `json:"codecs"`
	HeaderExtensions []HeaderExtension `json:"headerExtensions"`
}

type Codec struct {
	Channels    int32  `json:"channels,omitempty"`
	ClockRate   int64  `json:"clockRate"`
	MIMEType    string `json:"mimeType"`
	SDPFmtpLine string `json:"sdpFmtpLine,omitempty"`
}

type HeaderExtension struct {
	Direction string `json:"direction"`
	URI       string `json:"uri"`
}

type AliasedLineWidthRange struct {
	Integer    *int64
	IntegerMap map[string]int64
}

type Blend struct {
	Bool    *bool
	Integer *int64
}

type ColorWritemask struct {
	BoolArray []bool
	Integer   *int64
}

type CompressedTextureFormat struct {
	AudioWorklet *AudioWorklet
	Integer      *int64
}

// FromDLFingerprint Imports and normalizes output from https://github.com/kaliiiiiiiiii/driverless-fp-collector
func (b *Browser) FromDLFingerprint(response *DLFingerprint) error {
	b.UserAgent = response.UserAgent
	b.AppCodeName = response.AppCodeName
	b.AppName = response.AppName
	b.AppVersion = response.AppVersion
	b.DeviceMemory = response.DeviceMemory
	b.DoNotTrack = -1
	b.HardwareConcurrency = int32(response.HardwareConcurrency)
	b.Language = response.Language
	b.Languages = response.Languages
	b.MaxTouchPoints = int32(response.MaxTouchPoints)
	b.PdfViewerEnabled = response.PDFViewerEnabled
	b.Platform = response.Platform
	b.Product = response.Product
	b.ProductSub = response.ProductSub
	b.Vendor = response.Vendor
	b.VendorSub = response.VendorSub
	b.Webdriver = false // response.Webdiver
	b.DevicePixelRatio = response.DevicePixelRatio
	b.InnerWidth = int32(response.InnerWidth)
	b.InnerHeight = int32(response.InnerHeight)
	b.OuterWidth = int32(response.OuterWidth)
	b.OuterHeight = int32(response.OuterHeight)
	b.Screen = &Browser_BrowserScreen{
		AvailWidth:  int32(response.Screen.AvailWidth),
		AvailHeight: int32(response.Screen.AvailHeight),
		Width:       int32(response.Screen.Width),
		Height:      int32(response.Screen.Height),
		ColorDepth:  int32(response.Screen.ColorDepth),
		PixelDepth:  int32(response.Screen.PixelDepth),
		AvailLeft:   int32(response.Screen.AvailLeft),
		AvailTop:    int32(response.Screen.AvailTop),
		IsExtended:  response.Screen.IsExtended,
		Orientation: &Browser_BrowserScreen_Orientation{
			Angle: int32(response.Screen.Orientation.Angle),
			Type:  response.Screen.Orientation.Type,
		},
	}

	// Speech synthesis
	b.SpeechSynthesis = make([]*Browser_SpeechSynthesis, 0)
	for _, speechEngine := range response.SpeechSynthesis {
		b.SpeechSynthesis = append(b.SpeechSynthesis, &Browser_SpeechSynthesis{
			VoiceURI:     speechEngine.VoiceURI,
			Name:         speechEngine.Name,
			Lang:         speechEngine.Lang,
			LocalService: speechEngine.LocalService,
			Default:      speechEngine.Default,
		})
	}

	// Plugins
	b.Plugins = make([]*Browser_Plugin, 0)
	for _, plugin := range response.Plugins {
		pluginObj := &Browser_Plugin{
			Name:               plugin.Name,
			FileName:           plugin.Filename,
			Description:        plugin.Description,
			SupportedMIMETypes: make(map[string]*Browser_Plugin_MIMEType),
		}

		pluginObj.SupportedMIMETypes[plugin.The0.Type] = &Browser_Plugin_MIMEType{
			Type:        plugin.The0.Type,
			Suffixes:    plugin.The0.Suffixes,
			Description: plugin.The0.Description,
		}

		b.Plugins = append(b.Plugins, pluginObj)
	}

	// webGPU
	b.WebGPU = &Browser_WebGPU{
		Features: &Browser_WebGPU_Features{
			Size: response.Webgpu.Features.Size,
		},
		Limits:            response.Webgpu.Limits,
		IsFallbackAdapter: response.Webgpu.IsFallbackAdapter,
		Vendor:            response.Webgpu.Vendor,
		Architecture:      response.Webgpu.Architecture,
		Device:            response.Webgpu.Device,
		Description:       response.Webgpu.Description,
	}

	b.WebRTC = &Browser_WebRTC{
		Video: &Browser_WebRTC_CodecInformation{},
		Audio: &Browser_WebRTC_CodecInformation{},
	}
	for _, videoCodec := range response.Webrtc.Video.Codecs {
		b.WebRTC.Video.Codecs = append(b.WebRTC.Video.Codecs, &Browser_WebRTC_Codec{
			Channels:   0,
			ClockRate:  videoCodec.ClockRate,
			MimeType:   videoCodec.MIMEType,
			SdpFmtLine: "",
		})
	}
	for _, audioCodec := range response.Webrtc.Audio.Codecs {
		b.WebRTC.Audio.Codecs = append(b.WebRTC.Audio.Codecs, &Browser_WebRTC_Codec{
			Channels:   audioCodec.Channels,
			ClockRate:  audioCodec.ClockRate,
			MimeType:   audioCodec.MIMEType,
			SdpFmtLine: audioCodec.SDPFmtpLine,
		})
	}
	for _, videoHeaderExtension := range response.Webrtc.Video.HeaderExtensions {
		b.WebRTC.Video.HeaderExtensions = append(b.WebRTC.Video.HeaderExtensions, &Browser_WebRTC_HeaderExtension{
			Direction: videoHeaderExtension.Direction,
			Uri:       videoHeaderExtension.URI,
		})
	}
	for _, audioHeaderExtension := range response.Webrtc.Audio.HeaderExtensions {
		b.WebRTC.Audio.HeaderExtensions = append(b.WebRTC.Audio.HeaderExtensions, &Browser_WebRTC_HeaderExtension{
			Direction: audioHeaderExtension.Direction,
			Uri:       audioHeaderExtension.URI,
		})
	}

	b.AvailableFonts = &Browser_BrowserCollection{ListData: response.AvailabeFonts}
	b.Navigator = &Browser_BrowserCollection{ListData: response.Navigator}
	b.Window = &Browser_BrowserCollection{ListData: response.Window}
	b.Document = &Browser_BrowserCollection{ListData: response.Document}
	b.DocumentElement = &Browser_BrowserCollection{ListData: response.DocumentElement}
	b.AudioTypes = &Browser_BrowserCollection{MapData: response.AudioTypes}
	b.VideoTypes = &Browser_BrowserCollection{MapData: response.VideoTypes}

	b.Css = &Browser_BrowserCollection{MapData: make(map[string]string)}
	for key, value := range response.CSS {
		switch typedValue := value.(type) {
		case string:
			b.Css.MapData[key] = typedValue
			break
		}
	}

	if response.HighEntropyValues != nil {
		b.HighEntropyValues = &Browser_HighEntropyValues{
			Architecture:    response.HighEntropyValues.Architecture,
			Bitness:         response.HighEntropyValues.Bitness,
			Mobile:          response.HighEntropyValues.Mobile,
			Model:           response.HighEntropyValues.Model,
			Platform:        response.HighEntropyValues.Platform,
			PlatformVersion: response.HighEntropyValues.PlatformVersion,
			UsFullVersion:   response.HighEntropyValues.UaFullVersion,
			Brands:          nil,
			FullVersionList: nil,
		}
		brandHeaderElements := []string{}
		for _, brand := range response.HighEntropyValues.Brands {
			b.HighEntropyValues.Brands = append(b.HighEntropyValues.Brands, &Browser_HighEntropyValues_Brand{
				Brand:   brand.Brand,
				Version: brand.Version,
			})
			brandHeaderElements = append(brandHeaderElements, fmt.Sprintf("\"%s\";v=\"%s\"", brand.Brand, brand.Version))
		}
		if len(brandHeaderElements) > 0 {
			b.BrandHeader = strings.Join(brandHeaderElements, ",")
		}
		for _, brand := range response.HighEntropyValues.FullVersionList {
			b.HighEntropyValues.FullVersionList = append(b.HighEntropyValues.FullVersionList, &Browser_HighEntropyValues_Brand{
				Brand:   brand.Brand,
				Version: brand.Version,
			})
		}
	}

	// Parse GL capabilities
	b.Gl = &Browser_BrowserCollection{GlCapabilities: make(map[string]*Browser_GLCapability)}
	for key, value := range response.Gl {
		valueLength := len(value)
		enumElem := value[valueLength-1]

		data := &Browser_GLCapability{
			BoolValue:   []bool{},
			IntValue:    []int64{},
			FloatValue:  []float64{},
			StringValue: []string{},
			EnumValue:   int64(int32(enumElem.(float64))),
			EnumName:    key,
		}

		buildGLCapabilities(key, data, value[:valueLength-1])

		b.Gl.GlCapabilities[key] = data
	}
	b.Gl2 = &Browser_BrowserCollection{GlCapabilities: make(map[string]*Browser_GLCapability)}
	for key, value := range response.Gl2 {
		valueLength := len(value)
		enumElem := value[valueLength-1]

		data := &Browser_GLCapability{
			BoolValue:   []bool{},
			IntValue:    []int64{},
			FloatValue:  []float64{},
			StringValue: []string{},
			EnumValue:   int64(enumElem.(float64)),
			EnumName:    key,
		}

		buildGLCapabilities(key, data, value[:valueLength-1])

		b.Gl2.GlCapabilities[key] = data
	}
	b.GlExperimental = &Browser_BrowserCollection{GlCapabilities: make(map[string]*Browser_GLCapability)}
	for key, value := range response.GlExperimental {
		valueLength := len(value)
		enumElem := value[valueLength-1]

		data := &Browser_GLCapability{
			BoolValue:   []bool{},
			IntValue:    []int64{},
			FloatValue:  []float64{},
			StringValue: []string{},
			EnumValue:   int64(enumElem.(float64)),
			EnumName:    key,
		}

		buildGLCapabilities(key, data, value[:valueLength-1])

		b.GlExperimental.GlCapabilities[key] = data
	}

	return nil
}

func buildGLCapabilities(name string, data *Browser_GLCapability, elements []any) {
	for _, element := range elements {
		switch typedElement := element.(type) {
		case string:
			data.StringValue = append(data.StringValue, typedElement)
			break
		case float64:
			data.FloatValue = append(data.FloatValue, typedElement)
			break
		case bool:
			data.BoolValue = append(data.BoolValue, typedElement)
			break
		case []any:
			// Array of booleans
			buildGLCapabilities(name, data, typedElement)
			break
		case map[string]any:
			intermediate := make([]any, len(typedElement))
			for key, elementValue := range typedElement {
				intKey, err := strconv.Atoi(key)
				if err != nil {
					panic(err)
				}
				intermediate[intKey] = elementValue
			}
			buildGLCapabilities(name, data, intermediate)
			break
		default:
			fmt.Println(spew.Sdump(typedElement))
		}
	}
}

func (b *Browser) FromDLFingerprintRaw(data []byte) error {
	response := &DLFingerprint{}
	err := json.Unmarshal(data, response)
	if err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	return b.FromDLFingerprint(response)
}
