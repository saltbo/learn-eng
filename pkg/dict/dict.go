package dict

type Word struct {
}

type Dict interface {
	Lookup(word string) (Word, error)
}

func Lookup(word string) (Word, error) {

}
