A CLI tool for Concise Message Wrapper.

This tool can:
1. Create CMW collections from a set of individual files, with files containing Evidence CBOR.
2. Append to existing CMW collections, by passing in file(s) as input.

Currently in a prototype stage.
TODO LIST:

- ☑️ Serialise a CMW with  a basic evidence CBOR
- ☑️ Seralise multiple files into the same CMW
- Append CBOR evidence(s) to an existing CMW

CURRENT PROGRESS:
Currently, cmwcli can generate a CMW from a CBOR Collection, using the command `collection`. To run:

`go run main.go collection --data file.cbor [--data file2.cbor ...] --output out.cbor [--cmwc_t "URI/OID" (the name of the collection that wraps the multiple cbor files input)]`

Currently can only process and output CBOR files. (Should this be expanded to JSON and others as well?)
