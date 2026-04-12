package constants

// SubPlaceType merepresentasikan tipe dari sub_place/lokasi budidaya
type SubPlaceType string

const (
	Pond        SubPlaceType = "pond"
	PlantBed    SubPlaceType = "plant_bed"
	Greenhouse  SubPlaceType = "greenhouse"
)

// SubPlaceStatus merepresentasikan status operasional sub_place
type SubPlaceStatus string

const (
	Active   SubPlaceStatus = "active"
	Inactive SubPlaceStatus = "inactive"
)
