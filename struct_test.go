package utils

import (
	"fmt"
	"testing"
)

type createReq struct {
	Cpu        uint32 `json:"cpu"`       //cpu核心数
	Memory     uint32 `json:"memory"`    //内存大小
	Bandwidth  uint32 `json:"bandwidth"` //带宽 后续没有带宽 此字段作为iops的字段
	DBPassword string `json:"db_password"`
	DBUser     string `json:"db_user"`
	Disks      uint32 `json:"disks"`      //数据盘大小
	HardDisks  uint32 `json:"hard_disks"` //系统盘大小
	Months     uint32 `json:"months"`
}

type createService struct {
	Bandwidth   uint32 `json:"bandwidth"` //带宽 后续没有带宽 此字段作为iops的字段
	Cpu         uint32 `json:"cpu"`       //cpu核心数
	DBPassword  string `json:"db_password"`
	DBUser      string `json:"db_user"`
	Disks       uint32 `json:"disks"`      //数据盘大小
	HardDisks   uint32 `json:"hard_disks"` //系统盘大小
	Memory      uint32 `json:"memory"`     //内存大小
	Months      uint32 `json:"months"`
	DBName      string `json:"db_name"`
	DBCharset   string `json:"db_charset"`
	DBCollation string `json:"db_collation"`

	//回调参数
	Id         string `json:"id"`
	TimeStamp  string `json:"time_stamp"`
	NonceStr   string `json:"nonce_str"`
	Sign       string `json:"sign"`
	Page       int    `json:"page"`
	Size       int    `json:"size"`
	Search     string `json:"search"`
	Sid        string `json:"sid"` //临时使用存放登录信息
	TemplateId string `json:"template_id"`
	ImportFile string `json:"import_file"`
	SqlText    string `json:"sql_text"`
}

func TestCopyStructUsingJSON(t *testing.T) {
	var a createReq
	b := &createService{}
	a.Months = 111
	a.Disks = 111
	a.Memory = 1
	if err := CopyStructUsingJSON(a, b); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", b)
}

func TestIPv4InSubnet(t *testing.T) {
	ipAddress := "192.168.1.10"
	subnetCIDR := "192.168.1.0/24"

	isInSubnet := IPv4InSubnet(ipAddress, subnetCIDR)
	if !isInSubnet {
		fmt.Println("Error checking subnet membership:")
	} else {
		fmt.Printf("IP address %s is %t in subnet %s\n", ipAddress, isInSubnet, subnetCIDR)
	}
}
