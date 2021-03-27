resource "sentry_project" "project" {
  organization = var.organization
  team         = var.team
  name         = var.name
  slug         = var.slug
  platform     = var.platform
  resolve_age  = var.resolve_age
}
