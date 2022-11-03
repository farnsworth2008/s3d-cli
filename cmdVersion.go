package main

import (
	"encoding/json"
	"fmt"
)

// Format our version
func runVersion(cmd *command, args []string) {
	m := make(map[string]string)
	m["version"] = "0.0.7"
	enc, err := json.Marshal(m)
	kill(err)
	fmt.Println(string(enc))
}

// Return our version command
func cmdVersion() *command {
	return versionCmd
}

// Our versionCmd command structure
var versionCmd = &command{
	Run:   runVersion,
	Short: "Display the program version",
	Use:   "version",
}
