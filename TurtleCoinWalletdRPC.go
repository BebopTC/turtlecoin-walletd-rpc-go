package TurtleCoinWalletd

import(
  "./RPCBuilders"
  "net/http"
  "bytes"
  "encoding/json"

)

func sendRPCRequest(payload RPCPayload) {
  payloadJson, err := json.Marshall(payload)

  req, err := http.NewRequest("POST", "http://localhost:8080", bytes.NewBuffer(payloadjson))
}

//func reset(
//  viewSecretKey string
//)
