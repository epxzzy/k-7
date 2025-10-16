package helper

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)
func GetTextInput() string {
	if false {
		fmt.Println("[TEXT] Using default input...")
		return "NAHHHH input" 
	}

	fmt.Print("INPUT:\n")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("\nError reading input:", err)
		return "nien inplu biatch error" 
	}

	input = strings.TrimSpace(input)
	if strings.ToLower(input) == "n" {
		return "no input given" 
	}
	
	return input
}


func ShowSpinner(text string, delay time.Duration, chars []string) func() {
	if delay == 0 {
		delay = 50 * time.Millisecond
	}
	if chars == nil {
		chars = []string{
			"⠉⠉",
			"⠈⠙",
			"⠀⠹",
			"⠀⢸",
			"⠀⣰",
			"⢀⣠",
			"⣀⣀",
			"⣄⡀",
			"⣆⠀",
			"⡇⠀",
			"⠏⠀",
			"⠋⠁",
		}
    /*

		chars = []string{
			"(===)",
			"( ==)",
			"(  =)",
			"(   )",
			"(=  )",
			"(== )",
      "(===)",
      "(===)",
      "(===)",
      "(===)",
      "(===)",

		}
    */

	}

	done := make(chan struct{})
	ticker := time.NewTicker(delay)
	x := 0

	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// spain without the a
				fmt.Fprintf(os.Stdout, "\r %s %s", chars[x], text)
				os.Stdout.Sync()
				x = (x + 1) % len(chars)
			case <-done:
				// end && clearline 
				fmt.Fprintf(os.Stdout, "\r \u2714  %s\n", text)
				os.Stdout.Sync()
				return
			}
		}
	}()

	// no spain
	return func() {
		close(done)
	}
}

// func main() {
//     stop := ShowSpinner("Loading", 50*time.Millisecond, nil)
//     defer stop() // Ensure spinner stops even if we exit early
//     time.Sleep(5 * time.Second)
// }
