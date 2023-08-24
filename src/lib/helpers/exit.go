package helpers

func FinishOnError(err error) {
	if err == nil {
		return
	}

	panic(err)
}
