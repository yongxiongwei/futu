package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"futuxxx/futupb/InitConnect"
	"futuxxx/something"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:11111")
	if err != nil {
		fmt.Println("cannot connect server")
		os.Exit(1)
	}
	defer conn.Close()

	handlerWrite(conn)
	handlerRead(conn)

}

func handlerWrite(conn net.Conn) {

	pack := &futupack.FutuPack{}
	pack.SetProtoID(uint32(1001))
	// pack.SetBody([]byte("现在时间是:" + time.Now().Format("2006-01-02 15:04:05")))
	// pack.SetBody([]byte("test msg from client"))

	ClientVer := int32(307)
	ClientID := "test123456"

	fut := &InitConnect.Request{
		C2S: &InitConnect.C2S{
			ClientID:  &ClientID,
			ClientVer: &ClientVer,
		},
	}
	// fut.ClientVer = &ClientVer
	// fut.ClientID = &ClientID
	body, err := proto.Marshal(fut)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	pack.SetBody(body)

	pack.Pack(conn)

}

func handlerRead(conn net.Conn) {

	// for {
	// scanner
	scanner := bufio.NewScanner(conn)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		if !atEOF && data[0] == 'F' {
			if len(data) > 44 {
				length := uint32(0)
				binary.Read(bytes.NewReader(data[12:16]), binary.LittleEndian, &length)
				if int(length)+4 <= len(data) {
					return int(length) + 44, data[:int(length)+44], nil
				}
			}
		}
		return
	})

	for scanner.Scan() {
		scannedPack := new(futupack.FutuPack)
		err := scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		if err != nil {
			fmt.Println("unpack error", err)
			return
		}

		// fmt.Println(scannedPack)
		// return

		body := scannedPack.GetBody()

		fut := &InitConnect.Response{}
		err = proto.Unmarshal(body, fut)
		if err != nil {
			log.Fatal("unmarshaling error:", err)
		}

		fmt.Println("RetType: ", *fut.RetType, *(*fut.S2C).KeepAliveInterval)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("无效数据包")
	}

	// }
}
