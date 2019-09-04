package main

var (
	mapStr = make(map[string]int)
	mapInt = make(map[int]int)
)

func noop(a interface{}) interface{} {
	return a
}

func writeStr(n int) {
	j := 0
	for j < n {
		i := 0
		for i < len(strkeys) {
			mapStr[strkeys[i]] = i
			i++

		}
		j++
	}
}

func readStr(n int) {
	i := 0
	for i < n {
		j := 0
		for j <= len(strkeys) {
			a := mapStr[strkeys[i]]
			noop(a)
			j++
		}
		i++
	}
}

func writeInt(n int) {
	j := 0
	for j < n {
		i := 0
		for i < len(intkeys) {
			mapInt[intkeys[i]] = i
			i++
		}
		j++
	}
}

func readInt(n int) {
	i := 0
	for i < n {
		j := 0
		for j <= len(intkeys) {
			a := mapInt[intkeys[i]]
			noop(a)
			j++
		}
		i++
	}
}
