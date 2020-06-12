package main

import (
	"aging/lib"
	"aging/methods"
	"aging/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Reading csv file thread > START //
	read := make(chan models.Data)
	// create csv file to hold the results > START //
	balanceQty, err := os.Create("balanceBatches.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	// create csv file to hold the results < END //
	// create csv file to hold the results > START //
	issuedQty, err := os.Create("IssuedBatches.csv")
	if err != nil {
		log.Fatal("Unable to open output")
	}
	// create csv file to hold the results < END //

	// Reading csv file thread > START //
	go func(input_file string) {
		var data models.Data
		var inventories = new(models.Inventories)

		fr, err := os.Open(input_file)
		methods.FailOnError(err)
		defer fr.Close()
		r := csv.NewReader(fr)
		rows, err := r.ReadAll()
		methods.FailOnError(err)
		data.Header = rows[0]
		for _, row := range rows[1:] {
			var inventory = models.Inventory{}
			inventory.Warehouse = strings.TrimSpace(row[5])
			inventory.Resouce = strings.TrimSpace(row[0])
			inventory.Color = strings.TrimSpace(row[1])
			inventory.Size = strings.TrimSpace(row[3])
			inventory.Style = strings.TrimSpace(row[2])
			TrxReference := strings.TrimSpace(row[9])
			TrxDate := lib.GetTime(time.Parse(lib.Custom, strings.TrimSpace(row[7])))
			TrxQty := lib.GetInt(strconv.ParseFloat(strings.TrimSpace(row[11]), 64)) // multiple by 1_000, then convert to int
			//lib.GetFloat(strconv.ParseFloat(strings.TrimSpace(row[11]), 64))
			TrxBatch := ""
			if TrxQty > 0 {
				TrxBatch = TrxDate.Format("2006-01")
			}
			inventory.Batches = models.Lots{models.Lot{
				Reference: TrxReference,
				Date:      TrxDate,
				Key:       TrxBatch,
				Value:     TrxQty,
			}}

			*inventories = append(*inventories, inventory)
		}

		data.Lines = *inventories

		read <- data // send data to channel read
	}("Transactions.csv")
	data := <-read // rceive from channel 'read' and assign value to new data variable
	// Reading csv file thread < END //

	fmt.Printf("\nRecords read: %v\n", len(data.Lines))

	// Handling csv file thread > START //
	done := make(chan int)
	go func(data models.Data, balance *os.File, issued *os.File) {
		defer balance.Close()
		defer issued.Close()

		remainders := csv.NewWriter(balance)
		trx := csv.NewWriter(issued)
		// handle header
		outs := data.Lines.Clone()
		outs.Reductions()
		ins := data.Lines.Clone()
		ins.Additions()

		batchesBalance, batchesTransactions := outs.BuildBatchesFrom(ins)

		// Write the outputs to files
		lib.WriteOutputs(batchesBalance, remainders)
		lib.WriteOutputs(batchesTransactions, trx)

		close(done)
	}(data, balanceQty, issuedQty)
	<-done // without this program may exit before finishing the goroutine
	// Handling csv file thread < END //
	// Handling csv file thread > START //
}
