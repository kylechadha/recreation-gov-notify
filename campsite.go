package main

type Campsites struct {
	Campsites map[string]Campsite `json:"campsites"`
	Count     int                 `json:"count"`
}

type Campsite struct {
	Availabilities      map[string]string `json:"availabilities"`
	CampsiteID          string            `json:"campsite_id"`
	CampsiteReserveType string            `json:"campsite_reserve_type"`
	CampsiteType        string            `json:"campsite_type"`
	Loop                string            `json:"loop"`
	MaxNumPeople        int               `json:"max_num_people"`
	MinNumPeople        int               `json:"min_num_people"`
	Site                string            `json:"site"`
	TypeOfUse           string            `json:"type_of_use"`
}
