package main

import "fmt"

var linkListName = [6]string{"Swapnils List", "Rohit List", "Anand List", "Sunil List", "Onkar List", "Omkar List"}

var linkListIDs = [6]int{001, 002, 003, 004, 005, 006}

var linkListCount = [6][2]int{
	{001, 5},
	{002, 3},
	{003, 7},
	{004, 7},
	{005, 23},
	{006, 9},
}

func main() {
	fmt.Println("Start to code")
	printListArray()
	theTitleList()
}

func printListArray() {
	for i := 0; i < len(linkListName); i++ {
		fmt.Println("List Name >> ", linkListName[i])
	}
	fmt.Println("Multi D array > ", linkListCount)

}

func theTitleList() {
	titleOne := []string{"this is title one ", "save your keychain"}
	titleTwo := []string{"Common wins and odd mistakes", "collective intilligence", "world and it's rules"}
	titleThree := []string{"for we can", "design system"}
	titleFour := []string{"search bette", "true song"}
	titleFive := []string{"books to read"}
	titleSix := []string{"Good search", "tools to code"}

	fmt.Println("content of title 1-2 >> ", titleOne, titleTwo)
	fmt.Println("content of title 3-4 >> ", titleThree, titleFour)
	fmt.Println("content of title 5-6 >> ", titleFive, titleSix)
}
