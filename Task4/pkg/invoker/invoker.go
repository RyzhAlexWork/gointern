package invoker

type command interface {
	Execute() (done bool)
}

// Invoker ...
type Invoker interface {
	StoreCommand(command command)
	UnStoreCommand()
	Execute() []bool
}

type invoker struct {
	commands []command
}

// StoreCommand adds command.
func (n *invoker) StoreCommand(command command) {
	n.commands = append(n.commands, command)
}

// UnStoreCommand removes command.
func (n *invoker) UnStoreCommand() {
	if len(n.commands) != 0 {
		n.commands = n.commands[:len(n.commands)-1]
	}
}

// Execute all commands.
func (n *invoker) Execute() []bool {
	var result []bool
	for _, command := range n.commands {
		result = append(result, command.Execute())
	}
	return result
}

// NewInvoker create invoker implementation for interface Invoker.
func NewInvoker() Invoker {
	return &invoker{}
}
