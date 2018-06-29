package mercantile

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Extrema structure (Bounding Box)
type Extrema struct {
	W float64
	E float64
	N float64
	S float64
}

// TileID represents the id of the tile.
type TileID struct {
	X int64
	Y int64
	Z uint64
}

// Point represents a point in space.
type Point struct {
	X float64
	Y float64
}

// Returns the upper left (lon, lat) of a tile.
func Ul(tileid TileID) Point {
	n := math.Pow(2.0, float64(tileid.Z))
	lon_deg := float64(tileid.X)/n*360.0 - 180.0
	lat_rad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(tileid.Y)/n)))
	lat_deg := (180.0 / math.Pi) * lat_rad
	return Point{lon_deg, lat_deg}
}

//Returns the (lon, lat) bounding box of a tile.
func Bounds(tileid TileID) Extrema {
	a := Ul(tileid)
	b := Ul(TileID{tileid.X + 1, tileid.Y + 1, tileid.Z})
	return Extrema{W: a.X, S: b.Y, E: b.X, N: a.Y}
}

// Returns the (x, y, z) tile.
func Tile(lng float64, lat float64, zoom int) TileID {

	lat = lat * (math.Pi / 180.0)
	n := math.Pow(2.0, float64(zoom))
	xtile := int(math.Floor((lng + 180.0) / 360.0 * n))
	ytile := int(math.Floor((1.0 - math.Log(math.Tan(lat)+(1.0/math.Cos(lat)))/math.Pi) / 2.0 * n))

	return TileID{int64(xtile), int64(ytile), uint64(zoom)}
}

// Returns in string format like a geohash would be
func Tile_Geohash(lng float64, lat float64, zoom int) string {

	lat = lat * (math.Pi / 180.0)
	n := math.Pow(2.0, float64(zoom))
	xtile := int(math.Floor((lng + 180.0) / 360.0 * n))
	ytile := int(math.Floor((1.0 - math.Log(math.Tan(lat)+(1.0/math.Cos(lat)))/math.Pi) / 2.0 * n))

	return Tilestr(TileID{int64(xtile), int64(ytile), uint64(zoom)})
}

// Converts a tileid to tilestr representation
func Tilestr(tileid TileID) string {
	strval := fmt.Sprintf("%s/%s/%s", strconv.Itoa(int(tileid.X)), strconv.Itoa(int(tileid.Y)), strconv.Itoa(int(tileid.Z)))
	return strval
}

// From a tilestr representation back to a tileid
func Strtile(tileid string) TileID {
	vals := strings.Split(tileid, "/")
	x, _ := strconv.ParseInt(vals[0], 0, 64)
	y, _ := strconv.ParseInt(vals[1], 0, 64)
	z, _ := strconv.ParseInt(vals[2], 0, 64)

	return TileID{int64(x), int64(y), uint64(z)}
}

// Returns gets the children of a given child id.
func Children(tile TileID) []TileID {
	a := TileID{tile.X * 2, tile.Y * 2, tile.Z + 1}
	b := TileID{tile.X*2 + 1, tile.Y * 2, tile.Z + 1}
	c := TileID{tile.X*2 + 1, tile.Y*2 + 1, tile.Z + 1}
	d := TileID{tile.X * 2, tile.Y*2 + 1, tile.Z + 1}

	return []TileID{a, b, c, d}
}

// Returns the center of a given tileid.
func Center(tileid TileID) []float64 {
	bds := Bounds(tileid)
	return []float64{(bds.W + bds.E) / 2.0, (bds.N + bds.S) / 2.0}
}

// Returns the parent of a given tileid.
func Parent(tileid TileID) TileID {
	center := Center(tileid)
	return Tile(center[0], center[1], int(tileid.Z)-1)
}

// Converts a tileid to tilestr representation
// however this str conversion can be used for filenames
func TilestrFile(tileid TileID) string {
	strval := fmt.Sprintf("%s-%s-%s", strconv.Itoa(int(tileid.X)), strconv.Itoa(int(tileid.Y)), strconv.Itoa(int(tileid.Z)))
	return strval
}

// this function handles one of the many formats tilestr can be in
// basically every delimitter that could be withina  raw text tileid XYZ
func TileFromString(val string) TileID {
	var vals []string
	if strings.Contains(val, "-") {
		vals = strings.Split(val, "-")
	} else if strings.Contains(val, "/") {
		vals = strings.Split(val, "/")
	} else if strings.Contains(val, "_") {
		vals = strings.Split(val, "_")
	} else if strings.Contains(val, ",") {
		vals = strings.Split(val, ",")
	} else if strings.Contains(val, " ") {
		vals = strings.Split(val, " ")
	}

	x, _ := strconv.ParseInt(vals[0], 0, 64)
	y, _ := strconv.ParseInt(vals[1], 0, 64)
	z, _ := strconv.ParseInt(vals[2], 0, 64)
	return TileID{int64(x), int64(y), uint64(z)}
}

// returns a polygon from a given tile
func PolygonTile(tileid TileID) [][][]float64 {
	bds := Bounds(tileid)
	return [][][]float64{
		{
			{bds.E, bds.N},
			{bds.E, bds.S},
			{bds.W, bds.S},
			{bds.W, bds.N},
			{bds.E, bds.N},
		},
	}

}
