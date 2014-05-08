package main
import (
	"fmt"
	"encoding/json"
	"os"
	"bytes"
	"io"
)

// Structure of our rules
type Rule struct {
	Input string
	Direction string
	ToTape string
	ToState string
	Print string
}

// Structure of our States
type State struct {
    Name string
    Start bool
    Accept bool
    Rules []Rule
}

// Structure for whole thing
type Config struct {
	Input string
	States []State
}

// Print out the tape with the head as well
func PrintTape(tape string, head int) {
	fmt.Println("Current Tape:")
	for i := 0; i < len(tape); i++ {
		fmt.Printf("%s", string(tape[i]))
	}
	fmt.Printf("\n")
	for i := 0; i < len(tape); i++ {
		if (i == head) {
			fmt.Printf("^")
		} else {
			fmt.Printf(" ")
		}
	}
	fmt.Println("")
}

// Read in a text file
func ReadText(filename string) (string) {
	buf := bytes.NewBuffer(nil)
	f, _ := os.Open(filename) // Error handling elided for brevity.
	io.Copy(buf, f)           // Error handling elided for brevity.
	f.Close()
	s := string(buf.Bytes())
	return s
}

// Get state name by index
func GetStateName(States []State, Index int) (string) {
	return States[Index].Name
}

// Get state array index
func GetStateIndex(States []State, Name string) (int) {
	for i := 0; i < len(States); i++ {
		if (States[i].Name == Name) {
			return i
		}
	}
	return -1
}

// Main function
func main() {

	// Read our text file passed in from args our default
	var JsonFileName string
	if (len(os.Args) > 1) {
		JsonFileName = os.Args[1]
		fmt.Printf("Using %s as Turring Machine file!\n", JsonFileName)
	} else {
		fmt.Println("Defaulting to turring.json!")
	}

	// Read in our text file
	if (JsonFileName == "") {
		JsonFileName = "turring.json"
	}
	JsonFileText := ReadText(JsonFileName)

	// Read into array of states
	var config Config
	sbytes := []byte(JsonFileText)
	json.Unmarshal(sbytes, &config)

	// Vars
	head := 0
	accept := -1
	state := -1
	input := string(config.Input)
	var printer string

	// Find our start
	for i := 0; i < len(config.States); i++ {
	    if (config.States[i].Start == true) {
	    	state = i
	    }
	    if (config.States[i].Accept == true) {
	    	accept = i
	    }
	    if (state != -1 && accept != -1) {
	    	break
	    }
	}

	// Using input
	fmt.Println("Using Input: ", config.Input)
	fmt.Println("Starting State: ", config.States[state].Name)
	fmt.Println("")

	// Do it while we have not accepted
	for state != accept && state != -1 {

		// Print states
		fmt.Printf("Current State: %s\n", config.States[state].Name)
		var p string
		if (printer == "") {
			p = "[EMPTY]"
		} else {
			p = printer
		}
		fmt.Printf("Current Printer: %s\n", p)
		PrintTape(input, head)

		// Store the current state to see if we are stuck at end
		CurrentState := state

		// Go through the rules
		UsedRule := false
		for i := 0; i < len (config.States[state].Rules); i++ {

			// The current rule
			var MyRule Rule
			MyRule = config.States[state].Rules[i]

			// Check for condition
			fmt.Printf("Checking To See If [%s] == [%s]\n", MyRule.Input, string(input[head]))
			if (MyRule.Input == string(input[head])) {

				// We used one!
				UsedRule = true

				// It Does
				fmt.Printf("Moving To State: [%s]\n", MyRule.ToState)

				// Move to our next state
				if (MyRule.ToTape != "") {
					fmt.Printf("Storing Value [%s] Into Tape @ Head\n", MyRule.ToTape)

					// Create new "tape" with our new value
					// Golang can't directly assign
					b := []byte(input)
					copy(b[head:], []byte(string(MyRule.ToTape)))
					input = string(b)

					PrintTape(input, head)
				}

				// Move our head depending on direction
				if (MyRule.Direction == "R") {
					head += 1
				} else {
					head -= 1
				}

				// Stay in rage
				if (head < 0) {
					head = 0
				}

				// Move to our next state
				state = GetStateIndex(config.States, MyRule.ToState)

				// Print if we need to
				printer += MyRule.Print
				break
			}
		}

		fmt.Println("")

		// Are we stuck?
		if ((CurrentState == state && !UsedRule) || state == -1) {
			fmt.Println("Failed @ State:", GetStateName(config.States, state))
			if (state == -1) {
				fmt.Println("You tried to reach a state that does not exist!")
			}
			os.Exit(0)
		}

		// Check for enumerator and only run for 100 iterations
		if (len(printer) > 50) {
			fmt.Println("Stopping enumerator after 50 chars!")
			os.Exit(0)
		}
	}
}