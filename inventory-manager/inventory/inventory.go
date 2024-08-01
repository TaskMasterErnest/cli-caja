package inventory

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type item struct {
	Product_ID string
	Name       string
	Quantity   int64
	Price      float64
}

type Inventory []item

func List(file io.Reader) Inventory {
	r := csv.NewReader(file)
	inventory := Inventory{}
	header := true

	for {
		item := item{}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}

		if !header {
			for idx, value := range record {
				switch idx {
				case 0:
					value = strings.TrimSpace(value)
					item.Product_ID = value
				case 1:
					value = strings.TrimSpace(value)
					item.Name = value
				case 2:
					value = strings.TrimSpace(value)
					item.Quantity, err = strconv.ParseInt(value, 10, 64)
					if err != nil {
						log.Fatal(err)
					}
				case 3:
					value = strings.TrimSpace(value)
					item.Price, err = strconv.ParseFloat(value, 64)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
			inventory = append(inventory, item)
		}
		header = false
	}
	return inventory
}

// func convertStringToInt64(value string) (int64, error) {
// 	return strconv.ParseInt(value, 10, 64)
// }
// func convertStringtoFloat64(value string) (float64, error) {
// 	return strconv.ParseFloat(value, 64)
// }

// func (i *Inventory) Add(id string, name string, quantity int64, price float64) {
// 	s := item{
// 		Product_ID: id,
// 		Name:       name,
// 		Quantity:   quantity,
// 		Price:      price,
// 	}

// 	*i = append(*i, s)
// }

func AppendItem(filename string, id string, name string, quantity int64, price float64) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	defer w.Flush()

	record := []string{
		id,
		name,
		strconv.FormatInt(quantity, 10),
		fmt.Sprintf("%.2f", price),
	}

	err = w.Write(record)
	if err != nil {
		return err
	}

	return nil

}

// func AddProduct(filename string, args []string) error {
// 	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	// write CSV record to file
// 	w := csv.NewWriter(file)
// 	defer w.Flush()

// 	err = w.Write(args)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func SellProduct(filename string, id string, quantity int) Inventory {
// 	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
// 	if err != nil {
// 		panic(err)
// 	}

// 	r := csv.NewReader(file)
// 	inventories := Inventory{}

// 	header := true

// 	for {
// 		inventory := inventory{}

// 	}

// 	return nil
// }
