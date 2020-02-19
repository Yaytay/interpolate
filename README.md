#  interpolate

The interpolate tool takes in a matrix of floating point values as a CSV and outputs the same matrix with any unknown values
replaced by interpolating from non-diagonal neighbouring values.

All rows in the input matrix are expected to have the same number of fields.
If rows are shorter than the first row an error message will be printed, but processing will continue as if the missing fields were all nan.
If rows are longer than the first row an error message will be printed, but processing will continue and the excess fields will be dropped.
Any text value other than "nan" will cause an error message to be printed, but processing will continue as if the input was "nan".

The interpolate tool is written to handle relatively large files - as long as four lines can be held in memory it should work with any size file.

# Usage
```
interpolate [OPTIONS] <inputfile> <outputfile>
 inputfile may be "-" to input from stdin
 outputfile may be "-" to output to stderr
 options:
  -?    Print this message
```
# Compilation

The interpolate tool was written using go 1.13 on linux, though it should build and run on Windows too.

At this stage no build manager has been created, so the following commands should be used to compile:
* cd $GOPATH
* mkdir -p src/github.com/yaytay
* cd src/github.com/yaytay
* git clone https://github.com/Yaytay/interpolate.git
* cd $GOPATH
* git test github.com/Yaytay/interpolate
* git install github.com/Yaytay/interpolate
