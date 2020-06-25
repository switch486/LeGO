package main

import lib "lib"

func main() {
	// available pieces:
	// the pieces that are at stock
	var inventory = lib.OpenFile("../models/31011.xml")

	lib.PrintInventory(inventory)

	// official sets that can be built
	// all the sets that are available - to check if they can be build with the pieces on stock

}
