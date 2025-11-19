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