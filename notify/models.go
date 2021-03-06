package notify

type Campground struct {
	CampsiteReserveType []string `json:"campsite_reserve_type"`
	CampsiteTypeOfUse   []string `json:"campsite_type_of_use"`
	City                string   `json:"city"`
	CountryCode         string   `json:"country_code"`
	EntityID            string   `json:"entity_id"`
	EntityType          string   `json:"entity_type"`
	IsInventory         bool     `json:"is_inventory"`
	Lat                 string   `json:"lat"`
	Lng                 string   `json:"lng"`
	Name                string   `json:"name"`
	OrgID               string   `json:"org_id"`
	OrgName             string   `json:"org_name"`
	ParentEntityID      string   `json:"parent_entity_id"`
	ParentEntityType    string   `json:"parent_entity_type"`
	ParentName          string   `json:"parent_name"`
	PreviewImageURL     string   `json:"preview_image_url"`
	Reservable          bool     `json:"reservable"`
	StateCode           string   `json:"state_code"`
	Text                string   `json:"text"`
}

type Campsite struct {
	Availabilities      map[string]string `json:"availabilities"`
	CampsiteID          string            `json:"campsite_id"`
	CampsiteReserveType string            `json:"campsite_reserve_type"`
	CampsiteRules       interface{}       `json:"campsite_rules"`
	CampsiteType        string            `json:"campsite_type"`
	CapacityRating      string            `json:"capacity_rating"`
	Loop                string            `json:"loop"`
	MaxNumPeople        int               `json:"max_num_people"`
	MinNumPeople        int               `json:"min_num_people"`
	Quantities          interface{}       `json:"quantities"`
	Site                string            `json:"site"`
	TypeOfUse           string            `json:"type_of_use"`
}
