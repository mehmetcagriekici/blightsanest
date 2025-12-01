package pubsub

import(
        "sync"
)

// holds cancel functions for all consumers
type SubscriptionManager struct {
        mu         sync.Mutex
	cancellers []func() error
	closed     bool
}

// create a new subscription manager
func NewSubscriptionManager() *SubscriptionManager {
        return &SubscriptionManager{
	        cancellers: []func() error{},
	}
}

// store a cancel function so it can be closed later
func (m *SubscriptionManager) Add(cancel func() error) {
        m.mu.Lock()
	defer m.mu.Unlock()
	
	if m.closed {
	        cancel()
		return
	}
	
	m.cancellers = append(m.cancellers, cancel)
}

// cancel every subscription, prevent new additions
func (m *SubscriptionManager) CloseAll() error {
        m.mu.Lock()
	defer m.mu.Unlock()

        // already closed
        if m.closed {
	        return nil
	}

        m.closed = true

        var lastError error
	for _, cancel := range m.cancellers {
	        if err := cancel(); err != nil {
		      lastError = err
		}
	}
	m.cancellers = nil
	return lastError
}