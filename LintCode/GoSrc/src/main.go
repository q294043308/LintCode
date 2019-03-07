package main

func main() {
	Partition := []int{60, 3, 3, 3}
	var n uint32
	Map := make([]bool, 60)
	t := 0
	for n = 0; n < 100; n++ {
		partition := n % uint32(Partition[0])
		Map[partition] = true
		partition = n%uint32(Partition[0]-t) + uint32(0)
		Map[partition] = true
		partition = n % uint32(t)
		Map[partition] = true
	}
}
