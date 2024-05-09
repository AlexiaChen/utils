package utils

func TryCatchPanic(handler func(interface{})) {
	if err := recover(); err != nil {
		handler(err)
	}
}
