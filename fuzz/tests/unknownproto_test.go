//go:build gofuzz || go1.18

package tests

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/v2/codec/unknownproto"
	"github.com/cosmos/cosmos-sdk/v2/testutil/testdata"
)

func FuzzUnknownProto(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		msg := new(testdata.TestVersion2)
		resolver := new(unknownproto.DefaultAnyResolver)
		_, _ = unknownproto.RejectUnknownFields(b, msg, true, resolver)

		_, _ = unknownproto.RejectUnknownFields(b, msg, false, resolver)
	})
}
