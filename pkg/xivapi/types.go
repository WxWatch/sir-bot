package xivapi

type CharacterSearchResult struct {
	Results []*Character `json:"Results"`
}

type CharacterDetailsResult struct {
	Character *Character `json:"Character"`
}

type Character struct {
	Avatar string `json:"Avatar"`
	ID     int    `json:"ID"`
	Name   string `json:"Name"`
	Server string `json:"Server"`

	//Extended Information
	DC             string       `json:"DC"`
	ClassJobs      []*ClassJob  `json:"ClassJobs"`
	ActiveClassJob *ClassJob    `json:"ActiveClassJob"`
	Race           *GenericInfo `json:"Race"`
	Title          *GenericInfo `json:"Title"`
	Town           *GenericInfo `json:"Town"`
}

type ClassJob struct {
	Class *Class `json:"Class"`
	Job   *Job   `json:"Job"`
	Level int    `json:"Level"`
	Name  string `json:"Name"`
}

type Class struct {
	Abbreviation     string           `json:"Abbreviation"`
	ClassJobCategory ClassJobCategory `json:"ClassJobCategory"`
}

type Job struct {
	Abbreviation     string           `json:"Abbreviation"`
	ClassJobCategory ClassJobCategory `json:"ClassJobCategory"`
}

type ClassJobCategory struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}

type GenericInfo struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
}
