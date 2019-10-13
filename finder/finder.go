package finder

import (
	"github.com/37108/brandy/config"
	"github.com/ktr0731/go-fuzzyfinder"
)

func Finder(profiles []*config.Profile) (int, error) {
	idx, err := fuzzyfinder.Find(profiles, func(i int) string {
		return profiles[i].Name
	})
	if err != nil {
		return -1, err
	}
	return idx, nil
}
