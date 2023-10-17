package alert

import "fmt"

func Error(text string) {
	// Alert error message, takes string as parameter and println text
	fmt.Println(
		fmt.Sprint(RED, "[-] ", COLOR_RESET, text),
	)
}

func Warn(text string) {
	// Alert warning message, takes string as parameter and println text
	fmt.Println(
		fmt.Sprint(CYAN, "[!] ", COLOR_RESET, text),
	)
}

func Success(text string) {
	// Alert success message, takes string as parameter and println text
	fmt.Println(
		fmt.Sprint(GREEN, "[+] ", COLOR_RESET, text),
	)
}

func Info(text string) {
	// Alert information message, takes string as parameter and println text
	fmt.Println(
		fmt.Sprint(BLUE, "[?] ", COLOR_RESET, text),
	)
}

func Impo(text string) {
	// Alert important message, takes string as parameter and println text
	fmt.Println(
		fmt.Sprint(YELLOW, "[*] ", COLOR_RESET, text),
	)
}
