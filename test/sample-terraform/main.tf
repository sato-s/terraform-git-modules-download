resource "null_resource" "fake" {
}

module "hoge" {
  source = "git@github.com:sato-s/tfaction-playground//modules/hoge?ref=common-module"
}

module "fuga" {
  source = "git@github.com:sato-s/tfaction-playground//modules/hoge?ref=common-module"

}

module "fumu" {
  source = "git@github.com:sato-s/tfaction-playground//modules/hoge?ref=common-module"
}


