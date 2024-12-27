package exception

func PanicErrorBadRequest(err error) {
	if err != nil {
		panic(NewErrorBadRequest(err.Error()))
	}
}
