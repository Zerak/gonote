package varStore

var global *int

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	x := new(int)
	*x = 1
}