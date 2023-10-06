package alert

func Red(s string) string {
  // Red text, takes string as parameter and return text in red color
  return RED+s+COLOR_RESET
}

func Blue(s string) string {
  // Blue text, takes string as parameter and return text in blue color
  return BLUE+s+COLOR_RESET
}

func Green(s string) string {
  // Green text, takes string as parameter and return text in green color
  return GREEN+s+COLOR_RESET
}

func Yellow(s string) string {
  // Yellow text, takes string as parameter and return text in yellow color
  return YELLOW+s+COLOR_RESET
}

func White(s string) string {
  // White text, takes string as parameter and return text in white color
  return WHITE+s+COLOR_RESET
}

func Cyan(s string) string {
  // Cyan text, takes string as parameter and return text in cyan color
  return CYAN+s+COLOR_RESET
}

func Magenta(s string) string {
  // Magenta text, takes string as parameter and return text in magenta color
  return MAGENTA+s+COLOR_RESET
}

func Black(s string) string {
  // Black text, takes string as parameter and return text in black color
  return BLACK+s+COLOR_RESET
}
