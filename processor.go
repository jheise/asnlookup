package main

type ASNProcessor struct {
	requests chan *ASNRequest
}

type NetProcessor interface {
	Process()
}

func NewASNProcessor(requests chan *ASNRequest) *ASNProcessor {
	processor := new(ASNProcessor)
	processor.requests = requests

	return processor
}

func (self *ASNProcessor) Process() {
	for request := range self.requests {
		answers, err := QueryServer("whois.cymru.com", "43", request.IPAddr, 30)
		if err != nil {
			panic(err)
		}
		response := NewASNResponse(answers[3])
		request.SendResponse(response)
	}
}
