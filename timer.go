package app

import (
	"fmt"
	"time"
)

func NewTimer() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer1 completed.")
	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired.")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped.")
	}
	time.Sleep(2 * time.Second)
}
