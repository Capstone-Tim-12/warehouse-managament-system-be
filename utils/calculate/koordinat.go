package calculate

import (
	"fmt"
	"math"

	"github.com/spf13/cast"
)

const earthRadius = 6371.0 // Radius Bumi dalam kilometer

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	// Mengonversi koordinat latitude dan longitude dari derajat ke radian
	lat1, lon1, lat2, lon2 = toRadians(lat1), toRadians(lon1), toRadians(lat2), toRadians(lon2)

	// Haversine formula
	dlat := lat2 - lat1
	dlon := lon2 - lon1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Menghitung jarak
	distance := earthRadius * c

	resStr := fmt.Sprintf("%.2f", distance)
	return cast.ToFloat64(resStr)
}

// Mengonversi derajat menjadi radian
func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}
