package keywords

import (
	"sort"
	"strings"
)

func FormatKeywords(keywords []string) string {
	formatted := make([]string, len(keywords))
	existing := make(map[string]struct{})
	for i, kw := range keywords {
		kw = strings.ToLower(kw)
		kw = strings.TrimSpace(kw)
		if kw == "" {
			continue
		}
		if _, ok := existing[kw]; !ok { //去重
			formatted[i] = kw
			existing[kw] = struct{}{}
			formatted = append(formatted, kw)
		}
	}
	sort.Strings(formatted)
	return strings.Join(formatted, "|")

}
