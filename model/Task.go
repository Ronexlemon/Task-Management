package model

import (
	"strconv"
)

type Task struct{
	ID   string `json:"id"`
	Title  string `json:"title"`
	Priority string `json:"priority"`
	Status string `json:"status"`
	Assignee string `json:"assignee"`
}

var Tasks = []Task{
	{ID:"1",Title: "Code the codes",Priority: strconv.Itoa(3),Status: "in Progress",Assignee: "john doe"},
	{ID:"2",Title: "Get the bag",Priority: strconv.Itoa(1),Status: "in Progress",Assignee: "lemonr lemonr"},
	{ID:"3",Title: "All the goodies",Priority: strconv.Itoa(2),Status: "To do",Assignee: "john heck"},
	{ID:"4",Title: "Buy new Machine",Priority: strconv.Itoa(4),Status: "in Progress",Assignee: "Work station"},
	{ID:"5",Title: "Gain it all",Priority: strconv.Itoa(5),Status: "in Progress",Assignee: "john doe"},
}