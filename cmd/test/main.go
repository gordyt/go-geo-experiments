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
	input  = flag.String("--input-format", XYZ, "input format")
	output = flag.String("--output-format", ENU, "output format")
	help   = flag.Bool("--help", false, "if true, print help and exit")
)

const (
	AER = "AER"
	LLA = "LLA"
	XYZ = "XYZ"
	ENU = "ENU"
	//valConstructs = map[string]func
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
		log.Printf("[ERROR][main] expected 3 args, got %d", len(flag.Args()))
		os.Exit(1)
	}
	var args [3]float64
	for i, arg := range flag.Args() {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Printf("[ERROR][main], expected a string that looked like a float64, got %s", arg)
			os.Exit(1)
		}
		args[i] = val
	}
	// 161 Willow
	ref := geo.LLA{
		Latitude:  29.99504406861832,
		Longitude: -95.25508907331802,
		Altitude:  27.0,
	}
	log.Printf("[INFO][main] REF=%+v", ref)
	xyz := geo.XYZ{X: geo.Meters(args[0]), Y: geo.Meters(args[1]), Z: geo.Meters(args[2])}
	log.Printf("[INFO][main] XYZ=%+v", xyz)
	wsg := geo.WGS84()
	log.Printf("[INFO][main] WSG=%+v", wsg)
	enu := wsg.XYZToENU(ref, xyz)
	log.Printf("[INFO][main] ENU=%+v", enu)
	ned := wsg.XYZToNED(ref, xyz)
	log.Printf("[INFO][main] NED=%+v", ned)
	xyz2 := wsg.ENUToXYZ(ref, enu)
	log.Printf("[INFO][main] xyz2 = %+v", xyz2)

}
