package server

import (
	"reflect"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	tc := []struct {
		in, out interface{}
	}{
		{int64(1), "1"},
		{float64(1.2345), "1.234"},
		{"1.2345", "\"1.2345\""},
		{time.Duration(time.Hour), "3600"},
		{"minpv", "\"minpv\""},
	}

	for _, tc := range tc {
		t.Logf("%+v", tc)
		out, err := encode(tc.in)
		if err != nil {
			t.Error(err)
		}

		if out != tc.out {
			t.Errorf("expected %v (string), got %v (%s)",
				tc.out, out, reflect.TypeOf(out).Kind(),
			)
		}
	}
}
