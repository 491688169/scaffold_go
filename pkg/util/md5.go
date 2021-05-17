/*
 * @Author: Kim
 * @Date: 2021-05-14 10:40:14
 * @LastEditTime: 2021-05-14 10:40:15
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /template_go/pkg/util/md5.go
 */
package util

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
