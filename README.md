
# Intro

This repository is an initial attempt at a backwards-compatible port
of BSD `wc`, but written in a modern compiled language (Go), for
purposes of memory safety. 


# Features implemented

- Takes one file as an argument

- Takes arguments regarding bytes, lines, characters, and words, in
  any order, as long as they are combined into one argument, for
  example "-clmw"

- Maintains consistent output formatting regardless of order of
  arguments


#  Features not implemented

The following features are a part of the BSD wc implementation, but
are not found in this implementation, due to time constraints (and
funding).

- Pad all numerical output to length of largest number

- Ensure like-for-like compatibility regarding what constitutes a word
  (may require reading source code)

- Implement spacing between flags for arguments

- Interactivity when no arguments given

- Handle UTF-16 data equivalently

- Return codes

Note that this list is non-exhaustive, and is likely the reason people
are still using a C implementation of wc, for better or worse...


# Test data

Tests have been written for the following input data

- Empty file

- File with non-printable characters (control codes)

- Standard Lorem Ipsum placeholder text


# Tests not implemented

- Variety of permutations of arguments

- UTF-16 text


# Python packaging

To demonstrate that their is a better way, for the Python community, a
basic Python wrapper has been written to call this wc Go
implementation as a binary. This could be useful in the case of
distributing a multiarch binary.

A CLI callable from any directory has then been setup, via a carefully
crafted pyproject.toml file.

See python_packaging.md for more info.


# Original spec

An initial mini wc spec:

- There are no flag options.

- The output is linecount wordcount bytecount filepath (each separated
  by tab characters)

- There is exactly one filepath passed as an argument.

- That path argument is not stdin (i.e., "-")

- It should handle errors as wc for things inside the spec. I.e., if
wc would give an error for a one argument invocation, you should give
the same error.

- no imports/stdlib except arg parsing
