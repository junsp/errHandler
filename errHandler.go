package errHandler

func HandlerError(err error) {
	if err != nil {
		panic(err)
	}
}
