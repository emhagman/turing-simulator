Turing Machine Simulator
================

A Turing Machine (Single Tape) Simulator. Works as Enumerator as well.
For more details, please review the code and any of the JSON files available in the src/ folder.

To run the program:
	go run turing.go <jsonfilename.json>

If you do not provide a filename, turing.json will be used.

The format for the JSON files should be fairly self explanatory.
An example file looks like this:

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

