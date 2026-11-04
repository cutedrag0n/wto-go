package qr

import (
	"encoding/json"
	"os"
	"testing"
)

func TestWireScalars(t *testing.T) {
	for _, wire := range []string{`3`, `"3"`} {
		var v Integer
		if err := json.Unmarshal([]byte(wire), &v); err != nil || v != 3 {
			t.Fatal(v, err)
		}
	}
	for _, wire := range []string{`true`, `1`, `"1"`} {
		var v Boolean
		if err := json.Unmarshal([]byte(wire), &v); err != nil || !v {
			t.Fatal(v, err)
		}
	}
	for _, wire := range []string{`false`, `0`, `"0"`, `null`} {
		var v Boolean
		if err := json.Unmarshal([]byte(wire), &v); err != nil || v {
			t.Fatal(v, err)
		}
	}
	var i Integer
	if err := json.Unmarshal([]byte(`"x"`), &i); err == nil {
		t.Fatal("invalid integer")
	}
	var b Boolean
	if err := json.Unmarshal([]byte(`2`), &b); err == nil {
		t.Fatal("invalid boolean")
	}
}
func TestCurrentQRShape(t *testing.T) {
	var p Page[QR]
	err := json.Unmarshal([]byte(`{"data":[{"id":5487,"reporter_member":{"code":"C004"},"products":[{"product_code":"010310","hs_version":"H6"}],"measures":[{"interpreted":0,"mast_codes":["E311"]}],"termination_dt":null}],"links":{"next":null},"meta":{"current_page":1,"total":1}}`), &p)
	if err != nil {
		t.Fatal(err)
	}
	if len(p.Data) != 1 || p.Data[0].ReporterMember.Code != "C004" || p.Data[0].Products[0].ProductCode != "010310" || p.Meta.Total != 1 {
		t.Fatal(p)
	}
}

func TestOfficialFixtures(t *testing.T) {
	for _, tc := range []struct {
		file   string
		target any
	}{
		{"members", &Page[Member]{}}, {"products", &Page[Product]{}}, {"notifications", &Page[Notification]{}},
		{"qrs", &Page[QR]{}}, {"qrs_qrId", &Detail{}}, {"hs-versions", &Page[HSVersion]{}},
	} {
		t.Run(tc.file, func(t *testing.T) {
			b, err := os.ReadFile("testdata/" + tc.file + ".json")
			if err != nil {
				t.Fatal(err)
			}
			if err := json.Unmarshal(b, tc.target); err != nil {
				t.Fatal(err)
			}
		})
	}
}
