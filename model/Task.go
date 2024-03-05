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

var tasks = []Task{
	{ID:"1",Title: "Code the codes",Priority: strconv.Itoa(3),Status: "in Progress",Assignee: "john doe"},
	{ID:"1",Title: "Get the bag",Priority: strconv.Itoa(1),Status: "in Progress",Assignee: "lemonr lemonr"},
	{ID:"1",Title: "All the goodies",Priority: strconv.Itoa(2),Status: "To do",Assignee: "john heck"},
	{ID:"1",Title: "Buy new Machine",Priority: strconv.Itoa(4),Status: "in Progress",Assignee: "Work station"},
	{ID:"1",Title: "Gain it all",Priority: strconv.Itoa(5),Status: "in Progress",Assignee: "john doe"},
}