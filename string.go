package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// StringToFloat64 string转float64
func StringToFloat64(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	return f
}

// BytesToReadable bytes转可读字符串
func BytesToReadable(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

// StringToInt string转int
func StringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return i
}

// RandomString 生成随机字符串
func RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// StringToUint string 转uint
func StringToUint(str string) uint {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return uint(i)
}
