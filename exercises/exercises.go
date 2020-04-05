package main
import "fmt"

func main()  {
	slice()
	dMap()
	fPerson()
	usingVehicles()
	usingGators()
	usingSlices()
}

func slice() {
	mySlice := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(mySlice)

	for i := range mySlice{
		fmt.Println(i)
	}

	for i, v := range mySlice{
		fmt.Println(i, v)
	}
}

func dMap() {
	myMap := map[string]int{"Golang": 10, "Javascript": 9}
	
	fmt.Println(myMap)

	for k := range myMap {
		fmt.Println(k)
	}

	for k, v := range myMap {
		fmt.Println(k, v)
	}
}

type person struct {
	fName string
	lName string
	favFood []string
}

func (p person) walk() string{
	return fmt.Sprintln(p.fName, "is walking")
}

func fPerson() {
	p1 := person{
		"Jonh",
		"Snow",
		[]string{"Sushi", "Hamburguer"},
	}
	fmt.Println(p1)
	fmt.Println(p1.fName)
	fmt.Println(p1.favFood)

	for _, v := range p1.favFood{
		fmt.Println(v)
	}

	s:= p1.walk()
	
	fmt.Println(s)
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("I am a truck with", t.doors, "doors and I am durable.")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("I am a sedan with", s.doors, "doors and I am luxurious.")
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func usingVehicles() {
	t1 := truck {
		vehicle {
			2,
			"red",
		},
		true,
	}

	s1 := sedan {
		vehicle {
			4,
			"blue",
		},
		false,
	}

	fmt.Println(t1)
	fmt.Println(t1.color)
	report(t1)

	fmt.Println(s1)
	fmt.Println(s1.color)
	report(s1)

}

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I'm a gator")

}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(sc swampCreature) {
	sc.greeting()
}

func usingGators() {
	var g1 gator

	g1 = 3

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)

	var x int
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)
	fmt.Println(x)

	g1.greeting()
	bayou(g1)

	var f1 flamingo
	bayou(f1)
	
}

func usingSlices() {
	s := "I'm sorry dave I can't do that"
	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	fmt.Println(string([]byte(s)[:14]))
	fmt.Println(string([]byte(s)[10:22]))
	fmt.Println(string([]byte(s)[17:]))

	for _, v := range s{
		fmt.Println(string(v))
	}
}