package  main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

type Website struct {
	Url int32
}

func main()  {
	//writeBinFile()
	readBinFile()
}

func readBinFile() {
	file, err := os.Open("demo.bin")
	defer file.Close()
	if err != nil {
		fmt.Println("文件打开失败", err.Error())
		return
	}
	m := Website{}
	for i := 1; i <= 10; i++ {
		data := readNextBytes(file, 4)
		buffer := bytes.NewBuffer(data)
		err = binary.Read(buffer, binary.LittleEndian, &m)
		if err != nil {
			fmt.Println("二进制文件读取失败", err)
			return
		}
		fmt.Println("第", i, "个值为：", m)
	}
}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		fmt.Println("解码失败", err)
	}
	return bytes
}

func writeBinFile() {
	file,err := os.Create("demo.bin")

	for i:=1;i<10;i++{
		info:= Website{
			int32(i),
		}
		if err!=nil{
			fmt.Println("创建文件失败")
		}
		defer file.Close()
		var binBuf bytes.Buffer
		binary.Write(&binBuf,binary.LittleEndian,info)

		b:=binBuf.Bytes()
		_,err =file.Write(b)

		if err!=nil{
			fmt.Println("编码失败", err.Error())
			return
		}
	}
    fmt.Println("编码成功")
}