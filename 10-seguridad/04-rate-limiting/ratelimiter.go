// Rate limiter por IP usando un mapa con sync.Mutex.
// En producción se usa Redis para compartir el estado entre instancias del servicio.

package main

import (
	"sync"
	"time"
)

type clienteInfo struct {
	solicitudes int
	ventana     time.Time
}

type RateLimiter struct {
	mu          sync.Mutex
	clientes    map[string]*clienteInfo
	maxRequests int
	ventana     time.Duration
}

func NuevoRateLimiter(maxRequests int, ventana time.Duration) *RateLimiter {
	rl := &RateLimiter{
		clientes:    make(map[string]*clienteInfo),
		maxRequests: maxRequests,
		ventana:     ventana,
	}
	// Limpieza periódica de clientes inactivos para no acumular memoria.
	go rl.limpiar()
	return rl
}

// Permitir retorna true si el cliente puede hacer el request, false si excedió el límite.
func (rl *RateLimiter) Permitir(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	ahora := time.Now()
	info, existe := rl.clientes[ip]

	if !existe || ahora.After(info.ventana.Add(rl.ventana)) {
		// Primera solicitud o ventana expirada: reinicia el contador.
		rl.clientes[ip] = &clienteInfo{solicitudes: 1, ventana: ahora}
		return true
	}
	if info.solicitudes >= rl.maxRequests {
		return false
	}
	info.solicitudes++
	return true
}

func (rl *RateLimiter) limpiar() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		rl.mu.Lock()
		ahora := time.Now()
		for ip, info := range rl.clientes {
			if ahora.After(info.ventana.Add(rl.ventana * 2)) {
				delete(rl.clientes, ip)
			}
		}
		rl.mu.Unlock()
	}
}
