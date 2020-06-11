package main

import "fmt"

type TestError struct{}

func (t TestError) Error() string {
	return fmt.Sprint("Error!1! !!1")
}

func doError(shouldError bool) *TestError {
	if shouldError {
		return &TestError{}
	} else {
		return nil
	}
}

func Test(shouldError bool) error {
	return doError(shouldError)
}

func main() {
	if Test(true) != nil {
		fmt.Println("Error")
	}

	if Test(false) != nil {
		fmt.Println("Error")
	}
}
