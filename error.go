package gotang

func CheckError(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}
