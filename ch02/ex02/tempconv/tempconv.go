// tempconvパッケージは摂氏（Celsius）と華氏（Fahrenheit）の温度計算を行います
package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	AbsoluteZeroF Fahrenheit = -459.67
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

// CToFは摂氏の温度を華氏へ変換します
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToCは華氏の温度を摂氏へ変換します
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
