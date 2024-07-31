package inventory

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type inventory struct {
	Product_ID string
	Name       string
	Quantity   int64
	Price      float64
}

type Inventory []inventory

func ListContent(inventoryFile io.Reader) Inventory {
	r := csv.NewReader(inventoryFile)
	inventories := Inventory{}

	header := true

	for {
		inventory := inventory{}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if !header {
			for idx, value := range record {
				switch idx {
				case 0:
					value = strings.TrimSpace(value)
					inventory.Product_ID = value
				case 1:
					value = strings.TrimSpace(value)
					inventory.Name = value
				case 2:
					value = strings.TrimSpace(value)
					inventory.Quantity, err = strconv.ParseInt(value, 10, 64)
					if err != nil {
						log.Fatal(err)
					}
				case 3:
					value = strings.TrimSpace(value)
					inventory.Price, err = strconv.ParseFloat(value, 64)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			inventories = append(inventories, inventory)
		}
		header = false
	}
	return inventories
}

// func (i *Inventory) Add(id string, name string, quantity int, price float64) {
// 	s := inventory{
// 		Product_ID: id,
// 		Name:       name,
// 		Quantity:   int64(quantity),
// 		Price:      float64(price),
// 	}

// 	*i = append(*i, s)
// }

func AddContent(filename string, args []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// write CSV record to file
	w := csv.NewWriter(file)
	defer w.Flush()

	err = w.Write(args)
	if err != nil {
		return err
	}

	return nil
}
