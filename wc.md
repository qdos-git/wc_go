A mini wc spec
- There are no flag options.
- The output is linecount wordcount bytecount filepath (each separated by tab characters)
- There is exactly one filepath passed as an argument.
- That path argument is not stdin (i.e., "-")
- It should handle errors as wc for things inside the spec. I.e., if wc would give an error for a one argument invocation, you should give the same error.
- no imports/stdlib except arg parsing

todo research
  edge cases:
    - wc without an arg
    - empty file
    - 1 newline file
  return numbers?
  


to write
    parse argument - ?
    


joke python script
    calls go binary


##  Best, most backwards compatible way (and forwards, due to static compilation), to fix all ongoing Python problems

import sys
import subprocess

subprocess.run(["go", "run", "wc", sys.argv[1]])

##  Note: above untested, for obvious reasons.

args
  -c -m -l -w 
    just filter the wc output
  -L
    longest line

import (
    "flag"
    "fmt"
)

add test suite
