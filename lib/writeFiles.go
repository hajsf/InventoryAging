package lib

import (
	"aging/models"
	"encoding/csv"
	"log"
	"strconv"
)

// Write the data to the selected csv file
func WriteOutputs(data models.Inventories, file *csv.Writer) {
	header := []string{"Warehouse", "Resouce", "Size", "Color", "Style", "Batch.Reference", "Batch.Date", "Batch.Key", "Balance Qty"}
	if err := file.Write(header); err != nil {
		log.Fatal(err)
	}
	// Write file
	for _, inv := range data { // Inventories === []Invntory
		for _, batch := range inv.Batches {
			record := []string{
				inv.Warehouse,
				inv.Resouce,
				inv.Size,
				inv.Color,
				inv.Style,
				batch.Reference,
				batch.Date.String(),
				batch.Key,
				strconv.FormatFloat(float64(batch.Value)/1000, 'f', -1, 64), // strconv.FormatFloat(batch.Value, 'f', -1, 64), // strconv.FormatInt(batch.Value, 64)
				//batch.Value,
			}
			// out := append(data.show(), "100%")
			if err := file.Write(record); err != nil {
				log.Fatal(err)
			}
		}
	}
	file.Flush()
}
