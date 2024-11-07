#!/usr/bin/python3

import sys
import subprocess

subprocess.run(["go", "run", "wc", sys.argv[1]])

##  Note: above untested.
