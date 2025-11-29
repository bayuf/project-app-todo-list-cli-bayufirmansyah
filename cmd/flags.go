package cmd

type Flags struct {
	Add      string
	List     bool
	UpdateId int
	DelId    int
	Search   string

	// for update
	Priority string
	Status   string

	DoneID int
}
