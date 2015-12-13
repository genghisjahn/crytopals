package main

const tkeys = "abcdefghjklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type keylengthscore struct {
	Length int
	Score  float64
}

type singlekeyscore struct {
	RawText   string
	Key       string
	Score     int
	Decrypted string
}

type klscores []keylengthscore
type skscores []singlekeyscore

func (slice klscores) Len() int {
	return len(slice)
}

func (slice klscores) Less(i, j int) bool {
	return slice[i].Score < slice[j].Score
}

func (slice klscores) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (slice skscores) Len() int {
	return len(slice)
}

func (slice skscores) Less(i, j int) bool {
	return slice[i].Score < slice[j].Score
}

func (slice skscores) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
