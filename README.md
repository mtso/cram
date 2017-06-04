# cram

A line cramming and un-cramming utility.

Cramming replaces all newlines with tabs, writes to the filename plus `.cram`, and deletes the given file.
Un-cramming replaces all tabs with newlines, writes to the filename removing `.cram`, and deletes the given file.

[![Build Status](https://travis-ci.org/mtso/cram.svg?branch=master)](https://travis-ci.org/mtso/cram)

## Install

```
$ go get github.com/mtso/cram
```

## Command

```
cram [filepath]
```
