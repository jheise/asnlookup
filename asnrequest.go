package main

type ASNRequest struct {
	IPAddr       string
	ResponseChan chan *ASNResponse
}

func NewASNRequest(ipaddr string) *ASNRequest {
	request := new(ASNRequest)
	request.IPAddr = ipaddr
	request.ResponseChan = make(chan *ASNResponse)

	return request
}

func (self *ASNRequest) SendResponse(response *ASNResponse) {
	self.ResponseChan <- response
	close(self.ResponseChan)
}
