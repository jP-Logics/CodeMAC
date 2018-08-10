package vocabulary

import "CodeTeller/grammer"

// Vocabulary is a type
type Vocabulary struct {
	Elements map[string]map[string]interface{}
}

// New retuns new Vocabulary
func New(filename string, grammer grammer.Grammer) (*Vocabulary, error) {
	var err error
	v := &Vocabulary{}
	v.Elements = make(map[string]map[string]interface{})
	v.Elements["keywords"], err = grammer.GetKeywords(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["operators"], err = grammer.GetOeprators(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["punctuations"], err = grammer.GetPunctuations(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["literals"], err = grammer.GetLiterals(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["symbols"], err = grammer.GetSymbols(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["misc"], err = grammer.GetMisc(filename)
	if err != nil {
		return nil, err
	}

	return v, nil
}
