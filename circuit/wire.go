package circuit

type Wire struct {
	signal  bool
	actions []func()
}

func NewWire() *Wire {
	return &Wire{actions: make([]func(), 0)}
}

func (w *Wire) GetSignal() bool {
	return w.signal
}

func (w *Wire) SetSignal(newSig bool) {
	if w.signal != newSig {
		w.signal = newSig
		for _, handler := range w.actions {
			handler()
		}
	}
}

func (w *Wire) AcceptAction(action func()) {
	w.actions = append(w.actions, action)
}
