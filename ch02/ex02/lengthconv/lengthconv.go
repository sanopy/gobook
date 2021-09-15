// lengthconvパッケージはフィートとメートルの長さ計算を行います
package lengthconv

import "fmt"

type Feet float64
type Meter float64

const (
	OneFeetInMeters Meter = 0.3048
	OneMeterInFeet  Feet  = 3.28084
)

func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

// FToMはフィートをメートルへ変換します
func FToM(f Feet) Meter { return Meter(f * 0.3048) }

// MToFはメートルをフィートへ変換します
func MToF(m Meter) Feet { return Feet(m / 0.3048) }
