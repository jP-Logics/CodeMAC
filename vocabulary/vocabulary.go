package vocabulary

// Vocabular is an interface
type Vocabular interface {
	GetKeywords(file string) (map[string]interface{}, error)
	GetOeprators(file string) (map[string]interface{}, error)
	GetPunctuations(file string) (map[string]interface{}, error)
	GetLiterals(file string) (map[string]interface{}, error)
	GetSymbols(file string) (map[string]interface{}, error)
	GetMisc(file string) (map[string]interface{}, error)
}

// Vocabulary is a type
type Vocabulary struct {
	Elements map[string]map[string]interface{}
}

// New retuns new Vocabulary
func New(filename string, vocabular Vocabular) (*Vocabulary, error) {
	var err error
	v := &Vocabulary{}
	v.Elements = make(map[string]map[string]interface{})
	v.Elements["keywords"], err = vocabular.GetKeywords(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["operators"], err = vocabular.GetOeprators(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["punctuations"], err = vocabular.GetPunctuations(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["literals"], err = vocabular.GetLiterals(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["symbols"], err = vocabular.GetSymbols(filename)
	if err != nil {
		return nil, err
	}
	v.Elements["misc"], err = vocabular.GetMisc(filename)
	if err != nil {
		return nil, err
	}

	return v, nil
}
