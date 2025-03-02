package main

import "fmt"

var (

	Reset = "\033[0m" 
	Red = "\033[31m" 
	Green = "\033[32m" 
	Yellow = "\033[33m" 
	Blue = "\033[34m" 
	Magenta = "\033[35m" 
	Cyan = "\033[36m" 
	Gray = "\033[37m" 
	White = "\033[97m"
)


func main(){

	fmt.Print(Yellow + `
    ██╗   ██╗██╗     ██╗ █████╗ ██╗   ██╗
    ██║   ██║██║     ██║██╔══██╗╚██╗ ██╔╝
    ██║   ██║██║     ██║███████║ ╚████╔╝ 
    ╚██╗ ██╔╝██║██   ██║██╔══██║  ╚██╔╝  
      ████╔╝ ██║╚█████╔╝██║  ██║   ██║   
      ╚═══╝  ╚═╝ ╚════╝ ╚═╝  ╚═╝   ╚═╝   ` + Reset)
	fmt.Println("\n\tYour personal Agent")

	fmt.Print("\n\n")
	
	for {
		var input string
		fmt.Print(Red+"»"+Green+"» ")
		fmt.Scan(&input)
		fmt.Print("\n")

	}
}
