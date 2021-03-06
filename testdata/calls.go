package main

type Thing struct {
	name string
}

func (t Thing) String() string {
	return t.name
}

func main() {
	thing := &Thing{"foo"}

	// function pointers
	runFunc(hello, 5) // must be indirect to avoid obvious inlining

	// deferred functions
	testDefer()

	// Take a bound method and use it as a function pointer.
	// This function pointer needs a context pointer.
	testBound(thing.String)

	// closures
	func() {
		println("thing inside closure:", thing.String())
	}()
	runFunc(func(i int) {
		println("inside fp closure:", thing.String(), i)
	}, 3)
}

func runFunc(f func(int), arg int) {
	f(arg)
}

func hello(n int) {
	println("hello from function pointer:", n)
}

func testDefer() {
	i := 1
	defer deferred("...run as defer", i)
	i++
	defer func() {
		println("...run closure deferred:", i)
	}()
	i++
	defer deferred("...run as defer", i)
	i++

	println("deferring...")
}

func deferred(msg string, i int) {
	println(msg, i)
}

func testBound(f func() string) {
	println("bound method:", f())
}
