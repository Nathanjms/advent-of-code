# Advent of Code

Go the the relevant year and look at the README.md file for that year for details of that year's solutions.

I started with 2023, and am now working my way from 2015. Keep an eye on the list below for progress.

- [x] 2015
  - Go
- [ ] 2016
- [ ] 2017
- [ ] 2018
- [ ] 2019
- [ ] 2020
- [ ] 2021
- [ ] 2022
- [x] 2023
  - JavaScript

## Go

I'm using Go now for these challenges, starting for 2015 and working upwards.

For Go solutions, these can be ran from the homepage as follows:

```bash
go run ./2015/go/day01
```

to run with example input, or

```bash
go run ./2015/go/day01 ./input
```
to run with the input at the path given.

## Templates

### JavaScript

You can run
```bash
bash createJsDay.sh <day> <path>
```
to create the template for a new day in JavaScript. To do this, ensure that there is a `dayx` template folder in the destination.
_Tip: Add '0' to the start of the day number if it is less than 10, e.g. `01` instead of `1`, for better ordering!_

### Go

Run
```bash
bash createGoDay.sh <day> <path>
```

to create the template for a new day in Go. To do this, ensure that there is a `dayx` template folder in the destination.
_Tip: Add '0' to the start of the day number if it is less than 10, e.g. `01` instead of `1`, for better ordering!_