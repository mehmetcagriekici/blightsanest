package crypto

import(
        "log"
)

func ControlFeatureCommands(words []string) bool {
    // feature commands requires at least one more argument
    if len(words) < 2 {
        log.Printf("%s command requires at least one additional argument...\n", words[0])
        log.Printf("* %s crypto\n", words[0])
	return false
    }
    return true
}

// commands with at least 3 characters
func ControlFeatureSub(words []string) bool {
        // some feature commands require at least two more argument
	if len(words) < 3 {
	        log.Printf("%s command requires at least two additional arguments...\n", words[0])
		log.Printf("* %s group crypto\n", words[0])
		return false
	}
	return true
}
