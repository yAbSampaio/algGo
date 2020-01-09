package main

import (
	"fmt"
	"os"
	"strconv"
)
func changeUnity (iniUnity string) string {
	if iniUnity == "celsius" {
		return "fahrenheit"
	} else if iniUnity == "fahrenheit"{
		return "celsius"
	} else if iniUnity == "kilometers" {
		return"miles"
	
	} else if iniUnity == "miles" {
		return "kilometers"
	} else {
		fmt.Printf("%s not a known unit!\n",
			iniUnity)
		os.Exit(1)
		return "error"
	}
}

func changeValue (value float64, iniUnity string) float64 {
	if iniUnity == "celsius" {
		return value*1.8 + 32
	} else if iniUnity == "fahrenheit"{
		return (value-32)/1.8
	} else if iniUnity == "kilometers" {
		return value/1.60934
	
	} else  {
		return value*1.60934
	}
}

func main(){

	var endUnity string
	var endValue float64

	if len(os.Args) < 3 {
		fmt.Println("Use: converter <values> <unity>")
		os.Exit(1)
	}
	iniUnity := os.Args[len(os.Args)-1]
	iniValues := os.Args[1: len(os.Args)-1]

	endUnity = changeUnity(iniUnity)
	
	for i, v := range iniValues {
		iniValue, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf(
				"The value %s in position %d not is a valid number\n",
				v,i)
			os.Exit(1)
		}

		endValue = changeValue(iniValue,iniUnity)
		fmt.Printf("%.2f %s = %.2f %s\n",
			iniValue, iniUnity,
			endValue, endUnity)
	}
}