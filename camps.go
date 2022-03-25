package main

const availabilityURL = "https://www.recreation.gov/api/camps/availability/campground/%d/month"

type AvailabilityResponse struct {
	Campsites map[string]Campsite `json:"campsites"`
	Count     int                 `json:"count"`
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
