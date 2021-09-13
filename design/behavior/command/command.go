package command

type Request struct {
	command Command
}

func (r *Request) Input() {
	r.command.execute()
}

type Command interface {
	execute()
}

type CopyCommand struct {
	device Device
}

func (c *CopyCommand) execute() {
	c.device.Copy(12)
}

type CopyToCommand struct {
	device Device
}

func (c *CopyToCommand) execute() {
	c.device.CopyTo()
}

type Device interface {
	Copy(interface{})
	CopyTo() interface{}
}

type Mouse struct {
	some interface{}
}

func (m *Mouse) Copy(s interface{}) {
	m.some = s
}

func (m *Mouse) CopyTo() interface{} {
	return m.some
}
