

Allegedly the package must follow this structure:

package_name/
├── LICENSE
├── pyproject.toml
├── README.md
├── src/
│   └── package_name/
│       ├── __init__.py
│       └── example.py
└── tests/


The build backend converts source files into a Wheel package.

A build frontend takes a source distribution and builds a built
distribution, by calling build backends that have been explicitly
associated with each source tree.

Wheel (.whl) is a ZIP-format with a specific file name, based on
various PEP standards, and is what the end-user will download and
install. The files will only need to be moved to a location on the
target system; this is a built distribution.

By carefully configuring pyproject.toml, it is possible to then call
the Python script that calls the Go binary, from anywhere.