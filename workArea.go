package main

type workArea struct {
	numberItems int
	location    junctionNode
}

func (wa *workArea) incrementNumItems(newItems int) {
	wa.numberItems += newItems
}
