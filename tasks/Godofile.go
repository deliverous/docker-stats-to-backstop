// vim: ts=2 nowrap
package main

import (
	. "gopkg.in/godo.v1"
)

func tasks(p *Project) {
	p.Task("default", D{"tests"})

	p.Task("tests", func(c *Context) {
		Run("go fmt ./...")
		Run("go vet ./...")
		command := "go test ./..."
		if c.Args.ZeroBool("cover", "c") {
			command += " --cover"
		}
		Run(command)
	}).Watch("**/*.go").Debounce(1000)
}

func main() {
	Godo(tasks)
}
