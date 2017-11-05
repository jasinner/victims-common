# Victims-Common
[![Build Status](https://travis-ci.org/victims/victims-common.svg)](https://travis-ci.org/victims/victims-common/)

Common libraries for Victims-API and Victims-Upload

These projects were split so that Victims-API could provide a read-only view of the database while Victims-Upload is used for modifying data. Victims-Upload will not be exposed externally.

## Building

To make sure the library builds:

```
$ make build
```

Other targets can be found with:

```
$ make help
```
