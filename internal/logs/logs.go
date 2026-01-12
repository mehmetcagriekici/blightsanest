package logs

import(
       "os"
       "fmt"
       "bufio"
       "strings"
)

// get cli input
func GetInput() []string {
        fmt.Print("> ")
        scanner := bufio.NewScanner(os.Stdin)
	if isScanned := scanner.Scan(); !isScanned {
	        return nil
	}
	inp := strings.TrimSpace(scanner.Text())
	return strings.Fields(inp)
}

// control command lengths
func ControlFeatureCommands(words []string) bool {
    // feature commands requires at least one more argument
    if len(words) < 2 {
        fmt.Printf("%s command requires at least one additional argument...\n", words[0])
        fmt.Printf("* %s crypto\n", words[0])
	return false
    }
    return true
}

// commands with at least 3 characters
func ControlFeatureSub(words []string) bool {
        // some feature commands require at least two more argument
	if len(words) < 3 {
	        fmt.Printf("%s command requires at least two additional arguments...\n", words[0])
		fmt.Printf("* %s group crypto\n", words[0])
		return false
	}
	return true
}