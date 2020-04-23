package roti

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

// Table struct. User's data and formating sre stored here.
type Table struct {
	columns    [][]string
	rows       string
	colSizes   []int
	data       [][]string
	bg         uint8
	text       uint8
	padding    int
	headers    []string
	header     string
	headerBg   uint8
	headerText uint8
	margin     string
}

//Calculate and and save columns size based on user data
func (t *Table) saveColSizes(columns [][]string) []int {
	for _, c := range columns {
		for k, co := range c {
			if t.colSizes[k] < len(co) {
				t.colSizes[k] = len(co)
			}
		}
	}
	return t.colSizes
}

//AddRow adds row to a table
func (t *Table) AddRow(columns []string, textColor uint8, bgColor uint8, padding int) {
	t.data = append(t.data, columns)
	t.bg = bgColor
	t.text = textColor
	t.padding = padding
	t.colSizes = make([]int, len(columns))
	t.saveColSizes(t.data)

}

//AddHeader adds header
func (t *Table) AddHeader(headers []string, textColor uint8, bgColor uint8, padding int) {
	t.headers = append(headers)
	t.headerBg = bgColor
	t.headerText = textColor
	t.colSizes = make([]int, len(headers))
	t.saveColSizes(t.data)

}

// Helper method. All table formatting is happening here.
func (t *Table) formatHeader(headers []string, textColor uint8, bgColor uint8, padding int) string {
	p := " "
	header := ""
	header += aurora.Index(t.text, strings.Repeat(p, t.padding)).BgIndex(t.headerBg).String()
	for i, co := range headers {
		header += aurora.Index(t.headerText, co).BgIndex(t.headerBg).Bold().String() + aurora.Index(t.headerText, strings.Repeat(p, (t.colSizes[i]+padding)-len(co))).BgIndex(t.headerBg).String()
	}
	t.header += header
	return fmt.Sprintf("%v\n", t.header)
}

func (t *Table) addMargin(symbol string, color uint8, bg uint8) string {
	t.margin += aurora.Index(color, strings.Repeat(symbol, 23)).BgIndex(bg).String()
	t.margin += "\n"
	return fmt.Sprintf("%v", t.margin)
}

//Printing all rows
func (t *Table) formatTable() string {
	p := " "
	rows := ""
	header := t.formatHeader(t.headers, t.headerBg, t.headerText, t.padding)
	for _, c := range t.data {
		rows += aurora.Index(t.text, strings.Repeat(p, t.padding)).BgIndex(t.bg).String()
		for i, co := range c {
			rows += aurora.Index(t.text, co).BgIndex(t.bg).String() + aurora.Index(t.text, strings.Repeat(p, (t.colSizes[i]+t.padding)-len(co))).BgIndex(t.bg).String()
		}
		rows += "\n"
	}
	t.rows += rows
	return fmt.Sprintf("%v%v", header, t.rows)
}
