package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

func main() {
	fmt.Println("vim-go")

	_, _ := tfconfig.LoadModule(dir)
	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})
}
