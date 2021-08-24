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

func downloadFile(filePath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
func main() {
	fileURL := "http://standards-oui.ieee.org/oui.txt"
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println(err)
		return
	}
	filePath := cacheDir + "/maclookup_oui.txt"

	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		fmt.Println("Downloading OUI cache")
		if err = downloadFile(filePath, fileURL); err != nil {
			panic(err)
		}
	}
	if len(os.Args) != 2 {
		println("Mac is invalid. Usage $maclookup MAC_ADDRESS")
		return
	}
	inputMAC := os.Args[1]
	if _, err := net.ParseMAC(inputMAC); err != nil {
		println("Mac is invalid. Usage $maclookup MAC_ADDRESS")
		return
	}
	mac := strings.Replace(inputMAC, ":", "", -1)
	mac = strings.Replace(mac, "-", "", -1)
	mac = strings.Replace(mac, ".", "", -1)
	mac = mac[:6]
	mac = strings.ToUpper(mac)
	fptr := flag.String("fpath", filePath, "file path to read from")
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
			fmt.Print(inputMAC + " ")
			fmt.Println(vendor)
			return
		}
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
