package Common

type Queue struct {
	data []interface{}
}

func (self *Queue) Empty() bool {
	return self.Size() == 0
}

func (self *Queue) Size() int {
	return len(self.data)
}

func (self *Queue) Top() interface{} {
	if self.Size() == 0 {
		panic("Queue empty")
	}

	return self.data[0]
}

func (self *Queue) Push(num interface{}) {
	self.data = append(self.data, num)
}

func (self *Queue) Pop() interface{} {
	if self.Size() == 0 {
		panic("Queue empty")
	}

	num := self.data[0]
	self.data = self.data[1:]
	return num
}
