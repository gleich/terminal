package output

import "github.com/muesli/termenv"

type Colors struct {
	Blue  termenv.Color
	Green termenv.Color
	Grey  termenv.Color
}

func NewColors(colorProfile termenv.Profile) Colors {
	return Colors{
		Green: colorProfile.Color("#30CE75"),
		Blue:  colorProfile.Color("#2b95ff"),
		Grey:  colorProfile.Color("#383737"),
	}
}
