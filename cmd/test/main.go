package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"pault.ag/go/geo"
)

var (
	input = flag.String("--input-format", XYZ, "input format")
	output = flag.String("--output-format", ENU, "output format")
	help = flag.Bool("--help", false, "if true, print help and exit")
)

const (
	AER = "AER"
	LLA = "LLA"
	XYZ = "XYZ"
	ENU = "ENU"
	valConstructs = map[string]func
)



func usage() {
	fmt.Println("TODO help")
}

func main() {
	flag.Parse()
	log.Printf("[INFO][main] Input format: %s, Output Format=%s, Args=%s",
		strings.ToUpper(*input),
		strings.ToUpper(*output),
		flag.Args())
	if len(flag.Args()) != 3 {
		log.Println("[ERROR][main] expected 3 args, got %d", len(flag.Args()))
		os.Exit(1)
	}
	var args [3]float64{}
	for i, arg := range flag.Args() {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Printf("[ERROR][main], expected a string that looked like a float64, got %s", arg)
			os.Exit(1)
		}
		args[i] = val
	}
	input := geo.XYZ{X: geo.Meters(args[0]), Y: geo.Meters(args[1]), Z: geo.Meters(args[2])}
	wsg := geo.WGS84()
	lla := wsg.XYZToLLA(input)
	wsg.LLAToENU(lla)


}
