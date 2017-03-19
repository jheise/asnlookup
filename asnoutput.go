package main

type ASNOutput struct {
	Address string
	ASN     string
}

func NewASNOutput(addr string, asn string) *ASNOutput {
	output := new(ASNOutput)
	output.Address = addr
	output.ASN = asn

	return output
}
