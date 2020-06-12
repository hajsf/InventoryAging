package models

import "time"

type Inventories []Inventory
type Inventory struct { //instead of: map[string]map[string]Pairs
	Resouce   string
	Color     string
	Style     string
	Size      string
	Warehouse string
	Batches   Lots
}
type Lots []Lot
type Lot struct {
	Reference string
	Date      time.Time
	Key       string
	Value     int64 // float64
}

// Identifier should be upper case letter to be exported, regardless it is a type, field, function, ..
type Data struct {
	Header []string
	Lines  Inventories
}
