package bridge

import (
	"testing"
)

func TestBridge(t *testing.T) {
	pig := &pig{}
	cow := &cow{}
	pigCan := &canPlant{animal: pig}
	cowCan := &canPlant{animal: cow}
	pigSausage := &sausagePlant{animal: pig}
	cowSausage := &sausagePlant{animal: cow}
	pigCan.evaluation()
	cowCan.evaluation()
	pigSausage.evaluation()
	cowSausage.evaluation()
}

