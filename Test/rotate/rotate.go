package rotate

func main() {

}

func Rotate(n []int) []int {
	if len(n) < 2 {
		return n
	}
	l := n[len(n)-1]
	r := []int{l}
	return append(r, n[0:len(n)-1]...)
}
