package cmd

import (
	"testing"
)

const (
	src   = "123\rabc\n456\r\nefg\n\r..."
	dcr   = "123\rabc\r456\refg\r\r..."
	dlf   = "123\nabc\n456\nefg\n\n..."
	dcrlf = "123\r\nabc\r\n456\r\nefg\r\n\r\n..."
)

func TestToCR(t *testing.T) {
	if d := ToCR(src); d != dcr {
		t.Errorf("ToCR(%s)=%x, but %x got", src, dcr, d)
	}
}
func TestToLF(t *testing.T) {
	if d := ToLF(src); d != dlf {
		t.Errorf("ToLF(%s)=%x, but %x got", src, dlf, d)
	}
}

func TestToCRLF(t *testing.T) {
	if d := ToCRLF(src); d != dcrlf {
		t.Errorf("ToCR(%s)=%x, but %x got", src, dcrlf, d)
	}
}
