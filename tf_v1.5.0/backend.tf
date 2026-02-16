terraform {
  backend "http" {
    address        = "http://localhost:8000/client/v7ra2iphsfe/state"
    lock_address   = "http://localhost:8000/client/v7ra2iphsfe/lock"
    unlock_address = "http://localhost:8000/client/v7ra2iphsfe/unlock"
    lock_method    = "POST"
    unlock_method  = "POST"
  }
}
