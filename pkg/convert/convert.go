/*
 * @Author: Kim
 * @Date: 2021-03-08 16:40:50
 * @LastEditTime: 2021-03-08 16:50:43
 * @LastEditors: Kim
 * @Description:
 * @FilePath: /scaffold_go/pkg/convert/convert.go
 */
package convert

import "strconv"

type StrTo string

func (s StrTo) String() string {
	return s.String()
}

func (s StrTo) ToInt() (int, error) {
	v, err := strconv.Atoi(s.String())
	return v, err
}

func (s StrTo) MustToInt() int {
	v, _ := s.ToInt()
	return v
}

func (s StrTo) ToUint32() (uint32, error) {
	v, err := strconv.Atoi(s.String())
	return uint32(v), err
}

func (s StrTo) MustToUint32() uint32 {
	v, _ := s.ToUint32()
	return v
}
