package cli

import (
	"fmt"
	"os"
	"reflect"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

type CmdOption func(*cobra.Command)

func defaultCommand() *cobra.Command {
	return &cobra.Command{
		SilenceUsage: true,
	}
}

func Run(ops ...CmdOption) {
	if rootCmd == nil {
		rootCmd = defaultCommand()
	}

	for _, op := range ops {
		op(rootCmd)
	}

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}

// Cmd create and add a command. Zero must be of type 'struct { *one.Command }'.
func Cmd(zero interface{}, use string, short string, ops ...CmdOption) {
	if rootCmd == nil {
		rootCmd = defaultCommand()
	}

	// c is *UserType
	vc := reflect.ValueOf(zero)
	if vc.Kind() != reflect.Ptr {
		vc = reflect.New(vc.Type())
	}

	// cmd is *Cmd
	vcmd := vc.Elem().FieldByName("Command")
	vcmd.Set(reflect.New(vcmd.Type().Elem()))

	cmd, ok := vcmd.Interface().(*Command)
	if !ok {
		panic("zero must be of type 'struct { *one.Command }'")
	}
	ccmd := cmd.CorbaCommand(vc.Interface(), use, short)

	for _, op := range ops {
		op(ccmd)
	}
	rootCmd.AddCommand(ccmd)
}

func Func(run func(args []string), use string, short string, ops ...CmdOption) {
	if rootCmd == nil {
		rootCmd = defaultCommand()
	}

	ccmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(_ *cobra.Command, args []string) {
			run(args)
		},
	}

	for _, op := range ops {
		op(ccmd)
	}
	rootCmd.AddCommand(ccmd)
}

type Command struct {
	*cobra.Command
}

func (c *Command) CorbaCommand(ins interface{}, use string, short string) *cobra.Command {
	v := reflect.ValueOf(ins)

	ccmd := &cobra.Command{
		Use:          use,
		Short:        short,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			msg := "invalid Run method, shoule be 'Run(args []string)' or 'Run(args []string) error'"
			run := v.MethodByName("Run")
			if run.Kind() == reflect.Invalid {
				return fmt.Errorf(msg)
			}
			if run.Type().NumIn() != 1 {
				return fmt.Errorf(msg)
			}
			if run.Type().In(0) != reflect.TypeOf((*[]string)(nil)).Elem() {
				return fmt.Errorf(msg)
			}

			vargs := []reflect.Value{reflect.ValueOf(args)}
			var err error
			switch run.Type().NumOut() {
			case 0:
				run.Call(vargs)
			case 1:
				if run.Type().Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
					return fmt.Errorf(msg)
				}
				ret := run.Call(vargs)[0].Interface()
				if ret == nil {
					return nil
				}
				err = ret.(error)
			default:
				return fmt.Errorf(msg)
			}
			if err != nil {
				return err
			}
			return nil
		},
	}
	c.Command = ccmd

	msg := "invalid Build method, shoule be 'Build()'"
	build := v.MethodByName("Build")
	if build.Kind() != reflect.Invalid {
		if build.Type().NumIn() != 0 || build.Type().NumOut() != 0 {
			panic(msg)
		}
		build.Call(nil)
	}

	return ccmd
}

//
// Options
//

func Long(msg string) CmdOption {
	return func(c *cobra.Command) {
		c.Long = msg
	}
}

func Name(name string) CmdOption {
	return func(c *cobra.Command) {
		c.Use = name
	}
}
