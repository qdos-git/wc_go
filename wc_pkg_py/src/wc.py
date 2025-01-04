#!/usr/bin/python3

import sys
import subprocess

subprocess.call(['./wc_go.bin'] + sys.argv[1:])

