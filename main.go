//package main
//
//import (
//	"fmt"
//	"github.com/spf13/cobra"
//	"strings"
//)
//
////func initializeBaz(ctx context.Context) (engine.Baz, error) {
////	wire.Build(engine.MegaSet)
////	return engine.Baz{}, nil
////}
//
//func main() {
//	//cobra.AddTemplateFuncs()
//	//baz, err := a.InitializeBaz(context.Background())
//	//if err != nil {
//	//	panic(err)
//	//}
//	//fmt.Println(baz)
//	//engine := cmd2.NewCmd()
//	//engine.Run()
//}

package main

import (
	cmd2 "github.com/simon/mengine/infrastructure/engine"
	_ "github.com/sourcegraph/conc"
)

func main() {

	cmd := cmd2.NewCmd()
	cmd.Run()
	//
	//	var echoTimes int
	//
	//	var cmdPrint = &cobra.Command{
	//		Use:   "print [string to print]",
	//		Short: "Print anything to the screen",
	//		Long: `print is for printing anything back to the screen.
	//For many years people have printed back to the screen.`,
	//		Args: cobra.MinimumNArgs(1),
	//		Run: func(engine *cobra.Command, args []string) {
	//			fmt.Println("Print: " + strings.Join(args, " "))
	//		},
	//	}
	//
	//	var cmdEcho = &cobra.Command{
	//		Use:   "echo [string to echo]",
	//		Short: "Echo anything to the screen",
	//		Long: `echo is for echoing anything back.
	//Echo works a lot like print, except it has a child command.`,
	//		Args: cobra.MinimumNArgs(1),
	//		Run: func(engine *cobra.Command, args []string) {
	//			fmt.Println("Print: " + strings.Join(args, " "))
	//		},
	//	}
	//
	//	var cmdTimes = &cobra.Command{
	//		Use:   "times [# times] [string to echo]",
	//		Short: "Echo anything to the screen more times",
	//		Long: `echo things multiple times back to the user by providing
	//a count and a string.`,
	//		Args: cobra.MinimumNArgs(1),
	//		Run: func(engine *cobra.Command, args []string) {
	//			for i := 0; i < echoTimes; i++ {
	//				fmt.Println("Echo: " + strings.Join(args, " "))
	//			}
	//		},
	//	}
	//
	//	cmdTimes.Flags().IntVarP(&echoTimes, "times", "t", 1, "times to echo the input")
	//
	//	var rootCmd = &cobra.Command{Use: "app1111"}
	//	rootCmd.AddCommand(cmdPrint, cmdEcho)
	//	cmdEcho.AddCommand(cmdTimes)
	//	rootCmd.Execute()
}
