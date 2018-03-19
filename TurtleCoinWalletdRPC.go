package TurtleCoinWalletdRPC

import (
  "bufio"
  "bytes"
  "encoding/json"
  "errors"
  "io"
  "io/ioutil"
  "math/rand"
  "net/http"
  "os"
  "os/exec"
  "path/filepath"
  "strconv"
  "strings"
  "syscall"
  "time"

  log "github.com/sirupsen/logrus"
)

type RPCPayload struct {
  JSONRPC  string                  `json:"jsonrpc"`
  Method   string                  `json:"method"`
  Params   *map[string]interface{} `json:"params,omitempty"`
  Password string                  `json:"password"`
  ID       int                     `json:"id"`
}

func buildRPC(
  method       string,
  id           int,
  rpcPassword  string,
  params       RPCPayload
)
RPCPayload {
  return RPCPayload{
    JSONRPC   : "2.0",
    Method    : method,
    Password  : password,
    ID        : id,
    params    : params
  }
}

func reset(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "reset",
    id,
    rpcPassword,
    params
  )
}

func save(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "save",
    id,
    rpcPassword,
    params
  )
}

func save(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "save",
    id,
    rpcPassword,
    params
  )
}

func getViewKey(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "getViewKey",
    id,
    rpcPassword,
    params
  )
}

func getSpendKeys(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "getSpendKeys",
    id,
    rpcPassword,
    params
  )
}

func getStatus(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "getStatus",
    id,
    rpcPassword,
    params
  )
}

func getAddresses(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "getAddresses",
    id,
    rpcPassword,
    params
  )
}

func createAddress(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "createAddress",
    id,
    rpcPassword,
    params
  )
}

func deleteAddress(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "deleteAddress",
    id,
    rpcPassword,
    params
  )
}

func getBalance(
  id          int,
  rpcPassword string,
  params      RPCPayload
)
RPCPayload {
  return buildRPC(
    "getBalance",
    id,
    rpcPassword,
    params
  )
}
