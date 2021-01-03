package OrderedUUIDv1

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func UUIDStringToBytes(suid string) (uuid [16]byte) {
	var (
		ns string
		nb []byte
	)
	ns = strings.ReplaceAll(suid, "-", "")
	nb = append(nb, ns...)
	_, _ = hex.Decode(uuid[:], nb)
	return
}

func UUIDBytesToString(uuid [16]byte) (ns string) {

	ns = hex.EncodeToString(uuid[:])

	ns = fmt.Sprintf("%v-%v-%v-%v-%v",
		ns[0:8], ns[8:12], ns[12:16], ns[16:20], ns[20:])

	return
}
