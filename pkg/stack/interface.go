package stack

type Emptyer interface {
	Empty() bool
}

type Lenther interface {
	Len() int
}

type Interface interface {
	Emptyer
	Lenther
	Push(interface{})
	Pop() (interface{}, error)
}

type StackInt interface {
	Emptyer
	Lenther
	Push(int)
	Pop() (int, error)
}
