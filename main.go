package main

import (
	"fmt"
	goroutine "go-practice/advanced/goroutine"
	interface_example "go-practice/advanced/interface"
	account "go-practice/basics/account"
	calculator "go-practice/basics/calculator"
	closure "go-practice/basics/closure"
	panic "go-practice/basics/panic"
	student "go-practice/basics/student"
)

func main() {
	// Calculator
	calc := calculator.NewCalculator()
	sum := calc.Add(10, 5)
	fmt.Printf("10 + 5 = %.2f\n", sum)
	fmt.Println("Calculator History:", calc.GetHistory())

	diff := calc.Subtract(10, 5)
	fmt.Printf("10 - 5 = %.2f\n", diff)
	fmt.Println("Calculator History:", calc.GetHistory())

	product := calc.Multiply(10, 5)
	fmt.Printf("10 * 5 = %.2f\n", product)
	fmt.Println("Calculator History:", calc.GetHistory())

	quotient, err := calc.Divide(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("10 / 5 = %.2f\n", quotient)
	}
	fmt.Println("Calculator History:", calc.GetHistory()) // automatic dereferencing of pointer: (*calc).GetHistory()

	quotient, err = calc.Divide(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("10 / 0 = %.2f\n", quotient)
	}

	mod := calc.Modulus(10, 3)
	fmt.Printf("10 %% 3 = %.2f\n", mod)
	fmt.Println("Calculator History:", calc.GetHistory())

	pow := calc.Power(2, 3)
	fmt.Printf("2 ^ 3 = %.2f\n", pow)
	fmt.Println("Calculator History:", calc.GetHistory())

	numbers := []float64{1, 2, 3, 4, 5}
	total := calc.Sum(numbers...) // expand this slice into individual arguments
	fmt.Printf("Sum of %v = %.2f\n", numbers, total)
	fmt.Println("Calculator History:", calc.GetHistory())

	average := calc.Average(numbers...)
	fmt.Printf("Average of %v = %.2f\n", numbers, average)
	fmt.Println("Calculator History:", calc.GetHistory())

	calc.ClearHistory()
	fmt.Println("Calculator History:", calc.GetHistory())

	// Student Management
	manager := student.NewManager()
	student1 := student.Student{Name: "Alice", Age: 20, Grade: "A", ID: 1}
	student2 := student.Student{Name: "Bob", Age: 22, Grade: "B", ID: 2}

	manager.AddStudent(student1)
	manager.AddStudent(student2)

	s, err := manager.GetStudentByID(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Retrieved Student: %+v\n", s)
	}

	err = manager.UpdateStudentGrade(1, "A+")
	if err != nil {
		fmt.Println(err)
	} else {
		s, _ = manager.GetStudentByID(1)
		fmt.Printf("Updated Student Grade to A+: %+v\n", s)
	}

	students := manager.ListStudents()
	fmt.Printf("All Students: %+v\n", students)

	err = manager.RemoveStudentByID(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Removed Student with ID 2")
	}

	students = manager.ListStudents()
	fmt.Printf("All Students: %+v\n", students)

	// Account Management
	accountManager := account.NewManager()
	acc1 := accountManager.OpenAccount("John Doe")
	acc2 := accountManager.OpenAccount("Jane Smith")

	err = accountManager.Deposit(acc1.ID, 1000)
	if err != nil {
		fmt.Println(err)
	} else {
		balance, _ := accountManager.GetBalance(acc1.ID)
		fmt.Printf("Deposited $1000 to %s's account. New Balance: $%.2f\n", acc1.Name, balance)
	}

	err = accountManager.WithDraw(acc1.ID, 500)
	if err != nil {
		fmt.Println(err)
	} else {
		balance, _ := accountManager.GetBalance(acc1.ID)
		fmt.Printf("Withdrew $500 from %s's account. New Balance: $%.2f\n", acc1.Name, balance)
	}

	balance, err := accountManager.GetBalance(3) // non-existing account
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s's account balance: $%.2f\n", acc2.Name, balance)
	}

	err2 := accountManager.WithDraw(acc2.ID, 100) // insufficient funds
	if err2 != nil {
		fmt.Println(err2)
	} else {
		balance, _ := accountManager.GetBalance(acc2.ID)
		fmt.Printf("Withdrew $100 from %s's account. New Balance: $%.2f\n", acc2.Name, balance)
	}

	// Closure Examples 1 - Counter
	counter := closure.NewCounter()
	fmt.Println("Counter:", counter()) // 1
	fmt.Println("Counter:", counter()) // 2
	fmt.Println("Counter:", counter()) // 3

	// Closure Examples 2 - UI Handlers
	btnA := closure.Button{Label: "Button A"}
	btnB := closure.Button{Label: "Button B"}

	// Set up handlers using struct and method
	btnAHandler := closure.NewHandler("HandlerA")

	// Set up handlers using closures
	btnBHandler := closure.MakeHandler("HandlerB")

	// Simulate button clicks
	btnA.OnClick(btnAHandler.Handle) // method value as function
	btnB.OnClick(btnBHandler)
	btnA.OnClick(btnAHandler.Handle)
	btnB.OnClick(btnBHandler)

	// Closure Examples 3 - Price Strategies
	originalPrice := 100.0

	discount := closure.DiscountStrategy(0.25) // 25% off
	coupon := closure.WithCoupon(10.0)         // $10 off
	final := closure.Combine(discount, coupon)
	fmt.Printf("Original Price: $%.2f, Final Price after strategies: $%.2f\n", originalPrice, final(originalPrice))

	// Handle panic example
	err3 := panic.ProcessFile("README.md") // adjust the path as needed
	if err3 != nil {
		fmt.Println(err3)
	}

	// Interface Basis
	shapes := []interface_example.Shape{
		interface_example.Rectangle{Width: 10, Height: 5},
		interface_example.Circle{Radius: 7},
	}

	for _, shape := range shapes {
		fmt.Printf("Shape: %T, Area: %.2f, Perimeter: %.2f\n", shape, shape.Area(), shape.Perimeter())
	}

	// Empty Interface - Example 1
	var i interface_example.EmptyInterface

	i = 42
	fmt.Println(interface_example.PrintValue(i))

	i = "Hello, Go!"
	fmt.Println(interface_example.PrintValue(i))

	i = 3.14
	fmt.Println(interface_example.PrintValue(i))

	i = struct{ Name string }{Name: "Gopher"}
	fmt.Println(interface_example.PrintValue(i))

	// Empty Interface - Type Assertion and Type Switch
	interface_example.TypeAssertion("A string value")
	interface_example.TypeAssertion(100)

	i = []int{1, 2, 3}
	interface_example.TypeSwitch(i)

	i = "A string value"
	interface_example.TypeSwitch(i)

	i = 100
	interface_example.TypeSwitch(i)

	i = 3.14
	interface_example.TypeSwitch(i)

	// Empty Interface - Example 3: JSON parsing with unknown structure
	// data is of unknown structure, could be anything
	jsonData := `{
		"code": 0,
		"message": "ok",
		"data": {
			"id": 101,
			"name": "Alice",
			"age": 23
		}
	}`

	// Parse incoming JSON with unknown structure using empty interfaces
	user, err := interface_example.ParseJSON([]byte(jsonData))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Parsed User: %+v\n", user)

	// Outgoing Publishing with unknown payload using empty interface
	interface_example.Publish("UserCreated", user)

	// Interface Composition
	file := &interface_example.File{Name: "example.txt"}
	data := []byte("Hello, Interface Composition!")

	interface_example.Process(data, file)

	// Interface Polymorphism
	mysqlDB := &interface_example.MySQL{Connection: "user:pass@tcp(localhost:3306)/dbname"}
	result, err := interface_example.ExecuteQuery(mysqlDB, "SELECT * FROM users")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("MySQL Query Result: %v\n", result)
	}

	postgresDB := &interface_example.PostgreSQL{Connection: "postgres://user:pass@localhost:5432/dbname"}
	result, err = interface_example.ExecuteQuery(postgresDB, "SELECT * FROM products")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("PostgreSQL Query Result: %v\n", result)
	}

	// Goroutine Basis - WaitGroup example
	goroutine.WaitGroupExampe()

	// Goroutine - Channel examples
	goroutine.UnbufferedChannelExample()
	goroutine.BufferedChannelExample()
	goroutine.BufferedChannelExample2()

	// Goroutine - Select example
	goroutine.SelectBasis()
	goroutine.SelectNonBlockingWithDefault()
	goroutine.SelectWithTimeout()
	goroutine.FanInChannels()
}
