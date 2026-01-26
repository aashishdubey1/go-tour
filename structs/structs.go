package structs

import "fmt"

func Run() {
	fmt.Println("-------GO STRUCT-------")

	// WHAT IS STRUCT ?
	// => A struct is custom data type that lets you group multiple values(fields) types in one unit.
	// this is like real world object , ex a user has - name , age , email . so instead of seperate variables we can create struct

	type User struct {
		Name  string
		Age   int
		email string
	}

	// this is the syntax for creating a struct . we use type because we are creating custom type followed by type name and struct keyword
	// generally we write all fields of struct in capital letter that way we can export them means we can use them in different package
	// we can also use small letter in fields name then we can only use them in same package

	fmt.Println("-------STRUCT CREATED-------")
	// this is how we initialise struct , we can say that email fields here is private
	u := User{Name: "Aashish", Age: 54, email: "aashish@gmail.com"}
	fmt.Println(u)
	fmt.Println(u.Name)

	// Zero value struct
	// There's concept of Zero value struct , if we don't pass anything the fields got assigned to defalut values
	/*
		string → ""
		int → 0
		bool → false
		pointer → nil
	*/
	// like this
	fmt.Println("------- CAN UPDATE STRUCT-------")
	u.Name = "Aashish dubey"
	u.email = "newEmail.com"
	fmt.Println(u)
	fmt.Println(u.Name)

	fmt.Println("-------STRUCT POINTER-------")
	// Struct fields can be accessed through a struct pointer.
	// if we pass struct with value the copy will get update , if we want to update original value we pass pointer to struct.
	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{X: 45, Y: 90}
	fmt.Println(v)
	pV := &v
	pV.X = 90
	pV.Y = 45
	fmt.Println(v)
	// To access the field X of a struct when we have the struct pointer p we could write (*p).X
	// However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.

	// OPTIONS 2
	// We can also create Struct pointer with "new" keyword
	pv := new(Vertex)
	fmt.Println(pv)
	// new(User) allocates a zero-value Vertex and returns its pointer.

	fmt.Println("-------NESTED STRUCT-------")
	type Add struct {
		City string
		Pin  int
	}
	type Person struct {
		Name    string
		Address Add
	}
	p := Person{Name: "Aashish", Address: Add{City: "Ssm", Pin: 802215}}
	fmt.Println(p)

	fmt.Println("-------ANONYMOUS STRUCT-------")
	x := struct {
		X int
		Y int
	}{
		X: 45,
		Y: 54,
	}
	fmt.Println(x)

}

// Struct rules
// ✅ Struct is a collection of fields grouped into one type
// ✅ Define using type Name struct { ... }
// ✅ Field names must have a type (Name string)
// ✅ Struct values are copied when assigned or passed (by default)
// ✅ Use *StructType to pass reference / address
// ✅ Use &value to get pointer of a struct value
// ✅ Use new(Type) to allocate zero value and return pointer
// ✅ u.Field works even if u is a pointer (*User) (auto deref)
// ✅ Struct fields have zero values if not initialized
// ✅ You can create structs using:
// named fields ✅ recommended
// in-order fields ⚠️ not recommended
// ✅ Structs can be nested (struct inside struct)
// ✅ Anonymous structs are allowed
// ✅ Struct methods exist (receiver functions)
// ✅ Pointer receiver methods can modify original
// ✅ Struct tags are used for JSON, DB, validation, etc.
// ✅ Structs can be compared with == only if all fields are comparable
// ❌ Struct can’t contain itself directly (infinite size)
// ✅ but can contain pointer to itself:
// type Node struct { Next *Node }
// ✅ Export rule:
// Name (capital) = exported/public
// name (small) = unexported/private
