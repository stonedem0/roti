package roti

import (
	"fmt"
	"strings"

	"github.com/logrusorgru/aurora"
)

// - keeping all state locally for String()
// - add cell{} struct with bg, fg colors and text string so each cell is customizable

// Table struct. User's data and formating are stored here.
type Table struct {
	columns  []string
	headers  []string
	rows     string
	data     [][]string
	bg       uint8
	fg       uint8
	padding  int
	bgHeader uint8
	fgHeader uint8
}

//Calculate and and save columns size based on user data
func saveColSizes(data [][]string, colSizes []int) []int {
	for _, c := range data {
		for k, co := range c {
			if colSizes[k] < len(co) {
				colSizes[k] = len(co)
			}
		}
	}
	return colSizes
}

//AddRow adds row to a table
func (t *Table) AddRow(columns []string, textColor uint8, bgColor uint8, padding int) {
	t.columns = append(columns)
	t.data = append(t.data, columns)
	t.bg = bgColor
	t.fg = textColor
	t.padding = padding
}

//AddHeader adds header
func (t *Table) AddHeader(headers []string, textColor uint8, bgColor uint8, padding int) {
	t.headers = append(headers)
	t.bgHeader = bgColor
	t.fgHeader = textColor
}

// String() string

// largestColumn
// formatHeader() string
// formatRows() string

// Helper method. All table formatting is happening here.
func (t *Table) formatHeader(headers []string, textColor uint8, bgColor uint8, padding int, colSizes []int) string {
	p := " "
	header := ""
	header += aurora.Index(t.fg, strings.Repeat(p, t.padding)).BgIndex(t.bgHeader).String()
	for i, co := range headers {
		header += aurora.Index(t.fgHeader, co).BgIndex(t.bgHeader).Bold().String() + aurora.Index(t.fgHeader, strings.Repeat(p, (colSizes[i]+padding)-len(co))).BgIndex(t.bgHeader).String()
	}
	return fmt.Sprintf("%v\n", header)
}

// func (t *Table) addMargin(symbol string, color uint8, bg uint8) string {
// 	t.margin += aurora.Index(color, strings.Repeat(symbol, 23)).BgIndex(bg).String()
// 	t.margin += "\n"
// 	return fmt.Sprintf("%v", t.margin)
// }

//Printing all rows
func (t *Table) String() string {
	p := " "
	rows := ""
	colSizes := make([]int, len(t.columns))
	saveColSizes(t.data, colSizes)
	header := t.formatHeader(t.headers, t.bgHeader, t.fgHeader, t.padding, colSizes)

	for _, c := range t.data {
		rows += aurora.Index(t.fg, strings.Repeat(p, t.padding)).BgIndex(t.bg).String()
		for i, co := range c {
			rows += aurora.Index(t.fg, co).BgIndex(t.bg).String() + aurora.Index(t.fg, strings.Repeat(p, (colSizes[i]+t.padding)-len(co))).BgIndex(t.bg).String()
		}
		rows += "\n"
	}

	return fmt.Sprintf("%v%v", header, rows)
}
