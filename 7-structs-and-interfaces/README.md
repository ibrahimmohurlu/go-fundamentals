# Structs and Interfaces

## Structs

We can define new types with structs in Go. The type keyword introduces a new type. It's followed by the name of the type (Circle), the keyword struct to indicate that we are defining a struct type and a list of fields inside of curly braces. Each field has a name and a type. Like with functions we can collapse fields that have the same type:

```go
type Circle struct {
  x float64 // or 'x, y, r float64'
  y float64
  r float64
}
```

There are various ways to create instance of the new type

```go
var c Circle // sets values to defaults (0.0 for floats)
c := new(Circle) // returns pointer
c := Circle{x: 0, y: 0, r: 5} // creates instance and sets values
c := Circle{0, 0, 5} // fields will get values order they were defined

fmt.Println(c.x, c.y, c.r) // we can access fields by . notation
c.x = 10
c.y = 5
```

We can also attach methods(functions with special syntax) to the new type

```go
func (c *Circle) area() float64 { // (c *Circle) 'receiver' which means
  return math.Pi * c.r*c.r        // this function belongs to the Circle type
}

fmt.Println(c.area()) // we can also call this method with . notation
```

We can also embed other custom types too

```go
type Person struct {
  Name string
}
func (p *Person) Talk() {
  fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
  Person Person
  Model string
}

var me = Android{person: Person{Name: "John"}}
me.person.Talk() // we can call Talk method with named field
```

We can also embed other types without giving a name

```go
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model  string
}

me2 := Android{Person: Person{Name: "Doe"}}
me2.Talk() // this time we can call Talk method directly
me2.Person.Talk() // or its type name
```

## Interfaces

Like a struct an interface is created using the type keyword, followed by a name and the keyword interface. But instead of defining fields, we define a “method set”. A method set is a list of methods that a type must have in order to “implement” the interface.

```go
type Shape interface {
  area() float64
}
```

We can use interface types as arguments to functions

```go
func totalArea(shapes ...Shape) float64 { // spread syntax
  var area float64
  for _, s := range shapes {
    area += s.area()
  }
  return area
}
fmt.Println(totalArea(&c, &r))
```

Interfaces can also be used as fields

```go
type MultiShape struct {
  shapes []Shape
}
```

We can even turn <code>MultiShape</code> itself into a <code>Shape</code> by giving it an area method

```go
func (m *MultiShape) area() float64 {
  var area float64
  for _, s := range m.shapes {
    area += s.area()
  }
  return area
}
```

Now a <code>MultiShape</code> can contain <code>Circles</code>, <code>Rectangles</code> or even other <code>MultiShapes</code>.
