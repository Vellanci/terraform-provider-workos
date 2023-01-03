terraform {
  required_providers {
    workos = {
      source = "hashicorp.com/aleshchynskyi/workos"
    }
  }
}

provider "workos" {
  host = "https://api.workos.com"
}

data "workos_organizations" "example" {}

output "example_organizations" {
  value = data.workos_organizations.example
}