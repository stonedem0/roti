package roti

import (
	"testing"
)

const (
	softPink = 213
	purple   = 57
)

//Test for a table builder
func TestTable_AddRow(t *testing.T) {
	table := Table{}
	stonedemo := []string{"https://stonedemo.wtf", "42s", "142ms", "4242ms"}
	apex := []string{"https://apex.sh", "42s", "142ms", "4242ms"}
	table.AddRow(stonedemo, purple, softPink, 3)
	table.AddRow(apex, purple, softPink, 3)
}
