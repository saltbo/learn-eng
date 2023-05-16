package plan

type Plan interface {
	Run() error
}

type Word struct {
	Name          string
	Phonetic      string
	Meaning       string
	WordsGraph    string
	Roots         []string
	RootGraph     string
	MemorizeVideo string
	Examples      []Example
}

type Example struct {
}
