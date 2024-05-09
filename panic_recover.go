package utils

func TryCatchPanic(handler func()) {
	if err := recover(); err != nil {
		handler()
	}
}
