package gomaps

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func Run() {
	fmt.Println("--Go Map----------")
	// A map maps keys to values.
	// map[keyType]valueType
	// we use map when we have to
	// -> fast lookup by key
	// -> no fixed order
	// -> unique keys

	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m)

	fmt.Println("--------Map Literals--------")
	// Map literals are like struct literals, but the keys are required.
	// If the top-level type is just a type name, you can omit it from the elements of the literal
	u := map[string]int{"Aashish": 21, "Rahul": 54}
	fmt.Println(u)

	fmt.Println("--------Mutating Maps--------")
	// Insert or update an element in map m: => m[key] = elem
	// Retrieve an element: => elem = m[key]
	// Delete an element: => delete(m,key)
	// Test that a key is present with a two-value assignment: elem, ok := m[key]
	// If key is in m, ok is true. If not, ok is false.
	// If key is not in the map, then elem is the zero value for the map's element type.

	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)

}
