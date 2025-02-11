package registry

// ServiceRegistry menyimpan semua service yang digunakan dalam aplikasi
type ServiceRegistry struct {
	services map[string]interface{}
}

// NewServiceRegistry membuat instance ServiceRegistry baru
func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services: make(map[string]interface{}),
	}
}

// RegisterService mendaftarkan service ke registry
func (r *ServiceRegistry) RegisterService(name string, service interface{}) {
	r.services[name] = service
}

// GetService mengambil service berdasarkan nama
func (r *ServiceRegistry) GetService(name string) (interface{}, bool) {
	service, exists := r.services[name]
	return service, exists
}
