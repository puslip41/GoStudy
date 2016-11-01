package wtmp

import (
	"errors"
	"encoding/binary"
)

const (
	UT_UNKNOWN=0
	RUN_LVL=1
	BOOT_TIME=2
	NEW_TIME=3
	OLD_TIME=4
	INIT_PROCESS=5
	LOGIN_PROCESS=6
	USER_PROCESS=7
	DEAD_PROCESS=8
	ACCOUNTING=9

	UT_LINESIZE=32
	UT_NAMESIZE=32
	UT_HOSTSIZE=256
)

type ExitStatus struct {
	ETermination uint16
	EExit uint16
}

type Utmp struct {
    UtType uint16
	UtPid uint32
	UtLine [UT_LINESIZE]byte
	UtId [4]byte
	UtUser [UT_NAMESIZE]byte
	UtHost [UT_HOSTSIZE]byte
	UtExit ExitStatus
	UtTv struct {
		TvSec uint32
		TvUsec uint32
	}
	UtSession uint32
	UtAddrV6 [4]uint32
	Unused [20]byte
}

func Unmashal(b []byte, utmp *Utmp) error {
	if b == nil || utmp == nil {
		return errors.New("Null Reference Exception")
	}

	uint16 := binary.LittleEndian.Uint16
	uint32 := binary.LittleEndian.Uint32

	utmp.UtType = uint16(b[:2])
	utmp.UtPid = uint32(b[4:8])

	return nil
}

func (utmp *Utmp) String() string {
	if utmp == nil {
		return ""
	}
	return ""
}

