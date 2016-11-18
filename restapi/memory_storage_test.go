package restapi

import (
	"testing"
)

func TestRegister(t *testing.T) {
   cases := []struct {id, password, name, email string} {
		{"helloljho", "root1234", "lee jungho", "helloljho@igloosec.om"},
		{"puslip41", "root1234", "kim pilseop", "puslip41@igloosec.om"},
	}

	s := MemoryStorage{}

	for i, tc := range cases {
		err := s.Register(tc.id, tc.password, tc.name, tc.email)

		if err != nil {
			t.Errorf("MemoryStorage.Register() returns error. %s", err.Error())
		} else {
			if s.Count() != (i+1) {
				t.Errorf("expacted count: %d, got: %d", i+1, s.Count())
			}
		}
	}
}

func TestRegisterDuplicatedID(t *testing.T) {
	s := MemoryStorage{}

	s.Register("puslip41", "root1234", "kim pilseop", "puslip41@igloosec.om")

	if err := s.Register("puslip41", "root1234", "kim pilseop", "puslip41@igloosec.om"); err != nil {
		if err.Error() != "exist id: puslip41" {
			t.Errorf("expacted error: \"exist id: puslip41\", got \"%s\"", err.Error())
		}
	} else {
		t.Errorf("expacted error but not returns")
	}
}

func Test_Delete(t *testing.T) {
	s := MemoryStorage{}

	s.Register("puslip41", "root1234", "kim pilseop", "puslip41@igloosec.om")

	if err := s.Delete("puslip41", "root1234"); err != nil {
		t.Errorf("unexpacted error: %s", err.Error())
	} else {
		if s.Count() != 0 {
			t.Errorf("expacted count: 0, got: %d", s.Count())
		}
	}
}

func Test_DeleteNotEqualsPassword(t *testing.T) {
	s := MemoryStorage{}

	s.Register("puslip41", "root1234", "kim pilseop", "puslip41@igloosec.om")

	if err := s.Delete("puslip41", "root"); err == nil {
		t.Errorf("expacted error. but not returns")
	} else {
		if err.Error() != "not equals password" {
			t.Errorf("expacted error(not equals password), got: error(%s)", err.Error())
		}
	}
}

func Test_Modify(t *testing.T) {
	s := MemoryStorage{}

	s.Register("puslip41", "root1234", "lee jungho", "helloljho@igloosec.com")

	if err := s.Modify("puslip41", "root1234", "kim pilseop", "puslip41@igloosec.com"); err != nil {
		t.Errorf("unexpacted error: %s", err.Error())
	} else {
		if name, email, err := s.Query("puslip41"); err != nil {
			t.Errorf("unexpacted error: %s", err.Error())
		} else {
			if name != "kim pilseop" {
				t.Errorf("expacted : \"kim pilseop\", got: \"%s\"", name)
			}

			if email != "puslip41@igloosec.com" {
				t.Errorf("expacted : \"puslip41@igloosec.com\", got: \"%s\"", email)
			}
		}

	}
}

func Test_ModifyNotequalsPassword(t *testing.T) {
	s := MemoryStorage{}

	s.Register("puslip41", "root1234", "lee jungho", "helloljho@igloosec.com")

	if err := s.Modify("puslip41", "root", "kim pilseop", "puslip41@igloosec.com"); err == nil {
		t.Errorf("expacted error. but not returns")
	} else {
		if err.Error() != "not equals password" {
			t.Errorf("expacted error(not equals password), got: error(%s)", err.Error())
		}
	}
}

func Test_Query(t *testing.T) {
	cases := []struct {id, password, name, email string} {
		{"helloljho", "root1234", "lee jungho", "helloljho@igloosec.om"},
		{"puslip41", "root1234", "kim pilseop", "puslip41@igloosec.om"},
	}

	s := MemoryStorage{}

	for _, tc := range cases {
		s.Register(tc.id, tc.password, tc.name, tc.email)
	}

	item := cases[1]

	if name, email, err := s.Query(item.id); err != nil {
		t.Errorf("unexpacted error: %s", err.Error())
	} else {
		if name != item.name {
			t.Errorf("expacted : \"%s\", got: \"%s\"", item.name, name)
		}

		if email != item.email {
			t.Errorf("expacted : \"%s\", got: \"%s\"", item.email, email)
		}
	}
}

func Test_Count(t *testing.T) {
	s := MemoryStorage{}

	if s.Count() != 0 {
		t.Errorf("expacted : 0, got: %d", s.Count())
	}
}
