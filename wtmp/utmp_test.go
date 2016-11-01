package wtmp

import "testing"

func TestMashalInputNilByte (t *testing.T) {
	utmp := Utmp{}

	if err := Unmashal(nil, &utmp); err == nil {
		t.Error("Expected return erorr, got nil")
	}
}

func TestMashalInputNilUtmp (t *testing.T) {
	b := make([]byte, 10)

	if err := Unmashal(b, nil); err == nil {
		t.Error("Expected return error, got nil")
	}
}

func TestMashal(t *testing.T) {

	b := make([]byte, 384)
	n := copy(b, []byte{0x01, 0x00, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00})
	for ; n < 384; n++ {
		b[n] = 0x00
	}

	utmp := Utmp{}

	if err := Unmashal(b, &utmp); err != nil {
		t.Error("error: ", err)
	} else {
		if utmp.UtType != 1 {
			t.Error("Expected utmp.UtType 1, got: ", utmp.UtType)
		} else if utmp.UtPid != 2 {
			t.Error("Expected utmp.UtPid 2, got: ", utmp.UtPid)
		}
	}
}
