terraform_deploy "app" {
  source {
    owner    = "terrafire"
    repo     = "terraform"
    path     = "app/"
    revision = "xxxx"
  }

  workspace     = "dev"
  allow_destroy = true

  vars = {
    "foo_revision"          = "xxx"
    # "bar_revision"      = resolve_github_revision("terrafire-dev", "bar", "branch-do-something") # user-defined function
  }

  var_files = [
    "./app/variables.tfvars",
    "./app/secrets.tfvars.enc" # 末尾が enc だったら sops で復号
  ]
}