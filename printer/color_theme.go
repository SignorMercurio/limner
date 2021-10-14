package printer

import (
	"github.com/SignorMercurio/limner/color"
	"github.com/spf13/viper"
)

var (
	KeyColor    = color.Red
	StringColor = color.Green
	BoolColor   = color.Yellow
	NumberColor = color.Yellow
	NullColor   = color.Cyan
	HeaderColor = color.Blue

	ColumnColors = []color.Color{
		color.White,
		color.Cyan,
	}
)

// str2color takes in a string and translates it to Color
func str2color(colorStr string) color.Color {
	switch colorStr {
	case "Black":
		return color.Black
	case "Red":
		return color.Red
	case "Green":
		return color.Green
	case "Yellow":
		return color.Yellow
	case "Blue":
		return color.Blue
	case "Magenta":
		return color.Magenta
	case "Cyan":
		return color.Cyan
	default:
		return color.White
	}
}

// slice2color is the slice version of str2color
func slice2color(colorSlice []interface{}) []color.Color {
	var colors []color.Color
	for _, v := range colorSlice {
		colors = append(colors, str2color(v.(string)))
	}
	return colors
}

func InitColorTheme() {
	for _, key := range viper.AllKeys() {
		value := viper.Get(key)
		switch val := value.(type) {
		case string:
			colorVal := str2color(val)
			switch key {
			case "key_color":
				KeyColor = colorVal
			case "string_color":
				StringColor = colorVal
			case "bool_color":
				BoolColor = colorVal
			case "number_color":
				NumberColor = colorVal
			case "null_color":
				NullColor = colorVal
			case "header_color":
				HeaderColor = colorVal
			}
		case []interface{}:
			ColumnColors = slice2color(val)
		}
	}
}
