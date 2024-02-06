package backendgis

import (
	"fmt"
	"testing"
)

var privatekey = ""
var publickey = ""
var encode = ""
var dbname = "geojson"
var collname = "cijambe"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("mongoenv", dbname)
	coordinates := Point{
		Coordinates: []float64{
			107.70123687792598,-6.9133580309376015,
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

// func TestGeoWithin(t *testing.T) {
// 	mconn := SetConnection("mongoenv", dbname)
// 	coordinates := Polygon{
// 		Coordinates: [][][]float64{
// 			{
// 				{ 107.69161532452927, -6.909192149952432}, 
// 				{107.6928270495055, -6.909088574785045}, 
// 				{ 107.69101448076144, -6.90926005434234}, 
// 				{ 107.69161532452927, -6.909192149952432 }
// 			},
// 		},
// 	}
// 	datagedung := GeoWithin(mconn, collname, coordinates)
// 	fmt.Println(datagedung)
// }

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("mongoenv", "geojson", "cijambe")
	coordinates := Point{
		Coordinates: []float64{
			107.70123687792598,-6.9133580309376015,
		},
	}
	datagedung := Near(mconn, "GET", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("mongoenv", "geojson")
	coordinates := Point{
		Coordinates: []float64{
			107.70123339970226, -6.91335864870355,
		},
	}
	datagedung := NearSphere(mconn, "GET", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("mongoenv", "geojson")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{95.32345678901234, 5.567890123456789},
			{95.32355678901234, 5.567990123456789},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}
