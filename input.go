

// fmt.scan -> read user input

package main // entry point 
 
import "fmt" //formst package print 


func greeting(){
	var name string
	var age int
	fmt.Print("Enter name")

	fmt.Scan(&name)

	fmt.Scan(&age)

}
