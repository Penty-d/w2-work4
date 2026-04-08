package keywords

import (
	"sort"
	"strings"
)

func FormatKeywords(keywords []string) string {
	formatted := make([]string, 0, len(keywords))
	existing := make(map[string]struct{})
	for _, kw := range keywords {
		kw = strings.ToLower(kw)
		kw = strings.TrimSpace(kw)
		if kw == "" {
			continue
		}
		if _, ok := existing[kw]; !ok { //去重
			formatted = append(formatted, kw)
			existing[kw] = struct{}{}
		}
	}
	sort.Strings(formatted)
	return strings.Join(formatted, "|")

}
