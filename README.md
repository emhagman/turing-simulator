Turing Machine Simulator
================

A Turing Machine (Single Tape) Simulator. Works as Enumerator as well.
For more details, please review the code and any of the JSON files available in the src/ folder.

To run the program:
	
	go run turing.go <jsonfilename.json>

If you do not provide a filename, turing.json will be used. If at any point in time something is printed
to the "PRINTER" using the "Print" key in the Rules array, the Turing Machine will be assumed to be an Enumerator.
The Enumerator will stop after 50 characters has been printed. This will change very soon but you can modify the code
to change this cut-off value for now.

The format for the JSON files should be fairly self explanatory.
The following example is an ENUMERATOR of the language *(00)\**

	{
		"Input": "--------------------",
		"States": [
			{
				"Name": "Start",
				"Start": true,
				"Rules": [
					{
						"Input": "-",
						"Direction": "L",
						"ToState": "Q0",
						"ToTape": "$",
						"Print": ""
					}
				]
			},
			{
				"Name": "Q0",
				"Rules": [
					{
						"Input": "0",
						"Direction": "L",
						"ToState": "Q0",
						"ToTape": "",
						"Print": ""
					},
					{
						"Input": "$",
						"Direction": "R",
						"ToState": "Q1",
						"ToTape": "",
						"Print": ""
					}
				]
			},
			{
				"Name": "Q1",
				"Rules": [
					{
						"Input": "0",
						"Direction": "R",
						"ToState": "Q1",
						"ToTape": "",
						"Print": "0"
					},
					{
						"Input": "-",
						"Direction": "L",
						"ToState": "Q0",
						"ToTape": "0",
						"Print": "#"
					}
				]
			}
		]
	}

