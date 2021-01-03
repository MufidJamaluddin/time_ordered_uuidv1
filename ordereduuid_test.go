package OrderedUUIDv1

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name       string
	oldUuid    [16]byte
	wantNewUid [16]byte
}

func TestToOrderedUuid(t *testing.T) {

	var tests = []TestCase{
		{
			name:       "From UUID V1 to Ordered UUID",
			oldUuid:    UUIDStringToBytes("c5db1800-ce4c-11de-a5e2-1b45123c6e98"),
			wantNewUid: UUIDStringToBytes("11de-ce4c-c5db1800-a5e2-1b45123c6e98"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewUid := ToOrderedUuid(tt.oldUuid); !reflect.DeepEqual(gotNewUid, tt.wantNewUid) {
				t.Errorf(
					"ToOrderedUuid() = %v, want %v",
					UUIDBytesToString(gotNewUid),
					UUIDBytesToString(tt.wantNewUid))
			}
		})
	}
}

func TestFromOrderedUuid(t *testing.T) {

	var tests = []TestCase{
		{
			name:       "From Ordered UUID to UUID V1",
			oldUuid:    UUIDStringToBytes("11de-ce4c-c5db1800-a5e2-1b45123c6e98"),
			wantNewUid: UUIDStringToBytes("c5db1800-ce4c-11de-a5e2-1b45123c6e98"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewUid := FromOrderedUuid(tt.oldUuid); !reflect.DeepEqual(gotNewUid, tt.wantNewUid) {
				t.Errorf(
					"FromOrderedUuid() = %v, want %v",
					UUIDBytesToString(gotNewUid),
					UUIDBytesToString(tt.wantNewUid))
			}
		})
	}
}
