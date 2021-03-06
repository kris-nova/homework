package main

import (
	"fmt"

	"github.com/kris-nova/job2021/cs/junkdrawer"
	"github.com/kubicorn/kubicorn/pkg/namer"

	"github.com/kris-nova/job2021/cs/hash"
)

func main() {
	fmt.Println("<-----")
	// NewTable(512)
	table := hash.NewTable("Nóva", 512)

	// Set(small) X 10
	generateKeyValue(table, 10, 1024)
	// Set(large) X 10
	generateKeyValue(table, 2, 1024*1024)

	// Time (get(small, large))
	// assert small=small, large=large
	// 20 tests

	fmt.Println("Name: Nóva")
	table.Set("Name", "Nóva")
	fmt.Printf("Name: %s\n", table.Get("Name").String())

	fmt.Println("----->")
}

var keys []string

func generateKeyValue(table *hash.Table, itemCount int, itemSize int) error {
	for i := itemCount; i > 0; i-- {
		key := namer.RandomName()
		value := junkdrawer.RandomString(itemSize)
		fmt.Printf("%s : %d\n", key, itemSize)
		// Insert into table
		table.Set(key, value)
		keys = append(keys, key)
	}
	return nil
}
