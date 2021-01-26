package text

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

type baseConfig struct {
	wall, in, out string
}

func newBaseConfig() *baseConfig {
	return &baseConfig{"#", "i", "E"}
}

func (c *baseConfig) configure(config map[string]string) {
	if v, ok := config["wall"]; ok {
		validateOptionValue("wall", v)
		c.wall = v
		delete(config, "wall")
	}
	if v, ok := config["in"]; ok {
		validateOptionValue("in", v)
		c.in = v
		delete(config, "in")
	}
	if v, ok := config["out"]; ok {
		validateOptionValue("out", v)
		c.out = v
		delete(config, "out")
	}

	for k := range config {
		panic(fmt.Sprintf("Unknown option %q", k))
	}
}

func validateOptionValue(name, value string) {
	if "" == value {
		panic(fmt.Sprintf("Value for %q must not be empty", name))
	}
}

func repeatStringToLength(str string, length int) string {
	return string(
		[]rune(
			strings.Repeat(
				str,
				int(math.Ceil(
					float64(length)/float64(utf8.RuneCountInString(str)),
				)),
			),
		)[0:length],
	)
}
