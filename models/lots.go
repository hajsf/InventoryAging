package models

import (
	"time"
)

// To make Lots sortable
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

// Below to enable sorting: sort.Sort(Lots(lots))
func (l Lots) Len() int           { return len(l) }
func (l Lots) Less(i, j int) bool { return (l[i].Date).Before(l[j].Date) } // { return l[i].Key < l[j].Key }
func (l Lots) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }

func (p *Lots) Insert(x Lot) {
	*p = append(*p, x)
}

func (p *Lots) ReplaceBy(x *Lots) {
	*p = *x
}

func (p *Lots) Map() map[string]int64 { //
	sum := make(map[string]int64) // make(map[string]float64)
	for _, el := range *p {
		sum[el.Key] = sum[el.Key] + el.Value
	}
	return sum
}
func (p *Lots) Group() {
	lots := new(Lots)
	lots.FromMap(p.Map())
	p.ReplaceBy(lots)
}

func (p *Lots) Pop(i int) {
	*p = append((*p)[:i], (*p)[i+1:]...) // append(a, b...)  // append(a[:i], a[i+1:]...)
}

func (p *Lots) PopLast() {
	*p = append((*p)[:len(*p)-1])
}

func (p *Lots) PopFirst() {
	*p = append((*p)[1:len(*p)])
}

func (p *Lots) First() Lot {
	return (*p)[0]
}

func (p *Lots) Last() Lot {
	return (*p)[len(*p)]
}
func (p *Lots) FromMap(m map[string]int64) { // float64) {
	for k, v := range m {
		(*p).Insert(Lot{"", time.Now(), k, v})
	}
}
