package main

import (
	"fmt"
	"time"
)

func main() {

	//lab 1
	// theMine := [5]string{"rock", "ore", "ore", "rock"}
	// foundOre := finder(theMine)
	// mindedOre := miner(foundOre)
	// smelter(mindedOre)
	//

	//lab2 use go weight
	// theMine := [5]string{"rock", "ore", "ore", "rock"}
	// go finder1(theMine)
	// go finder2(theMine)
	// <-time.After(time.Second * 5)

	//lab3 channel
	// theMine1 := [5]string{"ore1", "ore2", "ore3", "ore4", "ore5"}
	// oreChan := make(chan string)

	// go func(mine [5]string) {
	// 	for _, m := range mine {
	// 		oreChan <- m //send
	// 	}
	// }(theMine1)

	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		foundOre := <-oreChan // receive
	// 		fmt.Println("Miner :Received " + foundOre + " from finder")
	// 	}

	// }()
	// <-time.After(time.Second * 5)

	//lab4 channel buffer
	// bufferedChan := make(chan string, 3)

	// go func() {
	// 	bufferedChan <- "first"
	// 	fmt.Println("Sent first")
	// 	bufferedChan <- "2nd"
	// 	fmt.Println("sent 2nd ")
	// 	bufferedChan <- "3th"
	// 	fmt.Println("sent 3rd")

	// }()

	// <-time.After(time.Second * 1)

	// go func() {

	// 	firstRead := <-bufferedChan
	// 	fmt.Println("Receiving...")
	// 	fmt.Println(firstRead)

	// 	secondRead := <-bufferedChan
	// 	fmt.Println(secondRead)

	// 	thirdRead := <-bufferedChan
	// 	fmt.Println(thirdRead)

	// }()
	// <-time.After(time.Second * 1)

	//lab5  putting  altogether
	theMine := [5]string{"rock", "ore", "ore", "rock"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	//finder
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item
			}

		}
	}(theMine)

	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel
			fmt.Println("From finder:  " + foundOre)
			minedOreChan <- "minededOre"
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			mindedOre := <-minedOreChan
			fmt.Println("From Miner:  " + mindedOre)
			fmt.Println("From Smelter:   Ore is Smelted~")
		}
	}()

	<-time.After(time.Second * 5)
}

func finder(mine [5]string) [5]string {

	fmt.Printf("From finder: %s \n", mine)
	return mine
}
func miner(mine [5]string) [5]string {
	fmt.Printf("From miner:%s  \n", mine)
	return mine
}

func smelter(mine [5]string) {
	fmt.Printf("From miner: %s  \n", mine)
}
func finder1(mine [5]string) {

	for _, m := range mine {
		if m == "ore" {
			fmt.Printf("finder1 fine ore! \n")
		}

	}
}
func finder2(mine [5]string) {

	for _, m := range mine {
		if m == "ore" {
			fmt.Printf("finder2 fine ore! \n")
		}

	}
}
