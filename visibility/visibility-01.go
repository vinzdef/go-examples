package visibility

// Visible to clients of `visibility` package
var VisibleVar int

func VisibleMethod(flag bool) bool {
	return !flag
}

// Local, only visible to package members
var anotherVar, _stillLocal = 9000, 42
