#  ggpt
<img src="https://github.com/islewis/ggpt/raw/main/logo/logo1.png" width="200">
ggpt is a simple tool for accessing GPT on the command line, written in Go.

# Usage Examples

GPT is great for creating sample data. This works particularly well when combined with a traditional CLI workflow.  
```
$ ggpt prompt "Output a sample csv file of new cars for sale. Include car model, cost, and mpg. Include no commas in the data. Give no explaination"
model,cost,mpg
Toyota Camry,25000,32
Honda Civic,22000,36
Ford Fusion,26000,28
```
If the output is suffient, `ggpt last` will return it again.
```
$ ggpt last > for_sale.csv
```
Taking advantage of command substitution, GPT can easily parse over this new data.
```
$ ggpt prompt "Output an example python file that reads in a file named for_sale.csv with data $(cat for_sale.csv), and prints the cost per mpg for each car model. Give no formatting or explanation."
import csv

with open('for_sale.csv', 'r') as file:
    reader = csv.reader(file)
    next(reader) #Skip header row
    for row in reader:
        cost_per_mpg = float(row[1]) / float(row[2])
        print(row[0] + ": " + str(cost_per_mpg))

$ ggpt last > cars.py && python3 cars.py
Toyota Camry: 781.25
Honda Civic: 611.1111111111111
Ford Fusion: 928.5714285714286
``` 
-------

Another possible usage of GPT is regex help (we all need it)
```
$ ggpt prompt "Output a regex that finds all december dates from 2011 to 2019. Use format month/day/year. Include no explanation."
\b12\/([1-9]|[1-2][0-9]|3[0-1])\/(201[1-9])\b
```
Cool, that's certainly a regex. We know GPT can (and will) make stuff up. Lets save that regex to an enviornment variable, and test it.
```
$ export REGEX=$(ggpt last)
$ ggpt prompt "Output a list of 20 dates, one on each line. Use format month/day/year. Include dates from the last ten years." | tee dates.csv
Assuming the current year is 2021, here is a list of 20 dates in the format month/day/year, including dates from the last ten years:

1. 03/14/2011
2. 11/11/2012
3. 05/16/2013
4. 08/08/2014
5. 01/27/2015
6. 06/12/2016
7. 12/31/2017
8. 09/22/2018                                                                                         
9. 04/09/2019
10. 11/03/2020
11. 01/10/2021
12. 04/22/2021
13. 07/04/2021
14. 09/15/2021
15. 11/25/2021
16. 01/01/2022
17. 03/17/2022
18. 06/19/2022
19. 08/14/2022
20. 12/01/2022

$ grep -E "$REGEX" dates.txt
7. 12/31/2017
```
Passes the eye test, pretty neat. What's it doing though?
```
ggpt prompt "can you explain what regex $REGEX is doing?"
The regex \b12\/([1-9]|[1-2][0-9]|3[0-1])\/(201[1-9])\b is looking for a date in the format "dd/mm/yyyy" where the month is specified as "12".

Here is a breakdown of the different parts of the regex:

\b - Matches the start or end of a word boundary

12\/ - Matches the string "12/" exactly

([1-9]|[1-2][0-9]|3[0-1]) - Matches a string of digits representing the day of the month from 1 to 31

\/ - Matches a forward slash character

(201[1-9]) - Matches a string of digits representing the year from 2011 to 2019

\b - Matches the end of a word boundary

Overall, the regex searches for any string of text that meets the aforementioned conditions, and returns that string as a match.
```


## Installation

Make sure go is installed using `go version`. If not, [install go](https://go.dev/doc/install).
```
$ go version      
go version go1.20.0 linux/amd64
```
Next, install ggpt using `go install`:
```
$ go install github.com/islewis/ggpt@latest
```
Thats it! To confirm ggpt is downloaded, run `ggpt --help`:
```
$ ggpt --help
ggpt is a CLI tool to interact with OpenAI's GPT language model. ggpt wraps OpenAI's completion feature, via their API, outputting the result directly in the terminal.

Usage:
  ggpt [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  configure   Configure your OpenAI API key
  help        Help about any command
  last        Returns the output of the previous query.
  prompt      Call GPT autocomplete with the given string as a prompt

Flags:
  -h, --help     help for ggpt
  -t, --toggle   Help message for toggle

Use "ggpt [command] --help" for more information about a command.
```
## Usage
todo
