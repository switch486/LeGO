package lib

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	. "lib/datamodel"
	"os"
)

// PrintInventory prints the inventory
func PrintInventory(inventory INVENTORY) {
	for _, value := range inventory.Items {
		fmt.Println(value)
	}
}

// OpenFile reads the xml into a data structure
func OpenFile(path string) INVENTORY {
	// Open our xmlFile
	xmlFile, err := os.Open(path)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	data, _ := ioutil.ReadAll(xmlFile)

	var inventory INVENTORY

	_ = xml.Unmarshal([]byte(data), &inventory)

	inventory.PiecesMap = ComputePiecesMap(inventory)
	inventory.SetName = path

	return inventory
}

// ComputePiecesMap creates a map where
// key: ITEMTYPE#ITEMID#COLOR
// and
// value: MINQTY
func ComputePiecesMap(inventory INVENTORY) map[string]int {

	sampleMap := map[string]int{}

	for _, value := range inventory.Items {
		sampleMap[ComputeKey(value)] = value.Minqty
	}

	return sampleMap
}

// ComputeKey returns a map key
func ComputeKey(item ITEM) string {
	return item.Itemtype + "#" + item.Itemid + "#" + item.Color
}

// CalculateMissingPieces returns a difference: destination - source
func CalculateMissingPieces(source INVENTORY, destination INVENTORY) map[string]int {

	missingPieces := map[string]int{}

	for key, value := range destination.PiecesMap {
		if source.PiecesMap[key] < value {
			missingPieces[key] = value - source.PiecesMap[key]
		}
	}

	return missingPieces
}

// CalculateSetRelation computes the relation between the source and destination set
// which includes missing pieces, the count of missing pieces, and the NO of missing piece types
func CalculateSetRelation(source INVENTORY, destination INVENTORY) SetRelationResult {

	setRelation := SetRelationResult{}

	setRelation.MissingPiecesMap = CalculateMissingPieces(source, destination)
	setRelation.MissingPiecesCount = CalculatePiecesCount(setRelation.MissingPiecesMap)
	setRelation.MissingPieceTypesCount = len(setRelation.MissingPiecesMap)

	setRelation.SourceSetName = source.SetName
	setRelation.DestinationSetName = destination.SetName

	return setRelation
}

// CalculateCompoundSetRelations computes overall set relations between the destination inventory and many
// other sets presented in the source
func CalculateCompoundSetRelations(source []INVENTORY, destination INVENTORY) CompoundSetRelationResult {
	setRelations := CompoundSetRelationResult{}

	for _, value := range source {
		relation := CalculateSetRelation(value, destination)
		setRelations.AllResults[value.SetName] = relation
		if relation.MissingPiecesCount == 0 {
			setRelations.PerfectMatchResults[value.SetName] = relation
			setRelations.PerfectMatchCount++
		}
	}

	return setRelations
}

// CalculatePiecesCount returns the overall pieces count
func CalculatePiecesCount(piecesMap map[string]int) int {
	result := 0
	for _, value := range piecesMap {
		result += value
	}

	return result
}
