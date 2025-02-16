package main

import (
	"fmt"
	"strings"
	"time"
  "net"
  "io"

  "go.bug.st/serial"

  "github.com/nottgy/http-router/router/device"
)

const DEVICE = device.DEVICE
const MTU = 256
const MAX_CHUNKS = 32

func onConnection(
  //pico net.Conn,
  pico serial.Port,
  conn net.Conn,
) error {
  //fmt.Printf("Request:\n")
	for i := 0; i < MAX_CHUNKS; i++ {
    buf := make([]byte, MTU)
    nread, err := conn.Read(buf)
    if err != nil {
      if err == io.EOF {
        break
      }
      return err
    }
    if nread == 0 {
      break
    }
    //fmt.Printf("[%d]%s", nread, string(buf[:nread]))

    to_write := make([]byte, nread)
    for idx := 0; idx < nread; idx++ {
      to_write[idx] = buf[idx]
    }

    // this breaks if you pass to_write
    nwrote, err := pico.Write(buf)
    if err != nil {
      fmt.Printf("pico write err\n")
      return err
    }
    if nwrote != nread {
      break
    }
    time.Sleep(100*time.Millisecond)
  }
  //fmt.Printf("Response:\n")

  response := ""
	//for i := 0; i < MAX_CHUNKS; i++ {
  for len(response) < MAX_CHUNKS * MTU {
    buf := make([]byte, MTU)
		nread, err := pico.Read(buf)
    //fmt.Printf("[%d]", nread)
		if err != nil {
			fmt.Println("Error reading from pico:", err)
			return err
		}
		if nread > 0 {
      to_write := make([]byte, nread)
      for idx := 0; idx < nread; idx++ {
        to_write[idx] = buf[idx]
      }
      nwrote, err := conn.Write(to_write)
      if err != nil {
        return err
      }
      if nwrote != nread {
        /*
        fmt.Printf(
          "Mismatched write-read: %d-%d\n",
          nwrote,
          nread,
        )
        */
      }

			str := strings.ReplaceAll(string(buf[:nread]), "\r", "")
      //fmt.Printf("%s", str)
      response = response + str
      http_part_index := strings.Count(response, "\n\n")
      if strings.Index(response, "\n") == strings.Index(response, "\n\n") {
        http_part_index++
      }
      N_HTTP_PARTS := 2
			if http_part_index >= N_HTTP_PARTS {
        //pico.Flush()
        //fmt.Printf("Finished\n")
				return nil
			}
		} else {
      //fmt.Printf("Finished\n")
      return nil
    }
	}
  //fmt.Printf("Too large\n")
  return nil
}

func main() {
	config := &serial.Mode{
		BaudRate: 115200,
	}
	var pico serial.Port
	//var pico net.Conn

  fmt.Printf("Started\n")
  ports, err := serial.GetPortsList()
  if err != nil {
    fmt.Println("Serial err: ", err)
  }
  if len(ports) == 0 {
    fmt.Println("No serial ports found!")
  }
  for _, port := range ports {
    fmt.Printf("Found port: %v\n", port)
  }
  fmt.Printf("Scan finished\n")


	for {
		var err error
    pico, err = serial.Open(DEVICE, config)
		//pico, err = net.Dial("unix", DEVICE)
		if err != nil {
			fmt.Println("Error opening serial port:", err)
			time.Sleep(time.Second)
			continue
		}
		break
	}
	defer pico.Close()


  port := 8000
  ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
  if err != nil {
    fmt.Println("Error listening:", err)
    return
  }
  fmt.Printf("listening on port %d\n", port)
	for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Error accepting:", err)
      return
    }
    // no `go` here, because usb can't multiplex
		err = onConnection(pico, conn)
    if err != nil {
      fmt.Printf("Following error occured:\n")
      fmt.Println(err)
    }
    conn.Close()
	}
}
