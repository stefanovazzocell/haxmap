package benchmark

import (
	"testing"
)

func setupGoMap() map[uintptr]uintptr {
	m := make(map[uintptr]uintptr, epochs)
	for i := uintptr(0); i < epochs; i++ {
		m[i] = i
	}
	return m
}

func BenchmarkSyncHaxMapReadsOnly(b *testing.B) {
	m := setupHaxMap()
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		j, _ := m.Get(i % epochs)
		if j != (i % epochs) {
			b.Fail()
		}
	}
}

func BenchmarkSyncHaxMapReadsWithWrites(b *testing.B) {
	m := setupHaxMap()
	var writer bool
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		// use 1 thread as writer
		if writer {
			m.Set(i%epochs, i%epochs)
		} else {
			j, _ := m.Get(i % epochs)
			if j != (i % epochs) {
				b.Fail()
			}
		}
		writer = !writer
	}
}

func BenchmarkSyncGoSyncMapReadsOnly(b *testing.B) {
	m := setupGoSyncMap()
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		j, _ := m.Load(i % epochs)
		if j != (i % epochs) {
			b.Fail()
		}
	}
}

func BenchmarkSyncGoSyncMapReadsWithWrites(b *testing.B) {
	m := setupGoSyncMap()
	var writer bool
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		// use 1 thread as writer
		if writer {
			m.Store(i%epochs, i%epochs)
		} else {
			j, _ := m.Load(i % epochs)
			if j != (i % epochs) {
				b.Fail()
			}
		}
		writer = !writer
	}
}

func BenchmarkSyncCornelkMapReadsOnly(b *testing.B) {
	m := setupCornelkMap()
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		j, _ := m.Get(i % epochs)
		if j != (i % epochs) {
			b.Fail()
		}
	}
}

func BenchmarkSyncCornelkMapReadsWithWrites(b *testing.B) {
	m := setupCornelkMap()
	var writer bool
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		// use 1 thread as writer
		if writer {
			m.Set(i%epochs, i%epochs)
		} else {
			j, _ := m.Get(i % epochs)
			if j != (i % epochs) {
				b.Fail()
			}
		}
		writer = !writer
	}
}

func BenchmarkSyncXsyncMapReadsOnly(b *testing.B) {
	m := setupXsyncMap()
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		j, _ := m.Load(i % epochs)
		if j != (i % epochs) {
			b.Fail()
		}
	}
}

func BenchmarkSyncXsyncMapReadsWithWrites(b *testing.B) {
	m := setupXsyncMap()
	var writer bool
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		// use 1 thread as writer
		if writer {
			m.Store(i%epochs, i%epochs)
		} else {
			j, _ := m.Load(i % epochs)
			if j != (i % epochs) {
				b.Fail()
			}
		}
		writer = !writer
	}
}

func BenchmarkSyncGoMapRWMutexReadsOnly(b *testing.B) {
	m := setupGoMapRWMutex()
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		m.l.RLock()
		j := m.m[i%epochs]
		m.l.RUnlock()
		if j != (i % epochs) {
			b.Fail()
		}
	}
}

func BenchmarkSyncGoMapRWMutexReadsWithWrites(b *testing.B) {
	m := setupGoMapRWMutex()
	var writer bool
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		// use 1 thread as writer
		if writer {
			m.l.Lock()
			m.m[i%epochs] = i % epochs
			m.l.Unlock()
		} else {
			m.l.RLock()
			j := m.m[i%epochs]
			m.l.RUnlock()
			if j != (i % epochs) {
				b.Fail()
			}
		}
		writer = !writer
	}
}

func BenchmarkSyncGoMapReadsOnly(b *testing.B) {
	m := setupGoMap()
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		j := m[i%epochs]
		if j != (i % epochs) {
			b.Fail()
		}
	}
}

func BenchmarkSyncGoMapReadsWithWrites(b *testing.B) {
	m := setupGoMap()
	var writer bool
	b.ResetTimer()
	for i := uintptr(0); i < uintptr(b.N); i++ {
		// use 1 thread as writer
		if writer {
			m[i%epochs] = i % epochs
		} else {
			j := m[i%epochs]
			if j != (i % epochs) {
				b.Fail()
			}
		}
		writer = !writer
	}
}
