package main

import (
	"machine"
  "time"
)

func main() {
  time.Sleep(5*time.Second)

  serial := machine.Serial

  led := machine.LED
  led.Configure(machine.PinConfig{
    Mode: machine.PinOutput,
  })

  for {
    data := make([]byte, 1)
    _, err := serial.Write([]byte("the name: "))
    if err != nil {
      println("Write err: ", err)
    }
    blink(led)

    for {
      if serial.Buffered() > 0 {
        inByte, _ := serial.ReadByte()
        err = serial.WriteByte(inByte)
        if err != nil {
          println("WriteByte err: ", err)
        }
        if inByte == byte(13) {
          break
        }
        data = append(data, inByte)
      }
    }
    blink(led)

    output := "\r\nThe name is the " + string(data) + "\r\n"
    _, err = serial.Write([]byte(output))
    if err != nil {
      println("Write err: ", err)
    }
  }
}

func blink(p machine.Pin) {
  p.High()
  time.Sleep(500*time.Millisecond)
  p.Low()
  time.Sleep(500*time.Millisecond)
}
