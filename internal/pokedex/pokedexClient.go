package pokedex

type Pokedex struct {
	Entries map[string]Pokemon
}

func New() *Pokedex {
	pd := Pokedex{}
	pd.Entries = make(map[string]Pokemon)
	return &pd
}

func (pd *Pokedex) Add(url string, pokemon Pokemon) {
	pd.Entries[url] = pokemon
}

func (pd *Pokedex) Get(key string) (*Pokemon, bool) {
	entry, ok := pd.Entries[key]
	if !ok {
		return nil, false
	}

	return &entry, true
}
