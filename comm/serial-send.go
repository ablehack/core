package comm

import (
	"encoding/hex"
	"fmt"
	"io"

	"github.com/jacobsa/go-serial/serial"
)

// SendSerial 发送数据到串口
func SendSerial(com string, baud uint, bits uint, stop uint, parity uint, data string) error {
	// 配置串口参数
	options := serial.OpenOptions{
		PortName:              com,
		BaudRate:              baud,
		DataBits:              bits,
		StopBits:              stop,
		InterCharacterTimeout: 0,
		MinimumReadSize:       1,
	}

	// 打开串口
	port, err := serial.Open(options)
	if err != nil {
		return err
	}
	defer port.Close()

	// 将十六进制字符串解码为字节数组
	byteData, err := hex.DecodeString(data)
	if err != nil {
		return fmt.Errorf("error decoding hex string: %v", err)
	}

	// 发送数据
	err = writeData(port, []byte(byteData))
	if err != nil {
		return err
	}
	return nil
}

func writeData(port io.Writer, data []byte) error {
	_, err := port.Write(data)
	if err != nil {
		return err
	}
	return nil
}
