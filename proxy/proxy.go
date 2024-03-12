// proxy.go
package proxy

type ProxyManager struct {
	Proxies []string
}

func NewProxyManager(proxies []string) *ProxyManager {
	return &ProxyManager{
		Proxies: proxies,
	}
}

// GetAllProxies returns a slice of all proxy URLs.
func (pm *ProxyManager) GetAllProxies() []string {
	return pm.Proxies
}
