package main

type ASNResponse struct {
	Answer string
}

func NewASNResponse(answer string) *ASNResponse {
	resp := new(ASNResponse)
	resp.Answer = answer

	return resp
}
