package OrderedUUIDv1

func ToOrderedUuid(uid [16]byte) (newUid [16]byte) {
	var (
		i uint
		c uint = 0
	)
	for i = 6; i < 8; i, c = i+1, c+1 {
		newUid[c] = uid[i]
	}
	for i = 4; i < 6; i, c = i+1, c+1 {
		newUid[c] = uid[i]
	}
	for i = 0; i < 4; i, c = i+1, c+1 {
		newUid[c] = uid[i]
	}
	for i = 8; i < 16; i, c = i+1, c+1 {
		newUid[c] = uid[i]
	}
	return
}

func FromOrderedUuid(uid [16]byte) (newUid [16]byte) {
	var (
		i uint
		c uint = 0
	)
	for i = 4; i < 8; i, c = i+1, c+1 {
		newUid[c] = uid[i]
	}
	for i = 2; i < 4; i, c = i+1, c+1 {
		newUid[c] = uid[i]
	}
	for i = 0; i < 2; i, c = i+1, c+1 {
		newUid[c] = uid[i]
	}
	for i = 8; i < 16; i, c = i+1, c+1 {
		newUid[c] = uid[i]
	}
	return
}
