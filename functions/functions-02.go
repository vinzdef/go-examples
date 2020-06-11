func exec(action string) (result int, err error) {
	result = doSomething(action)
	if result == nil {
		err = errors.New("Error!11	!1")
	}

	return
}