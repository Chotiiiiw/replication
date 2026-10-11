package main

import "fmt"

type ShoppingCart struct {
	Items []string
}
type VersionCart struct {
	version int
	items   ShoppingCart
}
type DatabaseEntry struct {
	latestVersion int
	items         []VersionCart
}

func write(entry *DatabaseEntry, cart ShoppingCart, clientVersion int) {
	// 1. Increment the version counter
	entry.latestVersion++
	// 2. Keep versions the client hasn't observed
	var remaining []VersionCart

	for _, existing := range entry.items {
		if existing.version > clientVersion {
			remaining = append(remaining, existing)
		}
	}

	// 3. Add the new cart with its assigned version
	newVersion := VersionCart{
		version: entry.latestVersion,
		items:   cart,
	}
	remaining = append(remaining, newVersion)

	// 4. Update the database entry
	entry.items = remaining
}
func main() {
	// Create an empty database
	db := DatabaseEntry{
		latestVersion: 0,
		items:         []VersionCart{},
	}

	// Client A writes milk without observing any version
	write(&db, ShoppingCart{Items: []string{"milk"}}, 0)

	fmt.Println("After Client A:")
	fmt.Printf("%+v\n", db)

	// Client B writes eggs, also without observing Client A
	write(&db, ShoppingCart{Items: []string{"eggs"}}, 0)

	fmt.Println("After Client B:")
	fmt.Printf("%+v\n", db)

	// Client A has observed version 1, but not version 2
	write(&db, ShoppingCart{Items: []string{"milk", "flour"}}, 1)

	fmt.Println("After Client A writes again:")
	fmt.Printf("%+v\n", db)
}
