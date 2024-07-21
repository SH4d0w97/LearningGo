package mypkg // Package declaration

//------------ BASIC METHOD DECLERATION ----------

// triangle represents a triangle shape with a given size.
type triangle struct {
    size int
}

// square represents a square shape with a given size.
type square struct {
    size int
}

// perimeter calculates the perimeter of a triangle.
func (t triangle) perimeter() int {
    return t.size * 3
}

// perimeter calculates the perimeter of a square.
func (s square) perimeter() int {
    return s.size * 4
}

// func main() {
//     t := triangle{3}  // Initialize a triangle with size 3
//     s := square{4}   // Initialize a square with size 4

//     // Calculate and print the perimeters
//     fmt.Println("Perimeter (triangle):", t.perimeter())
//     fmt.Println("Perimeter (square):", s.perimeter())
// }


//---------- Pointers in methods ----------
// doubleSize doubles the size of the triangle.
// The method takes a pointer receiver, allowing it to modify the original triangle.
func (t *triangle) doubleSize() {
    t.size *= 2 
}

// func main() {
//     t := triangle{3}      // Initialize a triangle with size 3
//     t.doubleSize()        // Double the size of the triangle
//     fmt.Println("Size:", t.size)          // Print the new size
//     fmt.Println("Perimeter:", t.perimeter()) // Print the perimeter (assuming 'perimeter' method is defined elsewhere)
// }

//---------- Declare methods for other types ----------

//	 import (
//	     "fmt"
//	     "strings"
//	 )

// upperstring is a custom type based on string.
//	type upperstring string

//	 Upper converts the upperstring to uppercase.
//	 func (s upperstring) Upper() string {
//	     return strings.ToUpper(string(s))
//	 }

// func main() {
//     s := upperstring("Learning Go!") // Create an upperstring value
//     fmt.Println(s)                  // Print the original string
//     fmt.Println(s.Upper())          // Print the uppercase version
// }

//---------- Embed methods ----------
// coloredTriangle is a struct that embeds the triangle struct and adds a color field.
type coloredTriangle struct {
    triangle // Embed the triangle struct
    color string
}

// func main() {
//     // Create a coloredTriangle instance with size 3 and color blue.
//     t := coloredTriangle{triangle{3}, "blue"} 
    
//     // Access the size field from the embedded triangle struct.
//     fmt.Println("Size:", t.size)    
    
//     // Call the perimeter method from the embedded triangle struct.
//     fmt.Println("Perimeter", t.perimeter()) 
// }

//--------- Overload methods ---------
// perimeter calculates the perimeter of a coloredTriangle (overloaded method).
func (t coloredTriangle) perimeter() int {
    return t.size * 3 * 2 // Calculation for colored triangle perimeter. 
}

// func main() {
//     // Create a coloredTriangle instance with size 3 and color blue.
//     t := coloredTriangle{triangle{3}, "blue"} 
    
//     // Access the size field from the embedded triangle struct.
//     fmt.Println("Size:", t.size)    
    
//     // Call the perimeter method of coloredTriangle (overloaded method).
//     fmt.Println("Perimeter", t.perimeter()) 
// }

//---------- Encapsulation in methods ----------
//  package geometry // Package declaration

// Triangle represents a triangle shape with a given size.
type Triangle struct {
    size int
}

// doubleSize doubles the size of the triangle.
// The method takes a pointer receiver, allowing it to modify the original triangle.
func (t *Triangle) doubleSize() {
    t.size *= 2
}

// SetSize sets the size of the triangle.
func (t *Triangle) SetSize(size int) {
    t.size = size
}

// Perimeter calculates and returns the perimeter of the triangle.
func (t *Triangle) Perimeter() int {
    t.doubleSize() // Doubles the size before calculating the perimeter.
    return t.size * 3
}

// func main() {
//     t := geometry.Triangle{} // Initialize a triangle.
//     t.SetSize(3)             // Set the size of the triangle to 3.
//     fmt.Println("Perimeter", t.Perimeter()) // Calculate and print the perimeter.
// }
