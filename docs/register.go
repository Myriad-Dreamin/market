package docs

var rel map[string]string

func Register(k, v string) {
	if rel == nil {
		rel = make(map[string]string)
	}
	rel[k] = v
}

func Get(k string) string {
	if rel == nil {
		return ""
	}
	return rel[k]
}


