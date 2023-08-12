package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"strconv"
	"time"
)

const interval int64 = 30

func GenerateOTP(secretKey string) string {
	m := hexStep()
	decodedKey := decodeKey(secretKey)
	hash := hmacSha1(m, decodedKey)
	offset := findOffset(hash)
	sub := binary.BigEndian.Uint32(hash[offset : offset+4])

	otp := strconv.Itoa(int((sub & 0x7fffffff) % 1000000))

	result := prefix(otp)

	return result
}

func hexStep() []byte {
	timestamp := time.Now().Unix()
	step := uint64(timestamp / interval)

	bs := make([]byte, 8)

	binary.BigEndian.PutUint64(bs, step)

	return bs
}

func decodeKey(key string) []byte {
	decoded, err := base32.StdEncoding.DecodeString(key)

	if err != nil {
		panic(err)
	}

	return decoded
}

func hmacSha1(m, key []byte) []byte {
	h := hmac.New(sha1.New, key)
	h.Write(m)

	return h.Sum(nil)
}

func findOffset(hash []byte) byte {
	return hash[19] & 15
}

func prefix(otp string) string {
	if len(otp) == 6 {
		return otp
	}

	limit := 6 - len(otp)

	for i := 0; i < limit; i++ {
		otp = "0" + otp
	}

	return otp
}
