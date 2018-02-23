package dos

import (
	"testing"
)

func TestGetVolumeName(t *testing.T){
	s,err := GetVolumeName(`C:\`)
	if err != nil {
		t.Fatal(err)
		return
	}
	println(s)
}
