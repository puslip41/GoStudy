package main

import (
	"os"
	"fmt"
	"encoding/binary"
	"bytes"
	"unsafe"
	"log"
)


func main() {
	wtmpFilename := `D:\SourceCode\Go\src\github.com\puslip41\GoStudy\wtmp\cmd\wtmp`

	if len(os.Args) == 2 {
		wtmpFilename = os.Args[1]
	}

	file, err := os.OpenFile(wtmpFilename, os.O_RDONLY, os.FileMode(644))
	if err != nil {
		log.Println(err)
		os.Exit(0)
	}
	log.Println("open wtmp file:", wtmpFilename)
	defer file.Close()

	fmt.Printf("utmp size : %d, utmpx size : %d\n", unsafe.Sizeof(UTMP{}), unsafe.Sizeof(UTMPX{}))

	utmpx := UTMPX{}

	for {
		if err := binary.Read(file, binary.LittleEndian, &utmpx); err != nil {
			log.Println(err)
			break;
		}

		log.Println(fmt.Sprintf("ut_type:%d", utmpx.Ut_type),
			fmt.Sprintf("ut_pid:%d", utmpx.Ut_pid),
			fmt.Sprintf("ut_line:%s", utmpx.Ut_line),
			fmt.Sprintf("ut_id:%s", utmpx.Ut_id),
			fmt.Sprintf("ut_user:%s", utmpx.Ut_user),
			fmt.Sprintf("ut_host:%s", utmpx.Ut_host),
			fmt.Sprintf("e_termination:%d", utmpx.Ut_exit.E_termination),
			fmt.Sprintf("e_exit:%d", utmpx.Ut_exit.E_exit),
			fmt.Sprintf("ut_session:%d", utmpx.Ut_session),
			fmt.Sprintf("tv_sec:%d", utmpx.Ut_tv.Tv_sec),
			fmt.Sprintf("tv_usec:%d", utmpx.Ut_tv.Tv_usec),
			fmt.Sprintf("ut_addr_v6:%d", utmpx.Ut_addr_v6),
			fmt.Sprintf("unused:%X", utmpx.Unused),
		)
	}

	/*
	fmt.Println("utmp:", binary.Size(utmpx))
	fmt.Println("Ut_type:", binary.Size(utmpx.Ut_type))
	fmt.Println("Ut_pid:", binary.Size(utmpx.Ut_pid))
	fmt.Println("Ut_line:", binary.Size(utmpx.Ut_line))
	fmt.Println("Ut_id:", binary.Size(utmpx.Ut_id))
	fmt.Println("Ut_user:", binary.Size(utmpx.Ut_user))
	fmt.Println("Ut_host:", binary.Size(utmpx.Ut_host))
	fmt.Println("Ut_exit:", binary.Size(utmpx.Ut_exit))
	fmt.Println("Ut_session:", binary.Size(utmpx.Ut_session))
	fmt.Println("Ut_tv:", binary.Size(utmpx.Ut_tv))
	fmt.Println("Ut_addr_v6:", binary.Size(utmpx.Ut_addr_v6))
	fmt.Println("Unused:", binary.Size(utmpx.Unused))
	*/

	/*
	b := make([]byte, 384)
	for {
		_, err := file.Read(b)
		if err != nil {
			if err == io.EOF {
				break;
			} else {
				log.Println(err)
			}
		}

		buffer := bytes.NewBuffer(b)
		err = binary.Read(buffer, binary.LittleEndian, &utmpx)

		fmt.Printf("ut_type:%d\tut_pid:%d\tut_line:%s\tut_id:%s\t\r\n",
			utmpx.Ut_type,
			utmpx.Ut_pid,
			utmpx.Ut_line,
			utmpx.Ut_id,
		)
		*/

	/*
		b := make([]byte, 384)

		_, err = file.Read(b)
		if err != nil {
			if err == io.EOF {
				log.Println("read EOF")
				break;
			} else {
				log.Println(err)
			}
		}
		*/

	//log.Println(fmt.Sprintf("% X", b))
	/*
		fmt.Println(fmt.Sprintf("ut_type:%d", binary.LittleEndian.Uint16(b[:2])),
			fmt.Sprintf("ut_pid:%d", binary.LittleEndian.Uint32(b[2:6])),
			fmt.Sprintf("ut_line:%s(% X)", string(b[6:38]), b[6:38]),
			fmt.Sprintf("ut_id:%s(% X)", string(b[38:42]), b[38:42]),
			fmt.Sprintf("ut_user:%s(% X)", string(b[42:74]), b[42:74]),
			fmt.Sprintf("ut_host:%s", string(b[74:330])),
			fmt.Sprintf("e_termination:%d", binary.LittleEndian.Uint16(b[330:332])),
			fmt.Sprintf("e_exit:%d", binary.LittleEndian.Uint16(b[332:334])),
			fmt.Sprintf("ut_session:%d", binary.LittleEndian.Uint32(b[334:338])),
			fmt.Sprintf("tv_sec:%d", binary.LittleEndian.Uint32(b[338:342])),
			fmt.Sprintf("tv_usec:%d", binary.LittleEndian.Uint32(b[342:346])),
			fmt.Sprintf("ut_addr_v6:%d", binary.LittleEndian.Uint32(b[346:362])),
			fmt.Sprintf("unused:%X", b[362:382]),
		)
		*/
	/*
		log.Println(fmt.Sprintf("ut_type:%d(% X)", binary.LittleEndian.Uint32(b[:4]), b[:4]),
			fmt.Sprintf("ut_pid:%d(% X)", binary.LittleEndian.Uint32(b[4:8]), b[4:8]),
			fmt.Sprintf("ut_line:%s(% X)", string(b[8:40]), b[8:40]),
			fmt.Sprintf("ut_id:%s(% X)", string(b[40:44]), b[40:44]),
			fmt.Sprintf("ut_user:%s(% X)", string(b[44:76]), b[44:76]),
			fmt.Sprintf("ut_host:%s", string(b[76:332])),
			fmt.Sprintf("e_termination:%d", binary.LittleEndian.Uint16(b[332:334])),
			fmt.Sprintf("e_exit:%d", binary.LittleEndian.Uint16(b[334:336])),
			fmt.Sprintf("ut_session:%d", binary.LittleEndian.Uint32(b[336:340])),
			fmt.Sprintf("tv_sec:%d", binary.LittleEndian.Uint32(b[340:344])),
			fmt.Sprintf("tv_usec:%d", binary.LittleEndian.Uint32(b[344:348])),
			fmt.Sprintf("ut_addr_v6:%d", binary.LittleEndian.Uint32(b[348:364])),
			fmt.Sprintf("unused:%X", b[364:384]),
		)
		*/

}
