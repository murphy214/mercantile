package main

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

type Bounds struct {
    w float64
    e float64
    n float64
    s float64
}

// TileID represents the id of the tile.
type TileID struct {
    x int64
    y int64
    z uint64
}

// Point represents a point in space.
type Size struct {
    deltaX float64
    deltaY float64
    linear float64
}

// Point represents a point in space.
type Point struct {
    X float64
    Y float64
}

func ul(tileid TileID) Point {
    //Returns the upper left (lon, lat) of a tile"""
    n := math.Pow(2.0, float64(tileid.z))
    lon_deg := float64(tileid.x)/n*360.0 - 180.0
    lat_rad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(tileid.y)/n)))
    lat_deg := (180.0 / math.Pi) * lat_rad
    return Point{lon_deg, lat_deg}
}

func bounds(tileid TileID) Bounds {
    //Returns the (lon, lat) bounding box of a tile"""
    a := ul(tileid)
    b := ul(TileID{tileid.x + 1, tileid.y + 1, tileid.z})
    return Bounds{w: a.X, s: b.Y, e: b.X, n: a.Y}
}

func tile(lng float64, lat float64, zoom int) TileID {
    // Returns the (x, y, z) tile"""

    lat = lat * (math.Pi / 180.0)
    n := math.Pow(2.0, float64(zoom))
    xtile := int(math.Floor((lng + 180.0) / 360.0 * n))
    ytile := int(math.Floor((1.0 - math.Log(math.Tan(lat)+(1.0/math.Cos(lat)))/math.Pi) / 2.0 * n))

    return TileID{int64(xtile), int64(ytile), uint64(zoom)}
}

func tile_geohash(lng float64, lat float64, zoom int) string {
    // Returns the (x, y, z) tile"""

    lat = lat * (math.Pi / 180.0)
    n := math.Pow(2.0, float64(zoom))
    xtile := int(math.Floor((lng + 180.0) / 360.0 * n))
    ytile := int(math.Floor((1.0 - math.Log(math.Tan(lat)+(1.0/math.Cos(lat)))/math.Pi) / 2.0 * n))

    return tilestr(TileID{int64(xtile), int64(ytile), uint64(zoom)})
}

func tilestr(tileid TileID) string {
    strval := fmt.Sprintf("%s/%s/%s", strconv.Itoa(int(tileid.x)), strconv.Itoa(int(tileid.y)), strconv.Itoa(int(tileid.z)))
    return strval
}

func strtile(tileid string) TileID {
    vals := strings.Split(tileid, "/")
    x, _ := strconv.ParseInt(vals[0], 0, 64)
    y, _ := strconv.ParseInt(vals[1], 0, 64)
    z, _ := strconv.ParseInt(vals[2], 0, 64)
    //fmt.Print(x)

    return TileID{int64(x), int64(y), uint64(z)}
}

func children(tile TileID) []TileID {

    a := TileID{tile.x * 2, tile.y * 2, tile.z + 1}
    b := TileID{tile.x*2 + 1, tile.y * 2, tile.z + 1}
    c := TileID{tile.x*2 + 1, tile.y*2 + 1, tile.z + 1}
    d := TileID{tile.x * 2, tile.y*2 + 1, tile.z + 1}

    return []TileID{a, b, c, d}
}
