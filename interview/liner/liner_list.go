package liner

type LinerListArr struct {
	Data []interface{}
	Last int
}

func (l *LinerListArr) Len() int {
	return l.Last + 1
}

type LinerListLink struct {
	lenghth int
}

func (l *LinerListLink) Len() int {
	return l.lenghth
}

type LinerList interface {
	Len() int
}

type List LinerListArr

func NewLinerList() *List {
	return &List{}
}

func FindKth(i int, l *List) interface{} {
	return nil
}

func Find(x interface{}, l *List) int {
	return 0
}

func Insert(x interface{}, i int, l *List) {

}

func Delete(i int, l *List) {

}

func Length(l *List) int {
	return l.Len()
}
