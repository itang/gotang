package gotang

// Deprecated.
// use MustNoError instead.
func CheckError(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}

func MustNoError(err error) {
	if err != nil {
		panic(err)
	}
}
