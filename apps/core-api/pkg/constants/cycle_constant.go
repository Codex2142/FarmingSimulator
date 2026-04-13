package constants

// CommodityType merepresentasikan jenis komoditas dalam siklus budidaya
type CommodityType string

const (
	Fish  CommodityType = "fish"
	Plant CommodityType = "plant"
)

// CycleStatus merepresentasikan status dari siklus budidaya
type CycleStatus string

const (
	Ongoing  CycleStatus = "ongoing"
	Finished CycleStatus = "finished"
	Failed   CycleStatus = "failed"
)
