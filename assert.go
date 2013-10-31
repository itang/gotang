package gotang

// Deprecated.
// use `MustNoError` instead.
func CheckError(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}

func AssertNoError(err error) {
	if err != nil {
		panic(err)
	}
}

func Assert(assertion bool, message string) {
	if !assertion {
		detail := message
		if detail != "" {
			detail = ": " + detail
		}
		panic("assertion failed" + detail)
	}
}
