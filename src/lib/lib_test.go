package lib

import (
	. "lib/datamodel"
	. "lib/testsupport"
	"testing"
)

func TestReadXml(t *testing.T) {
	var inventory = OpenFile("testModel_2parts.xml")

	AssertEqualInt(t, 2, len(inventory.Items))
	AssertDifferentString(t, "", inventory.SetName)

	if inventory.PiecesMap == nil {
		t.Error("Pieces map has not been initialized")
	}
}

func TestComputeKey(t *testing.T) {
	b := ITEM{
		Itemtype: "P",
		Itemid:   "43723",
		Color:    "3",
		Minqty:   1,
	}

	computed := ComputeKey(b)

	expected := "P#43723#3"

	AssertEqualString(t, expected, computed)
}

func TestComputePiecesMap(t *testing.T) {
	//given
	source := getInventory(b(), c())

	//when
	computedMap := ComputePiecesMap(source)

	AssertEqualInt(t, 1, computedMap[ComputeKey(b())])
	AssertEqualInt(t, 2, computedMap[ComputeKey(c())])
}

func TestCalculateMissingPieces(t *testing.T) {
	//given
	source := getInventory(b(), c())
	destinationSet := getInventory(c(), d3())

	//when
	missingPieces := CalculateMissingPieces(source, destinationSet)

	//then
	AssertEqualInt(t, 1, len(missingPieces))
	AssertEqualInt(t, 3, missingPieces[ComputeKey(d3())])

}

func TestCalculatePiecesCount(t *testing.T) {
	//given
	computedMap := map[string]int{}

	computedMap["asd"] = 2
	computedMap["ddd"] = 3

	//then
	AssertEqualInt(t, 5, CalculatePiecesCount(computedMap))
}

func TestCalculateSetRelation(t *testing.T) {
	//given
	source := getInventory(b(), c())
	destinationSet := getInventory(c(), d3())

	//when
	setRelation := CalculateSetRelation(source, destinationSet)

	//then
	AssertEqualInt(t, 1, setRelation.MissingPieceTypesCount)
	AssertEqualInt(t, 3, setRelation.MissingPiecesCount)
	AssertEqualString(t, source.SetName, setRelation.SourceSetName)
	AssertEqualString(t, destinationSet.SetName, setRelation.DestinationSetName)
}

func TestCalculateSetRelationSelf(t *testing.T) {
	//given
	source := getInventory(e(), b(), d(), c())
	destination := getInventory(c(), b())

	//when
	setRelation := CalculateSetRelation(source, destination)

	//then
	AssertEqualInt(t, 0, setRelation.MissingPieceTypesCount)
	AssertEqualInt(t, 0, setRelation.MissingPiecesCount)
}

func TestCalculateSetRelationSubset(t *testing.T) {
	//given
	source := getInventory(b(), c())

	//when
	setRelation := CalculateSetRelation(source, source)

	//then
	AssertEqualInt(t, 0, setRelation.MissingPieceTypesCount)
	AssertEqualInt(t, 0, setRelation.MissingPiecesCount)
}

func getInventory(items ...ITEM) INVENTORY {
	invent := INVENTORY{}
	invent.Items = items
	name := ""
	for _, value := range items {
		name += ComputeKey(value) + ":"
	}
	invent.SetName = name
	invent.PiecesMap = ComputePiecesMap(invent)
	return invent
}

func b() ITEM {
	return ITEM{
		Itemtype: "P",
		Itemid:   "66",
		Color:    "3",
		Minqty:   1,
	}
}

func b3() ITEM {
	return ITEM{
		Itemtype: "P",
		Itemid:   "66",
		Color:    "3",
		Minqty:   3,
	}
}

func c() ITEM {
	return ITEM{
		Itemtype: "P",
		Itemid:   "77",
		Color:    "2",
		Minqty:   1,
	}
}

func c2() ITEM {
	return ITEM{
		Itemtype: "P",
		Itemid:   "77",
		Color:    "2",
		Minqty:   2,
	}
}

func d() ITEM {
	return ITEM{
		Itemtype: "P",
		Itemid:   "66",
		Color:    "2",
		Minqty:   1,
	}
}
func d3() ITEM {
	return ITEM{
		Itemtype: "P",
		Itemid:   "66",
		Color:    "2",
		Minqty:   3,
	}
}

func e() ITEM {
	return ITEM{
		Itemtype: "P",
		Itemid:   "667",
		Color:    "2",
		Minqty:   1,
	}
}

func e3() ITEM {
	return ITEM{
		Itemtype: "P",
		Itemid:   "667",
		Color:    "2",
		Minqty:   3,
	}
}
