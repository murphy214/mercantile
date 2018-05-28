# Mercantile - Simple Conversion of points to tile x,y,z.

[![GoDoc](https://godoc.org/github.com/murphy214/mercantile?status.svg)](https://godoc.org/github.com/murphy214/mercantile)

This project is a [port](https://github.com/mapbox/mercantile
) from a mercantile implementation in python.



# Usage 
 ```golang
 package main

import m "github.com/murphy214/mercantile"
import "fmt"


func main() {
	long,lat := -90.0,40.0
	point := []float64{long,lat}


	// getting a tile id
	tileid := m.Tile(point[0],point[1],10)
	fmt.Printf("%+v\n",tileid)
	// {X:256 Y:387 Z:10}

	// getting Bounds
	bds := m.Bounds(tileid)
	fmt.Printf("%+v\n",bds)
	// {W:-90 E:-89.6484375 N:40.17887331434696 S:39.909736234537185}

	// getting children of a given tileid
	children := m.Children(tileid)
	fmt.Println(children)
	// [{512 774 11} {513 774 11} {513 775 11} {512 775 11}]

	// getting parent
	parent := m.Parent(tileid)
	fmt.Println(parent)
	// {128 193 9}

	// getting center
	center := m.Center(tileid)
	fmt.Println(center)	
	// [-89.82421875 40.04430477444207]

	// getting a string tile
	tilestr := m.Tilestr(tileid)
	fmt.Println(tilestr)
	// 256/387/10

	// getting a string tile
	strtile := m.Strtile(tilestr)
	fmt.Println(strtile)
	// {256 387 10}
}
 ```
