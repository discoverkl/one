package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var Root = NewCmd()

type FlagSet = pflag.FlagSet

type Cmd struct {
	cmd *cobra.Command

	Use string
	Short string
	Long string
}

func NewCmd() *Cmd {
	return &Cmd{
		cmd: &cobra.Command{},
	}
}

func (c *Cmd) Flags() *FlagSet{
	cc := c.cmd
	return cc.Flags()
}

func (c *Cmd) Run(r func(args []string) error) {
	c.refresh()
	cc := c.cmd
	cc.SilenceUsage = true
	cc.SilenceErrors = true

	cc.RunE = func(_ *cobra.Command, args []string) error {
		return r(args)
	}
	if err := cc.Execute(); err != nil {
		log.Fatal(err)
	}
}

func (c *Cmd) refresh() {
	cc := c.cmd
	cc.Use = c.Use
	cc.Short = c.Short
	cc.Long = c.Long
}