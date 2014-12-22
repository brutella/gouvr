package uvr

import(
    "os"
    "io"
    "bufio"
    "bytes"
)

type replayer struct {
    consumer BitConsumer
}
// NewReplayer returns a replayer which simulates a bit steam from a log file.
// The bit log strings must be created by using the result of calling `LogString(bit)`.
func NewReplayer(consumer BitConsumer) *replayer {
    replayer := &replayer{consumer:consumer}
    
    return replayer
}

func (r *replayer) ReadLines(filePath string) ([]string, error) {
    file, file_err := os.OpenFile(filePath, os.O_RDONLY, 0666)
    
    if file_err != nil {
        return nil, file_err
    }
    
    defer file.Close()
    
    var err error
    lines := make([]string, 0, 10)
    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        part, isPrefix, read_err := reader.ReadLine()
        if read_err != nil {
            err = read_err
            break;
        }
        
        buffer.Write(part)
        if isPrefix == false {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    
    if err == io.EOF {
        err = nil
    }

    return lines, err
}

// Starts replaying a bit log file
func (r *replayer) Replay(filePath string) (error) {
    lines, err := r.ReadLines(filePath)
    
    if err == nil {
        for _, line := range lines {
            bit, bit_err := BitFromLogString(line)
            if bit_err != nil {
                err = bit_err
            } else {
                r.consumer.Consume(bit)
            }
        }
    }
    
    return err
}