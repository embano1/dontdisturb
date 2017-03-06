package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var min int

func main() {

	defer disableDND()
	flag.IntVar(&min, "t", 15, "Time in minutes for \"Do not disturb\" mode")
	flag.Parse()

	timerChan := time.NewTimer(time.Minute * time.Duration(min)).C
	tickerChan := time.NewTicker(time.Minute * 1).C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go enableDND()

	for {
		select {
		case <-timerChan:
			return
		case <-tickerChan:
			min--
			fmt.Printf("Turning off \"Do not disturb\" mode in %d minutes \n", min)
		case <-sigChan:
			fmt.Println("Got SIGTERM")
			return
		}
	}

}

func enableDND() {
	fmt.Printf("Turning on \"Do not disturb\" mode for %d minutes \n", min)
	_, _ = exec.Command("sh", "-c", "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturb -boolean true").Output()
	_, err := exec.Command("sh", "-c", "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturbDate -date \"`date -u +\"%Y-%m-%d %H:%M:%S +000\"`\"").Output()
	if err != nil {
		log.Fatal("2", err)
	}
	_, _ = exec.Command("sh", "-c", "killall NotificationCenter").Output()

}

func disableDND() {
	_, _ = exec.Command("sh", "-c", "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturb -boolean false").Output()
	_, _ = exec.Command("sh", "-c", "killall NotificationCenter").Output()
	fmt.Println("Turning off \"Do not disturb\" mode")
}
