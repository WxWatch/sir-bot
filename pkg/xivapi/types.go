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
	ClassJobs []*ClassJob `json:"ClassJobs"`
}

type ClassJob struct {
	Class *Class `json:"Class"`
	Job   *Job   `json:"Job"`
	Level int    `json:"Level"`
	Name  string `json:"Name"`
}

type Class struct {
	Abbreviation string `json:"Abbreviation"`
}

type Job struct {
	Abbreviation string `json:"Abbreviation"`
}
