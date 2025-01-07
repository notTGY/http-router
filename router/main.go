package main

import (
	"fmt"
  "strings"
  "time"

	"github.com/tarm/serial"
)

func main() {
	config := &serial.Config{
		Name: "COM4", // Change this to your USB port, e.g., "/dev/ttyACM0" on Linux
		Baud: 115200,
	}
  var pico *serial.Port
  for {
    var err error
    pico, err = serial.OpenPort(config)
    if err != nil {
      fmt.Println("Error opening serial port:", err)
      time.Sleep(time.Second)
      continue
    }
    break
  }
	defer pico.Close()


  buf := make([]byte, 32)
  res := ""
  count := 0
  fmt.Printf("Ready\n")
  for i := 0; i < 10; i++ {
    n, err := pico.Read(buf)
    if err != nil {
      fmt.Println("Error reading from pico:", err)
      return
    }
    if n > 0 {
      str := string(buf[:n])
      res = res + str
      count += n
      if strings.HasSuffix(res, ": ") {
        break
      }
    }
  }
  fmt.Printf("Received %d: \"%s\"\n", count, res)

  if count != 10 {
    fmt.Printf("Early exit\n")
    return
  }

  body := []byte("Primeagen\r\n")
  n, err := pico.Write(body)
  if err != nil {
    fmt.Println("Error writing to pico:", err)
    return
  }
  fmt.Printf("Written %d bytes; %v\n", n, body)

  res = ""
  count = 0
  for i := 0; i < 10; i++ {
    n, err := pico.Read(buf)
    if err != nil {
      fmt.Println("Error reading from pico:", err)
      return
    }
    if n > 0 {
      str := string(buf[:n])
      res = res + str
      count += n
      if strings.HasSuffix(res, ": ") {
        break
      }
    }
  }
  fmt.Printf("Received %d: \"%s\"", count, res)
}
