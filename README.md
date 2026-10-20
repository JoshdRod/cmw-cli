A CLI tool for Concise Message Wrapper.

This tool can:
1. Create CMW collections from a set of individual files, with files containing Evidence CBOR.
2. Append to existing CMW collections, by passing in file(s) as input.

Currently in a prototype stage.
TODO LIST:
☑️ Serialise a CMW with basic evidence
- Serialise from both JSON and CBOR collections
- Output to both JSON and CBOR formats
- Seralise multiple files into the same CMW

CURRENT PROGRESS:
Currently, cmwcli can generate a CMW from a JSON Collection, using the command `collection`. To run:

`go run main.go collection --data collection.json --output out.json`

Currently can only process and output json files.
Currently cannot handle wrapping multiple collections.
