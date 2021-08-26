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
	input  = flag.String("--input-format", GPS, "input format")
	output = flag.String("--output-format", NED, "output format")
	help   = flag.Bool("--help", false, "if true, print help and exit")
)

const (
	AER = "AER"
	LLA = "LLA"
	XYZ = "XYZ"
	ENU = "ENU"
	NED = "NED"
	GPS = "GPS"			// latitude, longitude, altitude above mean sea level
	COE = 6371146
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
	wsg := geo.WGS84()
	lla := geo.LLA{Latitude: geo.Degrees(args[0]), Longitude: geo.Degrees(args[1]), Altitude: geo.Meters(args[2] + COE)}
	xyz := wsg.LLAToXYZ(lla)
	log.Printf("[INFO][main] XYZ (from args)=%+v", xyz)
	log.Printf("[INFO][main] WSG (just defining receiver)=%+v", wsg)
	// 161 Willow
	ref := lla
	log.Printf("[INFO][main] LLA REF (161 Willow)=%+v", ref)
	enu := wsg.XYZToENU(ref, xyz)
	log.Printf("[INFO][main] ENU=%+v", enu)
	ned := wsg.XYZToNED(ref, xyz)
	log.Printf("[INFO][main] NED=%+v", ned)
	xyz2 := wsg.ENUToXYZ(ref, enu)
	log.Printf("[INFO][main] XYZ (converted back from ENU)= %+v", xyz2)
	xyz3 := wsg.NEDToXYZ(ref, ned)
	log.Printf("[INFO][main] XYZ (converted back from NED)= %+v", xyz3)
}
