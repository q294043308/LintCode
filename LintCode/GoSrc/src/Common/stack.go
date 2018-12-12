package Common

type Stack struct {
	data []int
}

func (self *Stack) Empty() bool {
	return self.Size() == 0
}

func (self *Stack) Size() int {
	return len(self.data)
}

func (self *Stack) Top() int {
	if self.Size() == 0 {
		panic("Stack empty")
	}

	return self.data[len(self.data)-1]
}

func (self *Stack) Push(num int) {
	self.data = append(self.data, num)
}

func (self *Stack) Pop() int {
	if self.Size() == 0 {
		panic("Stack empty")
	}

	num := self.data[len(self.data)-1]
	self.data = self.data[:len(self.data)-1]
	return num
}
