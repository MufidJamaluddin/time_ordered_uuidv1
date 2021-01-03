package OrderedUUIDv1

import (
	"strings"
	"testing"
)

func TestUtils(t *testing.T) {
	t.Run("Utility Test", func(t *testing.T) {
		uuid1 := "c5db1800-ce4c-11de-a5e2-1b45123c6e98"
		t.Log(uuid1)

		uBytes := UUIDStringToBytes(uuid1)
		t.Log(uBytes)

		uuid2 := UUIDBytesToString(uBytes)
		t.Log(uuid2)

		if strings.Compare(uuid1, uuid2) != 0 {
			t.Errorf("Utils UUID %v != %v", uuid1, uuid2)
		}
	})
}
