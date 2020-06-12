package models

import (
	"aging/methods"
	"sort"
)

//Find(el.Warehouse, el.Resouce, el.Color, el.Size, el.Style); index == len(*inv)

func (i *Inventories) Insert(x Inventory) {
	*i = append(*i, x)
}

func (i *Inventories) GroupBatches() {
	inv := new(Inventories)
	for _, el := range *i {
		el.GroupBatches()
		inv.Insert(el)
	}
	(*i).ReplaceBy(inv)
}

func (i *Inventories) SortBatches() {
	inv := new(Inventories)
	for _, el := range *i {
		sort.Sort(Lots(el.Batches))
		inv.Insert(el)
	}
	(*i).ReplaceBy(inv)
}

func (i *Inventories) Additions() {
	inv := new(Inventories)
	for _, el := range *i {
		for _, b := range el.Batches {
			if b.Value > 0 {
				if index := inv.Find(el.Warehouse, el.Resouce, el.Size, el.Color, el.Style); index == len(*inv) {
					*inv = append(*inv, Inventory{
						Warehouse: el.Warehouse,
						Resouce:   el.Resouce,
						Size:      el.Size,
						Color:     el.Color,
						Style:     el.Style,
						Batches:   Lots{b},
					})
				} else {
					inv.InsertBatchAt(index, b)
				}
			}
		}
	}
	(*i).ReplaceBy(inv)
}

func (i *Inventories) Reductions() {
	inv := new(Inventories)
	for _, el := range *i {
		for _, b := range el.Batches {
			if b.Value < 0 {
				if index := inv.Find(el.Warehouse, el.Resouce, el.Size, el.Color, el.Style); index == len(*inv) {
					*inv = append(*inv, Inventory{
						Warehouse: el.Warehouse,
						Resouce:   el.Resouce,
						Size:      el.Size,
						Color:     el.Color,
						Style:     el.Style,
						Batches:   Lots{b},
					})
				} else {
					inv.InsertBatchAt(index, b)
				}
			}
		}
	}
	(*i).ReplaceBy(inv)
}

// Do NOT USE IT, replace it by CloneFrom()
func (c *Inventories) Clone() Inventories {
	var s = make(Inventories, len(*c))
	copy(s, *c)
	return s
}

func (i *Inventories) CloneFrom(c Inventories) {
	inv := new(Inventories)
	for _, v := range c {
		batches := Lots{}
		for _, b := range v.Batches {
			batches = append(batches, Lot{
				Reference: b.Reference,
				Date:      b.Date,
				Key:       b.Key,
				Value:     b.Value,
			})
		}

		*inv = append(*inv, Inventory{
			Warehouse: v.Warehouse,
			Resouce:   v.Resouce,
			Size:      v.Size,
			Color:     v.Color,
			Style:     v.Style,
			Batches:   batches,
		})
	}
	(*i).ReplaceBy(inv)
}

// Find returns the smallest index i at which x == a[i],
// or len(a) if there is no such index.
func (inv *Inventories) Find(warehouse string, resouce string, size string, color string, style string) int {
	for i, in := range *inv {
		if warehouse == in.Warehouse && resouce == in.Resouce && size == in.Size && color == in.Color && style == in.Style {
			return i
		}
	}
	return len(*inv)
}

func (outs Inventories) BuildBatchesFrom(ins Inventories) (batchesBalance Inventories, batchesTransactions Inventories) {
	batchesBalance.CloneFrom(ins) // := ins.Clone()
	var outgoing Inventories
	outgoing.CloneFrom(outs) // := outs.Clone()
	batchesOut := Inventories{}

	for _, in := range batchesBalance {
		for _, out := range outgoing {
			if out.Warehouse == in.Warehouse && out.Resouce == in.Resouce && out.Size == in.Size && out.Color == in.Color && out.Style == in.Style {
				batches := Lots{}
			OUTER:
				for {
					oldestBatch := in.Batches.First()
					// batchQty := math.Min(in.Batches.First().Value, math.Abs(out.Batches.First().Value))  // This works with FLoat64
					batchQty := methods.MinOf(in.Batches.First().Value, methods.AbsOf(out.Batches.First().Value)) // This works with Int64

					batch := Lot{oldestBatch.Reference, out.Batches.First().Date, oldestBatch.Key, batchQty}
					batches = append(batches, batch)
					batchesTransactions = append(batchesTransactions, Inventory{
						Warehouse: out.Warehouse,
						Resouce:   out.Resouce,
						Size:      out.Size,
						Color:     out.Color,
						Style:     out.Style,
						Batches:   Lots{batch},
					})

					out.Batches[0].Value = out.Batches.First().Value + batchQty // because out batches already in -ve
					in.Batches[0].Value = oldestBatch.Value - batchQty

					if in.Batches.First().Value == 0 {
						in.Batches.PopFirst()
					}
					if out.Batches.First().Value == 0 {
						out.Batches.PopFirst()
						if len(out.Batches) == 0 || len(in.Batches) == 0 {
							break
						} else {
							continue OUTER
						}
					} else {
						continue OUTER
					}
				}
				batchesOut = append(batchesOut, Inventory{
					Warehouse: out.Warehouse,
					Resouce:   out.Resouce,
					Size:      out.Size,
					Color:     out.Color,
					Style:     out.Style,
					Batches:   batches,
				})
			}
		}
		//os.Exit(3)
	}
	return batchesBalance, batchesTransactions
}

func (i *Inventories) ReplaceBy(x *Inventories) {
	*i = *x
}

func (i *Inventories) InsertBatchAt(index int, b Lot) {
	(*i)[index].Batches.Insert(b)
	//a = append(a[:index+1], a[index:]...)
	//*i = append(*i, x)
}

func (i *Inventory) InsertBatch(x Lot) {
	(*i).Batches = append((*i).Batches, x)
}

func (i *Inventory) GroupBatches() {
	(*i).Batches.Group()
}
