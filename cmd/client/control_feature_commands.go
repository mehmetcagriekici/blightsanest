package main

import(
        "log"
)

func controlFeatureCommands(words []string) bool {
    // feature commands requires at least one more argument
    if len(words) < 2 {
        log.Printf("%s command requires at least one additional argument...", words[0])
        log.Printf("* %s crypto", words[0])
	return false
    }
    return true
}