package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

var min int

func main() {

	// Make sure we enable Notifitcations when the program exits
	defer disableDND()
	flag.IntVar(&min, "t", 15, "Time in minutes for 'Do not disturb' mode")
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
			fmt.Printf("Turning off 'Do not disturb' mode in %d minutes\n", min)
		case <-sigChan:
			fmt.Println("Got SIGTERM")
			return
		}
	}

}

func enableDND() {
	fmt.Printf("Turning on 'Do not disturb' mode for %d minutes\n", min)

	cmd := exec.Command("sh", "-c", "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturb -boolean true")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not set doNotDisturb to true: %v", err)
		os.Exit(1)
	}

	cmd = exec.Command("sh", "-c", "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturbDate -date \"`date -u +\"%Y-%m-%d %H:%M:%S +000\"`\"")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not write time to ~/Library/Preferences/ByHost/com.apple.notificationcenterui: %v", err)
		os.Exit(1)
	}

	cmd = exec.Command("sh", "-c", "killall NotificationCenter")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not reset NotificationCenter: %v", err)
		os.Exit(1)
	}
}

func disableDND() {
	cmd := exec.Command("sh", "-c", "defaults -currentHost write ~/Library/Preferences/ByHost/com.apple.notificationcenterui doNotDisturb -boolean false")
	_, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not set doNotDisturb to false: %v", err)
		os.Exit(1)
	}

	cmd = exec.Command("sh", "-c", "killall NotificationCenter")
	_, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not reset NotificationCenter: %v", err)
		os.Exit(1)
	}

	fmt.Println("Turning off 'Do not disturb' mode")
}
