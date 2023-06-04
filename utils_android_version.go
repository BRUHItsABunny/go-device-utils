package device_utils

import (
	"errors"
	"strconv"
	"strings"
)

func (version AndroidDevice_Version) IsValid() bool {
	if version == 0 {
		return false
	}
	_, result := AndroidDevice_Version_name[int32(version)]
	return result
}

func AndroidVersionFromVersionString(versionStr string) (AndroidDevice_Version, error) {
	var err error
	// 6 => 6.0, 6.1 => 6.1
	versionSplits := strings.Split(versionStr, ".")
	if len(versionSplits) == 1 {
		versionSplits = append(versionSplits, "0")
	}
	val, ok := AndroidDevice_Version_value["V"+strings.ToUpper(strings.Join(versionSplits[:2], "_"))]
	if !ok {
		err = ErrAndroidVersionVersionUnsupported
		val = 0
	}
	return AndroidDevice_Version(val), err
}

func AndroidVersionFromSDKString(sdkStr string) (AndroidDevice_Version, error) {
	sdk, err := strconv.ParseInt(sdkStr, 10, 64)
	if err == nil {
		_, ok := AndroidDevice_Version_name[int32(sdk)]
		if !ok || sdk < 1 {
			err = ErrAndroidVersionSDKUnsupported
		}
		return AndroidDevice_Version(int(sdk)), err
	}
	return AndroidDevice_Version(0), err
}

func (version AndroidDevice_Version) ToAndroidVersion() string {
	return strings.ReplaceAll(version.String()[1:], "_", ".")
}

func (version AndroidDevice_Version) ToAndroidSDK() string {
	return strconv.Itoa(int(version))
}

var (
	ErrAndroidVersionSDKUnsupported     = errors.New("the supplied SDK is unsupported")
	ErrAndroidVersionVersionUnsupported = errors.New("the supplied version is unsupported")
)
