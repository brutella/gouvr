package uvr

type consoleLogger struct {
}

// NewConsoleLogger returns a logger which logs messages to the console
func NewConsoleLogger() *consoleLogger {
	return &consoleLogger{}
}

func (l *consoleLogger) Log(message string) {
	print(message)
}

func (l *consoleLogger) Flush() {
}
