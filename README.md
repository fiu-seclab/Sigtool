# sigtool
This a sigtool in Go for signed PE files.
Currently, only extracting digital signatures embeded in a PKCS#7 `SignedData` in a signed PE is supported.
Adding and deleting digital signatures will be supported soon.

Collects and store PKCS7 file of extracted certs in database => format: hex of byte array

## Command Line Usage
Example sigtool CLI usage:

	gosigtool 
