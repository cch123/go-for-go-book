package p

const k = 0

func f1() {
	type T [1]int
	_ = T{k: 0}
}

func f2() {
	type T struct{ k int }
	_ = T{k: 0}
}
