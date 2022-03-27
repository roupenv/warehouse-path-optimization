package main

type lineItem struct {
	sku      string
	quantity int
}

type order []lineItem

type orders []order
