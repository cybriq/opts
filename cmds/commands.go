package cmds

import (
	"github.com/p9c/opts"
)

// Commands are a slice of Command entries
type Commands []Command

// Command is a specification for a command and can include any number of subcommands
type Command struct {
	Name        string
	Description string
	Entrypoint  func(c *opts.Config) error
	Commands    Commands
}

// GetAllCommands returns all of the available command names
func (c Commands) GetAllCommands() (o []string) {
	for i := range c {
		o = append(o, c[i].Commands.GetAllCommands()...)
	}
	return
}

var tabs = "\t\t\t\t\t"

// Find the Command you are looking for. Note that the namespace is assumed to be flat, no duplicated names on different
// levels, as it returns on the first one it finds, which goes depth-first recursive
func (c Commands) Find(name string, hereDepth, hereDist int) (found bool, depth, dist int, cm *Command, e error) {
	if c == nil {
		dist = hereDist
		depth = hereDepth
		return
	}
	if hereDist == 0 {
		opts.D.Ln("searching for command:", name)
	}
	depth = hereDepth + 1
	opts.T.Ln(tabs[:depth]+"->", depth)
	dist = hereDist
	for i := range c {
		opts.T.Ln(tabs[:depth]+"walking", c[i].Name, depth, dist)
		if c[i].Name == name {
			opts.T.Ln(tabs[:depth]+"found", name, "at depth", depth, "distance", dist)
			found = true
			cm = &c[i]
			e = nil
			return
		} else {
			dist++
		}
		if found, depth, dist, cm, e = c[i].Commands.Find(name, depth, dist); opts.E.Chk(e) {
			opts.T.Ln(tabs[:depth]+"error", c[i].Name)
			return
		}
		if found {
			return
		}
	}
	opts.T.Ln(tabs[:hereDepth]+"<-", hereDepth)
	if hereDepth == 0 {
		opts.D.Ln("search text", name, "not found")
	}
	depth--
	return
}
