package util

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
)

func Hash(name, pwd string) string {
	hash := md5.Sum([]byte(name + pwd + "mbx0nr/a9wWm45mMJmrUwAvugF+BYyiDO93CSDgCo+0"))
	return hex.EncodeToString(hash[:])
}

func I2B(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func B2I(b []byte) int64 {
	i := binary.BigEndian.Uint64(b)
	return int64(i)
}
