package study

import (
	"fmt"
	"time"
)

func RoutineT1() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "alex", "emily", "patrick"}
	for _, person := range people {
		go whoIsSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func whoIsSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person
}
