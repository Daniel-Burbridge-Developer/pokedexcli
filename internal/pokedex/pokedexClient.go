package pokedex

type Pokedex struct {
	entries map[string]Pokemon
}

func New() *Pokedex {
	pd := Pokedex{}
	pd.entries = make(map[string]Pokemon)
	return &pd
}

func (pd *Pokedex) Add(url string, pokemon Pokemon) {
	pd.entries[url] = pokemon
}

func (pd *Pokedex) Get(key string) (*Pokemon, bool) {
	entry, ok := pd.entries[key]
	if !ok {
		return nil, false
	}

	return &entry, true
}
