package uvr

import(
    "os"
    "fmt"
)

type fileLogger struct {
    filePath string
    buffer []byte
}

// NewConsoleLogger returns a logger which logs messages to a file
func NewFileLogger(filePath string, buffer_size int) *fileLogger {
    l := &fileLogger{}
    l.filePath = filePath
    l.buffer = make([]byte, 0, buffer_size)
    
    return l
}

func (l *fileLogger) Log(message string) {
    capacity := cap(l.buffer)
    buffer := append(l.buffer, []byte(message)...)
    
    if capacity != cap(buffer) {
        l.writeBufferToFile()
        l.reset()
    } else {
        l.buffer = buffer
    }
}

// Flushes the stream to file
func (l *fileLogger) Flush() {
    l.writeBufferToFile()
    l.reset()
}

// Resets the memory buffer
func (l *fileLogger) reset() {
    l.buffer = make([]byte, 0, cap(l.buffer))
}

// Writes the memory buffer to file
func (l *fileLogger) writeBufferToFile() {
    file, err := os.OpenFile(l.filePath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Could not open file at path", l.filePath, err)
    }
    
    defer file.Close()
    
    fmt.Println("Save...")
    _, err = file.Write(l.buffer)
    file.Sync()
    if err != nil {
        fmt.Println("Could not write bytes to file.", err)
    }
}