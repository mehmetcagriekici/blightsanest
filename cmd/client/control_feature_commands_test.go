package main

import(
        "testing"
)

func TestControlFeatureCommands(t *testing.T) {
        cmdValidFeature    := []string{"get", "crypto"}
	cmdInvalidFeature  := []string{"get"}
        cmdValidSubFeature   := []string{"group", "crypto", "liquidity"}
	cmdInvalidSubFeature := []string{"group", "crypto"}

        resultFeatureValid := controlFeatureCommands(cmdValidFeature)
	if !resultFeatureValid {
	        t.Errorf("Expected: %v, got: %v", true, resultFeatureValid)
	}

        resultFeatureInvalid := controlFeatureCommands(cmdInvalidFeature)
	if resultFeatureInvalid {
	        t.Errorf("Expected: %v, got: %v", false, resultFeatureInvalid)
	}

        resultSubFeatureValid := controlFeatureSub(cmdValidSubFeature)
	if !resultSubFeatureValid {
	        t.Errorf("Expected: %v, got: %v", true, resultSubFeatureValid)
	}

        resultSubFeatureInvalid := controlFeatureSub(cmdInvalidSubFeature)
	if resultSubFeatureInvalid {
	        t.Errorf("Expected: %v, got: %v", false, resultSubFeatureInvalid)
	}
}