resource "random_string" "id" {
  length           = 8
  special          = true
  min_lower        = 8
}