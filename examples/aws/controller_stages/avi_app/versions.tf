terraform {
  required_version = ">= 0.13"
  required_providers {
    avi = {
      source = "vmware/avi"
      version = "20.1.4"
    }
  }
}
