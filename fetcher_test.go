package hotelaah

import "testing"

func TestGetAndExtract(t *testing.T) {
	f := NewFetcher("http://www.hotelaah.com/dijishi.html")
	pairs, err := f.GetAndExtract()
	if err != nil {
		t.Fatal(err)
	}
	// simple test
	if len(pairs) != 334 {
		t.Error("length not right\n")
	}
	if pairs[333].First != "遵义市" || pairs[333].Second != "贵州省" {
		t.Error("last element not right\n")
	}
}
