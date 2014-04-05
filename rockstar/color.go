package rockstar

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
)

func levelPrint(text string, level int) string {
	switch level {
	case 0:
		return whitePrint(text)
	case 1:
		return cyanPrint(text)
	case 2:
		return bluePrint(text)
	case 3:
		return greenPrint(text)
	case 4:
		return yellowPrint(text)
	case 5:
		return magentaPrint(text)
	case 6:
		return redPrint(text)
	}
	return whitePrint(text)
}

func colorTest() {
	fmt.Printf("%s", levelPrint("level", 6))
	fmt.Printf("%s", levelPrint("level", 5))
	fmt.Printf("%s", levelPrint("level", 4))
	fmt.Printf("%s", levelPrint("level", 3))
	fmt.Printf("%s", levelPrint("level", 2))
	fmt.Printf("%s", levelPrint("level", 1))
	fmt.Printf("%s", levelPrint("level", 0))
}

func blackPrint(text string) string {
	return color.Sprintf("@k%s", text)
}

func redPrint(text string) string {
	return color.Sprintf("@r%s", text)
}

func greenPrint(text string) string {
	return color.Sprintf("@g%s", text)
}

func yellowPrint(text string) string {
	return color.Sprintf("@y%s", text)
}

func bluePrint(text string) string {
	return color.Sprintf("@b%s", text)
}

func magentaPrint(text string) string {
	return color.Sprintf("@m%s", text)
}

func cyanPrint(text string) string {
	return color.Sprintf("@c%s", text)
}

func whitePrint(text string) string {
	return color.Sprintf("@w%s", text)
}
