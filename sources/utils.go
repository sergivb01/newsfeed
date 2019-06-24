package sources

func cutText(text string) string {
	cut := len(text)
	if len(text) > 150 {
		cut = 150
	}

	return text[:cut]
}
