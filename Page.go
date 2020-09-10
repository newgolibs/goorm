package goorm

import (
	"strconv"
)

func (this *Page) SqlLImit() string {
	i := (this.PageID - 1) * this.Prepage
	if i <= 0 {
		i = 0
	}
	return " limit " + strconv.Itoa(i) + "," + strconv.Itoa(this.Prepage)
}
