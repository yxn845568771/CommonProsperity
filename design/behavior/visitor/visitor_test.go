package visitor

import (
	"fmt"
	"testing"
)

func TestVisitor(t *testing.T) {
	nameUpdater := nameUpdater{
		pigName: "111",
		cowName: "222",
	}
	pig := &pig{name: "333"}
	cow := &cow{name: "444"}
	nameUpdater.visitForPig(pig)
	nameUpdater.visitForCow(cow)
	fmt.Println(pig.name,cow.name)
}
