package main

type RotateArray[K any] struct {
	tr    int
	tw    int
	count int
	data  []K
}

func New[K any](l int) *RotateArray[K] {
	return &RotateArray[K]{
		data: make([]K, l),
	}
}

func (a *RotateArray[K]) Read() (K, bool) {
	if a.count == 0 {
		var empty K
		return empty, false
	}

	v := a.data[a.tr]
	a.tr = (a.tr + 1) % len(a.data)
	a.count--

	return v, true
}

func (a *RotateArray[K]) Write(v K) {
	a.data[a.tw] = v
	a.tw = (a.tw + 1) % len(a.data)
	if a.count == len(a.data) {
		a.tr = (a.tr + 1) % len(a.data)
		return
	}

	a.count++
}
