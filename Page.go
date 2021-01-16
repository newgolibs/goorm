package goorm

import (
	"math"
	"strconv"
)

func (this *Page) SqlLImit() string {
	if this.PageID <= 0 {
		this.PageID = 1
	}
	i := (this.PageID - 1) * this.Prepage
	if i <= 0 {
		i = 0
	}
	return " limit " + strconv.Itoa(i) + "," + strconv.Itoa(this.Prepage)
}

func (this *Page) SetTotal(Total int) {
	this.Total = Total
	this.TotalPages = int(math.Ceil(float64(Total / this.Prepage)))
}
