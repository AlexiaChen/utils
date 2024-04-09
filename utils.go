package utils

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/gctx"
)

func GetConfig(confPath string, key string) *g.Var {
	var ctx = gctx.New()
	adapter, err := gcfg.NewAdapterFile(confPath)
	if err != nil {
		panic(err)
	}
	config := gcfg.NewWithAdapter(adapter)
	val, err := config.Get(ctx, key)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return val
}

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

// FormatAmount 格式化金额字符串
func FormatAmount(amount *big.Float) string {
	return amount.Text('f', 2)
}

// ToMap 将结构化数据转为Map
func ToMap(data interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	j, _ := json.Marshal(data)
	err := json.Unmarshal(j, &m)
	if _, ok := m["CreatedAt"]; ok {
		createTime, _ := time.Parse(time.RFC3339, m["CreatedAt"].(string))
		m["create_time"] = createTime.Format("2006-01-02 15:04:05")
	}
	if _, ok := m["UpdatedAt"]; ok {
		updateTime, _ := time.Parse(time.RFC3339, m["UpdatedAt"].(string))
		m["update_time"] = updateTime.Format("2006-01-02 15:04:05")
	}
	m["id"] = m["ID"]
	delete(m, "password")
	delete(m, "Password")
	delete(m, "CreatedAt")
	delete(m, "UpdatedAt")
	delete(m, "ID")
	delete(m, "DeletedAt")
	return m, err
}

// IsValidIP 验证IP是否合法
func IsValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// ValidateIpList 验证IP列表是否合法
func ValidateIpList(ipList string) bool {
	if strings.TrimSpace(ipList) == "" {
		return false
	}
	ips := strings.Split(ipList, "\n")
	for _, ip := range ips {
		if strings.TrimSpace(ip) != "" && !IsValidIP(strings.TrimSpace(ip)) {
			return false
		}
	}
	return true
}

func RemoveBlankElementInSlice(s []string) []string {
	var result []string
	for _, v := range s {
		if strings.TrimSpace(v) != "" {
			result = append(result, v)
		}
	}
	return result
}
