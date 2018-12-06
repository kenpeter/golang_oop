// package lower case
package employee
// print
import "fmt"

// struct type
// upper case
type Employee struct {
	// upper case
	FirstName string
	LastName string
	TotalLeave int
	UseLeave int	
} 

// func, param, type 
// name
// no return
func (e Employee) LeaveLeft() {
	fmt.Printf("leave left: %d", e.TotalLeave - e.UseLeave)
}
