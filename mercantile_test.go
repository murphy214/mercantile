package mercantile

import "testing"

// testing upper left bound
func Test_Ul(t *testing.T) {  
	tileid := TileID{0,0,0}
	expected_point := Point{-180, 85.05112877980659}
    point := Ul(tileid)
    if (point.X == expected_point.X && point.Y == expected_point.Y) == false {
       t.Errorf("Ul was incorrect, got: %v, want: %v.", point, expected_point)
    }
}


// testing bounds
func Test_Bounds(t *testing.T) {  
	tileid := TileID{0,0,0}
	expected_bds := Extrema{-180, 180, 85.05112877980659, -85.05112877980659}
    bds := Bounds(tileid)
    if (bds.N == expected_bds.N && bds.S == expected_bds.S && 
    	bds.E == expected_bds.E && bds.W == expected_bds.W) == false {
       t.Errorf("Bounds was incorrect, got: %v, want: %v.", bds, expected_bds)
    }
}

// testing tile
func Test_Tile(t *testing.T) {  
	lat,long := 40.0,-90.0
	zoom := 10
	
	expected_tileid := TileID{256, 387, 10}
    tileid := Tile(long,lat,zoom)
    if ((tileid.Z == expected_tileid.Z) && (tileid.Y == expected_tileid.Y) &&
    (tileid.X == expected_tileid.X)) == false {
       t.Errorf("Tile was incorrect, got: %v, want: %v.", tileid, expected_tileid)
    }
}


// testing tile
func Test_Tilestr(t *testing.T) {  
	tileid := TileID{0,0,0}

	
	expected_tilestr := "0/0/0"
    tilestr := Tilestr(tileid)
    if tilestr != expected_tilestr {
       t.Errorf("Tilestr was incorrect, got: %s, want: %s.", tilestr, expected_tilestr)
    }
}

// testing tile _geohash
func Test_Tile_Geohash(t *testing.T) {  
	lat,long := 40.0,-90.0
	zoom := 10
	
	expected_tilestr := "256/387/10"
    tilestr := Tile_Geohash(long,lat,zoom)
    if tilestr != expected_tilestr {
       t.Errorf("Tilestr was incorrect, got: %s, want: %s.", tilestr, expected_tilestr)
    }
}


// testing str_tile
func Test_Strtile(t *testing.T) {  
	strtile := "256/387/10"

	expected_tileid := TileID{256,387,10}
    tileid := Strtile(strtile)
    if ((tileid.Z == expected_tileid.Z) && (tileid.Y == expected_tileid.Y) &&
    (tileid.X == expected_tileid.X)) == false {
       t.Errorf("Strtile was incorrect, got: %v, want: %v.", tileid, expected_tileid)
    }
}


// testing str_tile
func Test_Children(t *testing.T) {  
	tileid := TileID{0,0,0}

	expected_children := []TileID{TileID{0, 0, 1}, TileID{1, 0, 1}, TileID{1, 1, 1}, TileID{0, 1, 1}}
    children := Children(tileid)

    for i := range children {
    	expected_tileid := expected_children[i]
    	tileid = children[i]

	    if ((tileid.Z == expected_tileid.Z) && (tileid.Y == expected_tileid.Y) &&
	    (tileid.X == expected_tileid.X)) == false {
	       t.Errorf("Children was incorrect, got: %v, want: %v.", tileid, expected_tileid)
	    }
	}
}

// testing str_tile
func Test_Center(t *testing.T) {  
	tileid := TileID{0,0,0}

	expected_center := []float64{0,0}
    center := Center(tileid)

    if ((expected_center[0] == center[0]) && (expected_center[1] == center[1])) == false {
       t.Errorf("Tile_Center was incorrect, got: %s, want: %s.", center, expected_center)
    }
}

// testing parent
func Test_Parent(t *testing.T) {
	child := TileID{256,387,10}

	expected_tileid := TileID{128, 193, 9}
	tileid := Parent(child)
	if ((tileid.Z == expected_tileid.Z) && (tileid.Y == expected_tileid.Y) &&
	(tileid.X == expected_tileid.X)) == false {
		t.Errorf("Children was incorrect, got: %v, want: %v.", tileid, expected_tileid)
	}
}	





