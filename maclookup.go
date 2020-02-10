package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"net"
	"fmt"
	"strings"
)

func main() {
	inputmac := os.Args[1]
	if _, err := net.ParseMAC(inputmac); err != nil {
		println("Mac is invalid. Usage $maclookup MAC_ADDRESS")
	} else {
		mac := strings.Replace(inputmac, ":", "", -1)
		mac = mac[:6]
		mac = strings.ToUpper(mac)
		fptr := flag.String("fpath", "oui.txt", "file path to read from")
		flag.Parse()

		f, err := os.Open(*fptr)
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			if err = f.Close(); err != nil {
				log.Fatal(err)
			}
		}()
		s := bufio.NewScanner(f)
		for s.Scan() {
			if strings.Contains(s.Text(), mac) {
				vendor := s.Text()[22:]
				fmt.Print(inputmac+" ")
				fmt.Println(vendor)
			}
		}
		err = s.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}
