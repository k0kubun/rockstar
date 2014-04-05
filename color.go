package rockstar

import (
	"github.com/wsxiaoys/terminal/color"
)

func coloredUser(text string, star int) string {
	if star >= 5000 {
		return levelPrint(text, 6)
	} else if star >= 2000 {
		return levelPrint(text, 5)
	} else if star >= 500 {
		return levelPrint(text, 4)
	} else if star >= 100 {
		return levelPrint(text, 3)
	} else if star >= 50 {
		return levelPrint(text, 2)
	} else if star >= 10 {
		return levelPrint(text, 1)
	} else {
		return levelPrint(text, 0)
	}
}

func coloredRepository(text string, star int) string {
	if star >= 5000 {
		return levelPrint(text, 6)
	} else if star >= 1000 {
		return levelPrint(text, 5)
	} else if star >= 100 {
		return levelPrint(text, 4)
	} else if star >= 50 {
		return levelPrint(text, 3)
	} else if star >= 20 {
		return levelPrint(text, 2)
	} else if star >= 1 {
		return levelPrint(text, 1)
	} else {
		return levelPrint(text, 0)
	}
}

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
	default:
		return whitePrint(text)
	}
}

func coloredLanguage(language string) string {
	switch language {
	case "Ruby", "C++":
		return redPrint(language)
	case "JavaScript":
		return magentaPrint(language)
	case "Go":
		return yellowPrint(language)
	case "Scala", "C":
		return greenPrint(language)
	case "Shell", "PHP":
		return bluePrint(language)
	case "Perl", "Objective-C":
		return cyanPrint(language)
	default:
		return whitePrint(language)
	}
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
