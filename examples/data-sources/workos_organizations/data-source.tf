data "workos_organizations" "example" {}

output "example_organizations" {
  value = data.workos_organizations.example
}
