package main

import (
	"fmt"
	"github.com/itz-1411/gocron/cron"
)

func main() {
	fmt.Println("hello world")
	
	schedular  := cron.NewScheduler()
	schedular.AddJob("*/10 * * * * *",  func(){
		fmt.Print("hello world\n")
	})
	schedular.Start()
	
	select {}
}
