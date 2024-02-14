package reports

type Report struct {
	Id      int
	Name    string
	Rows    int
	Columns int
	Query   string
}

// add into base queries that struct
