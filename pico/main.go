package main

import (
	"machine"
	"strings"
	"time"
)

const MTU = 128
const MAX_CHUNKS = 32

func main() {
	serial := machine.Serial

	led := machine.LED
	led.Configure(machine.PinConfig{
		Mode: machine.PinOutput,
	})

	led.High()
	time.Sleep(5 * time.Second)
	led.Low()

	for {
		/*
			_, err := serial.Write([]byte("the name: "))
			if err != nil {
				println("Write err: ", err)
			}
			blink(led)
		*/

    request := ""
    for i := 0; i < MAX_CHUNKS; i++ {
      for serial.Buffered() == 0 { }
      buf := make([]byte, MTU)
      j := 0
      for ; serial.Buffered() > 0 && j < MTU; j++ {
        inByte, _ := serial.ReadByte()

        /*
        err := serial.WriteByte(inByte)
        if err != nil {
          println("WriteByte err: ", err)
        }
        */

        buf[j] = inByte
      }
      request = request + string(buf[:j])

      parts := strings.Count(
        strings.ReplaceAll(request, "\r", ""),
        "\n\n",
      )
      if parts >= 1 {
        break
      }
    }
		blink(led)

		response := `HTTP/1.1 200 OK

Hello, World!

`
		_, err := serial.Write([]byte(response))
		if err != nil {
			println("Write err: ", err)
		}
	}
}

func blink(p machine.Pin) {
	p.High()
	time.Sleep(500 * time.Millisecond)
	p.Low()
	time.Sleep(500 * time.Millisecond)
}
