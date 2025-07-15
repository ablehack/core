package str

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// 传入ASCII字符串，返回其对应的HEX字符串
func ATH(s string) string {
	var hexString string
	for _, char := range s {
		hexString += hex.EncodeToString([]byte(string(char)))
	}
	return hexString
}

// 传入HEX字符串, 返回其对应的ASCII字符串, 无法匹配时使用\xFF形式转义
func HTA(s string) string {
	var result []rune
	for i := 0; i < len(s); i += 2 {
		if i+1 >= len(s) {
			result = append(result, rune(0xFF)) // 处理奇数长度的情况
			continue
		}
		byteStr := s[i : i+2]
		bytes, err := hex.DecodeString(byteStr)
		if err != nil || len(bytes) == 0 {
			result = append(result, []rune(fmt.Sprintf("\\x%02X", 0xFF))...)
		} else {
			char := bytes[0]
			if char >= 32 && char <= 126 { // 可打印ASCII字符
				result = append(result, rune(char))
			} else {
				result = append(result, []rune(fmt.Sprintf("\\x%02X", char))...)
			}
		}
	}
	return string(result)
}

// 传入整数，返回其对应的HEX字符串
func ITH(n int) string {
	return fmt.Sprintf("%02x", n)
}

// 传入HEX字符串，返回其对应的整数
func HTI(s string) (int, error) {
	num, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0, err
	}
	return int(num), nil
}
