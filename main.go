package main

import (
	"fmt"
)

func main() {
	trafficColors := map[string]string{
		"Green":  "Go",
		"Yellow": "Slow Down",
		"Red":    "Stop",
	}
	fmt.Println(trafficColors)

	printColorMsg(trafficColors)

	ChangeColor(trafficColors)
	if trafficColors["Green"] == "Dance" {
		fmt.Printf("Now you can %s\n", trafficColors["Green"])
	}

	fmt.Println(trafficColors)
}

func printColorMsg(trafficColors map[string]string) {
	fmt.Println()
	for color, msg := range trafficColors {
		fmt.Printf("If the Signal is %s then you should %s\n", color, msg)
	}
}

func ChangeColor(trafficColors map[string]string) {
	trafficColors["Green"] = "Dance"
}
