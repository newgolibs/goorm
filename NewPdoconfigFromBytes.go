package goorm

import (
	"encoding/json"
)

/**    从json的字符串中，生成数据库连接池对象    */
func NewPdoconfigFromBytes(bytes []byte) *Pdoconfig {
	var pdoconfig Pdoconfig
	// 配置还原成对象
	json.Unmarshal(bytes, &pdoconfig)
	return &pdoconfig
}
