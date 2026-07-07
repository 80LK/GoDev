package patches

import (
	"strconv"
	"strings"

	"github.com/80LK/godev/internal/pipeline"
	"golang.org/x/sync/errgroup"
)

type ParallelPatch struct {
	Items []pipeline.Patch
}

func (p ParallelPatch) Apply() error {
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
	if len(p.Items) == 0 {
		return ""
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
