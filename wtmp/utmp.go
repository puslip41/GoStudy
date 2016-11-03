package wtmp

import (
	"errors"
	"encoding/binary"
	"fmt"
	"bytes"
	"time"
	"net"
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
	EExit        uint16
}

type UtTv struct {
	TvSec uint32
	TvUsec uint32
}

type Utmp struct {
    UtType uint16
	UtPid uint32
	UtLine []byte
	UtId []byte
	UtUser []byte
	UtHost []byte
	UtExit ExitStatus
	UtTv UtTv
	UtSession uint32
	UtAddrV6 []byte
	Unused []byte
}

func Unmashal(b []byte, utmp *Utmp) error {
	if b == nil || utmp == nil {
		return errors.New("Null Reference Exception")
	}

	uint16 := binary.LittleEndian.Uint16
	uint32 := binary.LittleEndian.Uint32

	utmp.UtType = uint16(b[:2])
	utmp.UtPid = uint32(b[4:8])
	utmp.UtLine = b[8:40]
	utmp.UtId = b[40:44]
	utmp.UtUser = b[44:76]
	utmp.UtHost = b[76:332]
	utmp.UtExit.ETermination = uint16(b[334:336])
	utmp.UtExit.EExit = uint16(b[332:334])
	utmp.UtTv.TvSec = uint32(b[340:344])
	utmp.UtTv.TvUsec = uint32(b[336:340])
	utmp.UtSession = uint32(b[344:348])
	utmp.UtAddrV6 = b[348:364]
	utmp.Unused = b[364:384]

	return nil
}

func (utmp *Utmp) String() string {
	if utmp == nil {
		return ""
	}
	return fmt.Sprintf("ut_type:%d, ut_pid:%d, ut_line:%s, ut_id:%d, ut_user:%s, ut_host:%s, ut_exit.e_termination:%d, ut_exit.e_exit:%d, ut_tv:%s, ut_session:%d, ut_addr_v6:%s",
		utmp.UtType,
		utmp.UtPid,
		bytes.Trim(utmp.UtLine, "\x00"),
		binary.LittleEndian.Uint32(utmp.UtId),
		bytes.Trim(utmp.UtUser, "\x00"),
		bytes.Trim(utmp.UtHost, "\x00"),
		utmp.UtExit.ETermination,
		utmp.UtExit.EExit,
		time.Unix(int64(utmp.UtTv.TvSec), int64(utmp.UtTv.TvUsec)).Format("20060102150405"),
		utmp.UtSession,
		net.IP(bytes.Trim(utmp.UtAddrV6, "\x00")).String(),
	)
}

