package helpers

import "net/http"

func HxRe(w http.ResponseWriter, target string, swap string) {
	if target != "" {
		w.Header().Set("HX-Retarget", target)
	}

	if swap != "" {
		w.Header().Set("HX-Reswap", target)
	}
}
