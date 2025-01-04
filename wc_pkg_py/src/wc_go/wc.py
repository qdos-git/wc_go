#!/usr/bin/python3

import sys
import subprocess
import site


def main():

    full_path = site.getsitepackages()[0]

    full_path += "/wc_go/bin/wc_go.bin"
    
    subprocess.call([full_path] + sys.argv[1:])

    
if __name__ == "__main__":
    main()
