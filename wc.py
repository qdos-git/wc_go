#!/usr/bin/python3

##  How Python is suppose to be...

import sys
import subprocess

subprocess.run(["go", "run", "wc.go", sys.argv[1]])

##  Note: above untested.

##  How Python really is...

exec(go run wc.go)
