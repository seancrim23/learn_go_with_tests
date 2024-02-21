package hello

import "fmt"

const (
	emptyString        = ""
	spanish            = "Spanish"
	french             = "French"
	helloString        = "Hello, "
	helloStringSpanish = "Hola, "
	helloStringFrench  = "Bonjour, "
)

func Hello(name, language string) string {
	if name == emptyString {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	prefix := ""
	switch language {
	case spanish:
		prefix = helloStringSpanish
	case french:
		prefix = helloStringFrench
	default:
		prefix = helloString
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Sean", ""))
}
