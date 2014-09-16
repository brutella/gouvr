package uvr

// Logs messages to the console
type consoleLogger struct {
}

func (l *consoleLogger) Log(message string) {
    print(message)
}