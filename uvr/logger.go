package uvr

import(
    "os"
    "fmt"
)

type Logger interface {
    Log(message string)
}

type consoleLogger struct {
}

func (l *consoleLogger) Log(message string) {
    print(message)
}

type fileLogger struct {
    filePath string
    buffer []byte
}

func NewFileLogger(filePath string) *fileLogger {
    l := &fileLogger{}
    l.filePath = filePath
    l.buffer = make([]byte, 0, 255)
    
    return l
}

func (l *fileLogger) reset() {
    l.buffer = make([]byte, 0, cap(l.buffer))
}

func (l *fileLogger) writeBufferToFile() {
    file, err := os.OpenFile(l.filePath, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
    if err != nil {
        fmt.Println("Could not open file at path", l.filePath, err)
    }
    
    defer file.Close()
    
    _, err = file.Write(l.buffer)
    file.Sync()
    if err != nil {
        fmt.Println("Could not write bytes to file.", err)
    }
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