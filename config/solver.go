package config

// this is the address of the user
const X_COOPHIVE_USER_HEADER = "X-Lilypad-User"

// this is the signature of the message
const X_COOPHIVE_SIGNATURE_HEADER = "X-Lilypad-Signature"

// the context name we keep the address
const CONTEXT_ADDRESS = "address"

// the sub path any API's are served over
const API_SUB_PATH = "/api/v1"

// the sub path the websocket server is mounted on
const WEBSOCKET_SUB_PATH = "/ws"
