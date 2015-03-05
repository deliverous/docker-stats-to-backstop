// vim: ts=2 nowrap
package docker

type Container struct {
	Id    string   ``
	Image string   ``
	Names []string ``
}

func (container Container) Name() string {
	if len(container.Names) > 0 {
		return container.Names[0]
	} else {
		return container.Id
	}
}
