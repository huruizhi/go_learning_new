package main

import "go_learning_new/day6/homework/log"

func main() {
	l := log.NewLogging(log.LogFile)
	l.SetLogFile("logfile")
	msg := `adafasddafadfbcdefg abcdefg
kljdfiae ewrwe
ewrwew erwewwr`
	l.Trace(msg)
}
