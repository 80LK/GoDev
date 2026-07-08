package patches

import (
	"strconv"
	"strings"

	"golang.org/x/sync/errgroup"
)

type ParallelPatch struct {
	Items []Patch
}

func (p ParallelPatch) Apply() error {
	l := len(p.Items)
	switch l {
	case 0:
		return nil
	case 1:
		return p.Items[0].Apply()
	}

	var g errgroup.Group

	for _, item := range p.Items {
		g.Go(func() error {
			return item.Apply()
		})
	}

	return g.Wait()
}

func (p ParallelPatch) Diff() (string, error) {
	return "", nil
}

func (p ParallelPatch) Summary() string {
	l := len(p.Items)
	switch l {
	case 0:
		return ""
	case 1:
		return p.Items[0].Summary()
	}

	var str strings.Builder

	str.WriteString("Parallel:")
	i := 1
	for _, item := range p.Items {
		sum := item.Summary()
		if sum == "" {
			continue
		}

		str.WriteString("\n\t" + strconv.Itoa(i) + ". " + item.Summary())
		i++
	}

	return str.String()
}
