# Line Art Encoder/Decoder

## Description
This is a Go program that encodes and decodes line art. It provides various options including encoding multiple lines of input, reading input from a file, writing output to a file, and adding color to the output.

## Usage
To run the program, use the following command:
bash go run . [options] [input]

Options include:
- `-m`: Enable to enter multiple lines of input
- `-e`: Select encoding mode for input line
- `-o`: Write output to specified file
- `-i`: Read input from a file
- `-io`: Read input from file & write this to output file
- `-c`: Select color for stdout

### Encode a single line  
go run . -e "[5 #][5 -_]-[5 #]"

### Encode multiple lines from the console  
go run . -m -e

### Decode multiple lines from a file and write to an output file  
go run . -io input.txt output.txt  

## Features

**Balanced Brackets Check**: Ensures the input has balanced square brackets.  
**Colorful Output**: Optionally adds color to the decoded ASCII art for a visually appealing display.  
**File Input/Output**: Supports reading input from files and writing output to files.  

## Examples

### Encoding

Input: **[5 #][5 -_]-[5 #]**

Output: **#####-_-_-_-_-_-#####**

### Decoding

Input: **#####-_-_-_-_-_-#####**

Output: **[5 #][5 -_]-[5 #]**
