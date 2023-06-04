package device_utils

import (
	"encoding/base64"
	"encoding/hex"
	"math/rand"
	"strconv"
)

func NewAndroidID() *AndroidDevice_ID {
	result := &AndroidDevice_ID{}
	_ = result.Random()
	return result
}

func (id *AndroidDevice_ID) FromHex(idStr string) error {
	result, err := strconv.ParseUint(idStr, 16, 64)
	if err == nil {
		id.SetID(result)
	}
	return err
}

func (id *AndroidDevice_ID) Random() error {
	b := make([]byte, 8)
	_, err := rand.Read(b)
	if err == nil {
		err = id.FromHex(hex.EncodeToString(b))
	}
	return err
}

func (id *AndroidDevice_ID) ToDecimalString() string {
	return strconv.FormatUint(id.GetID(), 10)
}

func (id *AndroidDevice_ID) ToHexString() string {
	return strconv.FormatUint(id.GetID(), 16)
}

func (id *AndroidDevice_ID) ToBase64String() string {
	hByte, _ := hex.DecodeString(id.ToHexString())
	return base64.StdEncoding.EncodeToString(hByte)
}

func (id *AndroidDevice_ID) Equals(comparison *AndroidDevice_ID) bool {
	return id.GetID() == comparison.GetID()
}

func (id *AndroidDevice_ID) IsNull() bool {
	return id.GetID() < 1
}

func (id *AndroidDevice_ID) GetID() uint64 {
	result := id.Id
	return result
}

func (id *AndroidDevice_ID) SetID(idN uint64) {
	id.Id = idN
}
