// vim: ts=2 nowrap
package main

import (
	"fmt"
	. "gopkg.in/godo.v1"
	"time"
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
		now := time.Now()
		fmt.Printf("Last run finished at %d-%02d-%02d %02d:%02d:%02d\n",
			now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second())
	}).Watch("**/*.go").Debounce(1000)

	p.Task("run", D{"tests"}, func() {
		Start("server.go", In{"server"})
	}).Watch("**/*.go").Debounce(1000)
}

func main() {
	Godo(tasks)
}
