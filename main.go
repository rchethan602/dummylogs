package main

import (
	"fmt"
	"time"
)

func main() {

	for {
		fmt.Println("[2024-12-03T06:38:45.111Z] \"GET /portal HTTP/1.1\" 200 - 0 1111 5 5 \"10.47.1.68\" \"Amazon-Route53-Health-Check-Service (ref 1bd07fd9-450e-4420-a0f0-383a797282bb; report http://amzn.to/1vsZADi)\" \"364d24ae-63df-904e-aa9c-96777b10a6ea\" \"keycore-yosemite-https.ingress.k8s-prod.us-east-1.prod.cdsf.io\" \"10.20.102.211:80\"")
		time.Sleep(2 * time.Second)
		fmt.Println("[2024-12-02T14:38:38.474Z] \"POST /1/syncgateway/ecumessages HTTP/1.0\" 200 - 501 19 427 30 \"10.47.2.81\" \"ACCMba8fb7c283e3f49d\" \"e3e58c98-b366-9f21-aacd-5cd5467ff804\" \"keycore-yosemite-http.ingress.k8s-prod.us-east-1.prod.cdsf.io\" \"10.20.103.89:80\"")
		time.Sleep(1 * time.Millisecond)
		fmt.Println("[2024-12-02T14:39:38.474Z] \"GET /1/virtualkeys HTTP/1.0\" 503 - 501 19 427 30 \"10.47.2.81\" \"ACCMba8fb7c28321f49d\" \"e3e58c98-b366-9f87-aacd-5cd5467ff804\" \"keycore-yosemite-http.ingress.k8s-prod.us-east-1.prod.cdsf.io\" \"10.20.103.89:80\"")
		time.Sleep(3 * time.Second)
		fmt.Println("[2024-12-02T14:40:38.474Z] \"PUT /1/user HTTP/1.0\" 409 - 501 19 427 30 \"10.47.2.81\" \"ACCMba8fb7c28321f49d\" \"e3e58c98-b366-9f87-aacd-5cd5467ff804\" \"keycore-yosemite-http.ingress.k8s-prod.us-east-1.prod.cdsf.io\" \"10.20.103.89:80\"")
	}
}
