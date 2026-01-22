# Quadchecker

## Overview

**quadchecker** is a Go program that identifies which `quad` function(s) generated a given ASCII rectangle.
It analyzes a string input representing a rectangle and determines the matching quad type(s) along with their dimensions.

If the input does not correspond to any valid quad function, the program outputs:

```
Not a quad function
```

## Project Structure

The **main** branch contains only the files required for submission:

```
quadchecker (main branch)
├── go.mod
└── main.go
```

Other branches contain additional **quad executables** to allow testing the program:

```
quadchecker (boookumu / dosure / piregi branches)
├── go.mod
├── main.go
├── quadA
├── quadB
├── quadC
├── quadD
├── quadE
├── quadchecker
└── README.md
```

## Branches

* `main` – Contains only the files required for submission.
* `boookumu`, `dosure`, `piregi` – Include `quad` executables for testing purposes.

## Rules & Requirements

* The program takes **a string as input** (usually piped from another program).
* It prints the name(s) of the matching quad function(s) and the rectangle dimensions.
* All outputs **must end with a newline (`\n`)**.
* If **multiple quad functions match**, they must be:

  * Displayed **alphabetically**
  * Separated by `||`
* If **no quad matches**, print:

  ```
  Not a quad function
  ```

## Output Format

```
[quadX] [width] [height]
```

Example with multiple matches:

```
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

## Usage

### Example: quadA (3 × 3)

```
$ ./quadA 3 3 | go run .
[quadA] [3] [3]
```

### Example: quadC / quadD / quadE (1 × 1)

```
$ ./quadC 1 1 | go run .
[quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]
```

### Example: quadC / quadE (1 × 2)

```
$ ./quadE 1 2 | go run .
[quadC] [1] [2] || [quadE] [1] [2]
```

### Invalid Quad Input

```
$ echo 0 0 | go run .
Not a quad function
```

```
$ echo -n "o--o"$'\n'"|"$'\n'"o" | go run .
Not a quad function
```

## Notes

* The program reads from **standard input (stdin)**.
* Dimensions are inferred from the structure of the input string.
* Designed to work seamlessly with existing `quadA`, `quadB`, `quadC`, `quadD`, and `quadE` generators.
* For testing, switch to branches with quad executables (`boookumu`, `dosure`, `piregi`).

## Author

Project completed as part of a Go programming raid exercise.
