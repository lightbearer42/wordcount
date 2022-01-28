package processor

import (
    "strings"
)

func Count(s string) int {
    return len(strings.Fields(s))
}
