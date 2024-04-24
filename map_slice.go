package utils

import (
	"encoding/json"
	"strings"
	"time"
)

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

// RemoveBlankElementInSlice 删除slice中空的值
func RemoveBlankElementInSlice(s []string) []string {
	var result []string
	for _, v := range s {
		if strings.TrimSpace(v) != "" {
			result = append(result, v)
		}
	}
	return result
}
