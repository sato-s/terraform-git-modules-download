package main

import (
	"fmt"

	"github.com/go-git/go-git/v5"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func main() {
	fmt.Println("vim-go")

	dir := "test/sample-terraform/"
	modules, dials := tfconfig.LoadModule(dir)
	fmt.Println(modules)
	fmt.Println(dials)
	fmt.Println(git.PlainOpen)
	//_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
	//	URL:      "https://github.com/go-git/go-git",
	//	Progress: os.Stdout,
	//})
}
