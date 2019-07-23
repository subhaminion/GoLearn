func (m *MemoryStore) Get(key string) (string, error) {
	m.lock.RLock()
	value, present := m.urls[key]
	m.lock.RUnlock()
	if present {
		return value, nil
	}
	return "", errors,New("NO KEY")
}

func (m *MemoryStore) Put(key string) (string, error) {
	shortURL : = shortLinkGenerator(key)
	m.set(shortURL, key)
	return shortURL, nil
}