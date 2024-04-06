package main

type UniverseResponse struct {
	Name     string          `json:"name"`
	Ship     Ship            `json:"ship"`
	Universe [][]interface{} `json:"universe"`
}

type UniverseEntry struct {
	PlanetFrom string `json:"planetFrom"`
	PlanetTo   string `json:"planetTo"`
	Distance   int    `json:"distance"`
}

type Ship struct {
	CapacityX int         `json:"capacityX"`
	CapacityY int         `json:"capacityY"`
	FuelUsed  int         `json:"fuelUsed"`
	Garbage   interface{} `json:"garbage"` // В данном примере типы данных "garbage" и "planet" не определены, можно использовать пустой интерфейс для их представления
	Planet    interface{} `json:"planet"`
}

type TravelRequest struct {
	Planets []string `json:"planets"`
}

type TravelResponse struct {
	FuelDiff      int                      `json:"fuelDiff"`
	PlanetDiffs   []map[string]interface{} `json:"planetDiffs"`
	PlanetGarbage map[string][][]int       `json:"planetGarbage"`
	ShipGarbage   map[string][][]int       `json:"shipGarbage"`
}

type CollectRequest struct {
	Garbage map[string][][]int `json:"garbage"`
}

type CollectResponse struct {
	Garbage map[string][][]int `json:"garbage"`
	Leaved  []string           `json:"leaved"`
}

type Round struct {
	StartAt     string `json:"startAt"`
	EndAt       string `json:"endAt"`
	IsCurrent   bool   `json:"isCurrent"`
	Name        string `json:"name"`
	PlanetCount int    `json:"planetCount"`
}

type RoundsResponse struct {
	Rounds []Round `json:"rounds"`
}
