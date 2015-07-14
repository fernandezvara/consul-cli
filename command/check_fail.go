package command

import (
	"strings"
)

type CheckFailCommand struct {
	Meta
	note		string
}

func (c *CheckFailCommand) Help() string {
	helpText := `
Usage: consul-cli check-fail [options] checkId

  Mark a local check as critical

Options:
` + c.ConsulHelp() + 
`  --note			Message to associate with check status
				(default: not set)
`

	return strings.TrimSpace(helpText)
}

func (c *CheckFailCommand) Run(args []string) int {
	flags := c.Meta.FlagSet(false)
	flags.StringVar(&c.note, "note", "", "")
	flags.Usage = func() { c.UI.Output(c.Help()) }

	if err := flags.Parse(args); err != nil {
		return 1
	}

	extra := flags.Args()
	if len(extra) < 1 {
		c.UI.Error("Check name must be specified")
		c.UI.Error("")
		c.UI.Error(c.Help())
		return 1
	}
	checkId := extra[0]

	consul, err := c.Client()
	if err != nil {	
		c.UI.Error(err.Error())
		return 1
	}

	client := consul.Agent()
	err = client.FailTTL(checkId, c.note)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	return 0
}

func (c *CheckFailCommand) Synopsis() string {
	return "Mark a local check as critical"
}
