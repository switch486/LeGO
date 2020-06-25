package datamodel

import "encoding/xml"

type INVENTORY struct {
	XMLName   xml.Name `xml:"INVENTORY"`
	Items     []ITEM   `xml:"ITEM"`
	PiecesMap map[string]int
	SetName   string
}

type ITEM struct {
	Itemtype string `xml:"ITEMTYPE"`
	Itemid   string `xml:"ITEMID"`
	Color    string `xml:"COLOR"`
	Minqty   int    `xml:"MINQTY"`
}

type SetRelationResult struct {
	MissingPiecesMap       map[string]int
	MissingPiecesCount     int
	MissingPieceTypesCount int
	SourceSetName          string
	DestinationSetName     string
}

type CompoundSetRelationResult struct {
	AllResults          map[string]SetRelationResult
	PerfectMatchResults map[string]SetRelationResult
	PerfectMatchCount   int
}
