// weightconvパッケージはポンドとキログラムの重さ計算を行います
package weightconv

import "fmt"

type Pound float64
type Kilogram float64

const (
	OneKilogramInPound Pound    = 2.204623
	OnePoundInKilogram Kilogram = 0.453592
)

func (p Pound) String() string     { return fmt.Sprintf("%glb", p) }
func (kg Kilogram) String() string { return fmt.Sprintf("%gkg", kg) }

// PToKgは摂氏の温度を華氏へ変換します
func PToKg(p Pound) Kilogram { return Kilogram(p / 2.204623) }

// KgToPは華氏の温度を摂氏へ変換します
func KgToP(kg Kilogram) Pound { return Pound(kg * 2.204623) }
