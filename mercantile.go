package mercantile

import (
    "fmt"
    "math"
    "strconv"
    "strings"
)

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

func Ul(tileid TileID) Point {
    //Returns the upper left (lon, lat) of a tile"""
    n := math.Pow(2.0, float64(tileid.Z))
    lon_deg := float64(tileid.X)/n*360.0 - 180.0
    lat_rad := math.Atan(math.Sinh(math.Pi * (1 - 2*float64(tileid.Y)/n)))
    lat_deg := (180.0 / math.Pi) * lat_rad
    return Point{lon_deg, lat_deg}
}

func Bounds(tileid TileID) Extrema {
    //Returns the (lon, lat) bounding box of a tile"""
    a := Ul(tileid)
    b := Ul(TileID{tileid.X + 1, tileid.Y + 1, tileid.Z})
    return Extrema{W: a.X, S: b.Y, E: b.X, N: a.Y}
}

func Tile(lng float64, lat float64, zoom int) TileID {
    // Returns the (x, y, z) tile"""

    lat = lat * (math.Pi / 180.0)
    n := math.Pow(2.0, float64(zoom))
    xtile := int(math.Floor((lng + 180.0) / 360.0 * n))
    ytile := int(math.Floor((1.0 - math.Log(math.Tan(lat)+(1.0/math.Cos(lat)))/math.Pi) / 2.0 * n))

    return TileID{int64(xtile), int64(ytile), uint64(zoom)}
}

func Tile_Geohash(lng float64, lat float64, zoom int) string {
    // Returns the (x, y, z) tile"""

    lat = lat * (math.Pi / 180.0)
    n := math.Pow(2.0, float64(zoom))
    xtile := int(math.Floor((lng + 180.0) / 360.0 * n))
    ytile := int(math.Floor((1.0 - math.Log(math.Tan(lat)+(1.0/math.Cos(lat)))/math.Pi) / 2.0 * n))

    return Tilestr(TileID{int64(xtile), int64(ytile), uint64(zoom)})
}

func Tilestr(tileid TileID) string {
    strval := fmt.Sprintf("%s/%s/%s", strconv.Itoa(int(tileid.X)), strconv.Itoa(int(tileid.Y)), strconv.Itoa(int(tileid.Z)))
    return strval
}

func Strtile(tileid string) TileID {
    vals := strings.Split(tileid, "/")
    x, _ := strconv.ParseInt(vals[0], 0, 64)
    y, _ := strconv.ParseInt(vals[1], 0, 64)
    z, _ := strconv.ParseInt(vals[2], 0, 64)
    //fmt.Print(x)

    return TileID{int64(x), int64(y), uint64(z)}
}

func Children(tile TileID) []TileID {

    a := TileID{tile.X * 2, tile.Y * 2, tile.Z + 1}
    b := TileID{tile.X*2 + 1, tile.Y * 2, tile.Z + 1}
    c := TileID{tile.X*2 + 1, tile.Y*2 + 1, tile.Z + 1}
    d := TileID{tile.X * 2, tile.Y*2 + 1, tile.Z + 1}

    return []TileID{a, b, c, d}
}
