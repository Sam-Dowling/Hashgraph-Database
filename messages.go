package main

/*
Code

0 : event sync -> event count
1 : events list

*/

type Message struct {
	Code int64
	Data interface{}
}

type EventCount struct {
	Count []int
}

type Events struct {
	Events []Event
}
