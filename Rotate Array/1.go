package main

type CircleBuffer[K any] struct {
	tr    int
	tw    int
	count int
	data  []K
}

func New[K any](l int) *CircleBuffer[K] {
	return &CircleBuffer[K]{
		data: make([]K, l),
	}
}

func (a *CircleBuffer[K]) Read() (K, bool) {
	if a.count == 0 {
		var empty K
		return empty, false
	}

	v := a.data[a.tr]
	a.tr = (a.tr + 1) % len(a.data)
	a.count--

	return v, true
}

func (a *CircleBuffer[K]) Write(v K) {
	a.data[a.tw] = v
	a.tw = (a.tw + 1) % len(a.data)
	if a.count == len(a.data) {
		a.tr = (a.tr + 1) % len(a.data)
		return
	}

	a.count++
}
