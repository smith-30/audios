package main

import (
	"fmt"

	"github.com/winlinvip/go-fdkaac/fdkaac"
)

const BUFSIZE = 1024 // 読み込みバッファのサイズ

func main() {
	//file, err := os.Open("./adts_data.aac")
	//if err != nil {
	//	fmt.Println("e, ", err)
	//}
	//buf := make([]byte, BUFSIZE)
	//for {
	//	n, err := file.Read(buf)
	//	if n == 0 {
	//		break
	//	}
	//	if err != nil {
	//		// Readエラー処理
	//		break
	//	}
	//
	//	//fmt.Print(string(buf[:n]))
	//}

	d := fdkaac.NewAacDecoder()

	asc := []byte{0x12, 0x10}
	if err := d.InitRaw(asc); err != nil {
		fmt.Println(err)
		return
	}
	defer d.Close()

	if b, err := d.Decode([]byte{0x17, 0xf6, 0x65, 0x15, 0x00, 0x48, 0xa9, 0x80, 0x00, 0x38,
		0xff, 0xf1, 0x50, 0x80, 0x0c, 0x40, 0xfc,
		0x21, 0x17, 0x55, 0x35, 0xa1, 0x0c, 0x2f, 0x00, 0x00, 0x50, 0x23, 0xa6, 0x81, 0xbf, 0x9c, 0xbf}); err != nil {
		fmt.Println(b)
		return
	} else {
		fmt.Println(b)
	}
}
