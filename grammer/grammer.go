package grammer

// KeywordGrammer is a interface for keyword functionality
type KeywordGrammer interface {
	GetKeywords(file string) (map[string]interface{}, error)
	IsKeyword(string) bool
}

// OperatorGrammer is a interface for Operators functionality
type OperatorGrammer interface {
	GetOeprators(file string) (map[string]interface{}, error)
	IsOperator(string) bool
}

// PunctuationGrammer is a interface for Punctuation functionaity
type PunctuationGrammer interface {
	GetPunctuations(file string) (map[string]interface{}, error)
	IsPunctuation(string) bool
}

// LiteralGrammer is an interface for literal functionality
type LiteralGrammer interface {
	GetLiterals(file string) (map[string]interface{}, error)
	IsLiteral(string) bool
}

// SymbolGrammer is an interface for Symbol functionality
type SymbolGrammer interface {
	GetSymbols(file string) (map[string]interface{}, error)
	IsSymbol(string) bool
}

// TokenGrammer is an interface for token functionality
type TokenGrammer interface {
	GetTokens(file string) (map[string]interface{}, error)
	IsToken(string) bool
}

// MiscGrammer is an interface for misc(any missing grammer type) functionality
type MiscGrammer interface {
	GetMisc(file string) (map[string]interface{}, error)
	IsMisc(string) bool
}

// Grammer is an interface that consists of language grammers
type Grammer interface {
	KeywordGrammer
	OperatorGrammer
	PunctuationGrammer
	LiteralGrammer
	SymbolGrammer
	TokenGrammer
	MiscGrammer
}
