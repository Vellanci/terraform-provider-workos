resource "workos_organization" "example" {
  name    = "Provided Org by Terraform"
  domains = ["provided-org.org", "provided-org.ua"]
}

output "example_organizations" {
  value = workos_organization.example
}
