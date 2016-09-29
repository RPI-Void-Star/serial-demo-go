package main

import (
    "log"
    "bufio"
    "fmt"
    "io"
    "os"
    "github.com/goburrow/serial"
)

func main() {
    fmt.Println("Please input dev port to use: ")
    in := bufio.NewReader(os.Stdin)
    message, err := in.ReadString('\n')
    if err != nil {
        log.Fatal(err)
    }
    message = message[:len(message)-1]

    port, err := serial.Open(&serial.Config{Address: message})
    if err != nil {
        log.Fatal(err)
    }
    // Auto closes when the function ends.
    defer port.Close()
    // Launch output copying on another thread.
    go io.Copy(os.Stdout, port)
    for {
        messagein, _ :=in.ReadString('\n')
        if _, err = port.Write([]byte(messagein)); err != nil {
            log.Fatal(err)
        }
    }
}
