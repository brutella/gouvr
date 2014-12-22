package uvr

// Logs messages to the console
type consoleLogger struct {
}

func NewConsoleLogger() *consoleLogger {
    return &consoleLogger{}
}

func (l *consoleLogger) Log(message string) {
    print(message)
}

func (l *consoleLogger) Flush() {
}