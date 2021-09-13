package command

import "testing"

func TestCommand(t *testing.T) {
	mouse := &Mouse{}
	copy := &CopyCommand{device: mouse}
	copyTo := &CopyToCommand{device: mouse}
	copyRequest := Request{command: copy}
	copyToRequest := Request{command: copyTo}
	
	copyRequest.Input()
	copyToRequest.Input()
}
