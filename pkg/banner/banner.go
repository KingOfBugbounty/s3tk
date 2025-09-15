package banner

import (
	"fmt"
	"math/rand"
	"time"
)

const Reset = "\033[0m"

func PrintBanner() {
	color1 := GetRandomColor()
	color2 := GetRandomColor()
	color3 := GetRandomColor()

	banner := `
%s███████╗██████╗ ███████╗ ██████╗ █████╗ ███╗   ██╗     ██╗ █████╗  █████╗  █████╗ ██╗  ██╗%s
%s██╔════╝╚════██╗██╔════╝██╔════╝██╔══██╗████╗  ██║     ██║██╔══██╗██╔══██╗██╔══██╗██║  ██║%s
%s███████╗ █████╔╝███████╗██║     ███████║██╔██╗ ██║     ██║███████║███████║███████║███████║%s
%s╚════██║ ╚═══██╗╚════██║██║     ██╔══██║██║╚██╗██║██   ██║██╔══██║██╔══██║██╔══██║██╔══██║%s
%s███████║██████╔╝███████║╚██████╗██║  ██║██║ ╚████║╚█████╔╝██║  ██║██║  ██║██║  ██║██║  ██║%s
%s╚══════╝╚═════╝ ╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝ ╚════╝ ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═╝%s

%s                    S3 Bucket Security Scanner for Bug Bounty%s
%s                          Created for Penetration Testing%s
`

	fmt.Printf(banner,
		color1, Reset,
		color2, Reset,
		color3, Reset,
		color1, Reset,
		color2, Reset,
		color3, Reset,
		color1, Reset,
		color2, Reset,
	)
	fmt.Println()
}

var colors = []string{
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
	"\033[91m", // Bright Red
	"\033[92m", // Bright Green
	"\033[93m", // Bright Yellow
	"\033[94m", // Bright Blue
	"\033[95m", // Bright Magenta
	"\033[96m", // Bright Cyan
}

func GetRandomColor() string {
	rand.Seed(time.Now().UnixNano())
	return colors[rand.Intn(len(colors))]
}
