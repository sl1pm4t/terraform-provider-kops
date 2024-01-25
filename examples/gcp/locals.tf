locals {
  controlPlaneType = "e2-standard-2"
  nodeType         = "e2-standard-2"
  clusterName      = "cluster.example.com"
  dnsZone          = "example.com"
  network          = "default"
  region           = "us-central1"
  zones = [
    "us-central1-a",
    "us-central1-b",
    "us-central1-c",
  ]

}
