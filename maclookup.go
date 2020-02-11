package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
func main() {
	fileUrl := "http://standards-oui.ieee.org/oui.txt"
	cachedir, err := os.UserCacheDir()
	filepath := cachedir + "/maclookup_oui.txt"

	if _, err = os.Stat(filepath); os.IsNotExist(err) {
		fmt.Println("Downloading OUI cache")
		if err = DownloadFile(filepath, fileUrl); err != nil {
			panic(err)
		}
	}
	inputmac := os.Args[1]
	if _, err := net.ParseMAC(inputmac); err != nil {
		println("Mac is invalid. Usage $maclookup MAC_ADDRESS")
	} else {
		mac := strings.Replace(inputmac, ":", "", -1)
		mac = mac[:6]
		mac = strings.ToUpper(mac)
		fptr := flag.String("fpath", filepath, "file path to read from")
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
				fmt.Print(inputmac + " ")
				fmt.Println(vendor)
				return
			}
		}
		err = s.Err()
		if err != nil {
			log.Fatal(err)
		}
	}
}
