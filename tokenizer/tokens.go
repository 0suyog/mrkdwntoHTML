package tokenizer

var SingleLetterTokens = map[string]TokenType{
	"_":  UNDERSCORE,
	"*":  ASTERIK,
	"\t": TAB,
	"\n": NEWLINE,
}
