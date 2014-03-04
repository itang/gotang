package gotang

// Deprecated.
// use `AssertNoError` instead.
func CheckError(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}

func AssertNoError(errs ...error) {
	for _, err := range errs {
		if err != nil {
			panic(err)
		}
	}
}

func Assert(assertion bool, message string) {
	if !assertion {
		err := "assertion failed"
		if message != "" {
			err += ": "+message
		}
		panic(err)
	}
}
