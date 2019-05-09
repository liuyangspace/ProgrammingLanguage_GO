package os

import (
	"strings"
	"os"
	"fmt"
	"flag"
	"os/exec"
)

func TestReadCommandArgs()  {
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Good Morning", who)
}

func TestParseCommandArgs()  {
	NewLine := flag.Bool("n", false, "print newline") // echo -n flag, of type *bool
	const (
		Space   = " "
		Newline = "\n"
	)
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine { // -n is parsed, flag becomes true
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

func TestRunSystemCommand()  {
	cmd := exec.Command("c:\\dir")  // this opens a gedit-window
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error %v executing command!", err)
		os.Exit(1)
	}
	fmt.Printf("The command is %v", cmd)
}

func main()  {
	TestReadCommandArgs()
	TestParseCommandArgs()
}