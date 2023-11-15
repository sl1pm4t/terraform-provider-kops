terraform {
  required_providers {
    kops = {
      source  = "github/sl1pm4t/kops"
      version = "0.0.1"
    }
  }
}

provider "kops" {
  state_store = "file://./store/"
  mock        = true
  aws {
    region = "us-test-1"
  }
}
