package streamdeck

// An ApplicationDidLaunchHandler responds to ApplicationDidLaunchEvents.
type ApplicationDidLaunchHandler interface {
	ApplicationDidLaunch(*ApplicationDidLaunchEvent)
}

// An ApplicationDidLaunchHandlerFunc responds to ApplicationDidLaunchEvents.
type ApplicationDidLaunchHandlerFunc func(*ApplicationDidLaunchEvent)

// ApplicationDidLaunch calls f(e).
func (f ApplicationDidLaunchHandlerFunc) ApplicationDidLaunch(e *ApplicationDidLaunchEvent) {
	f(e)
}

// An ApplicationDidTerminateHandler responds to ApplicationDidTerminateEvents.
type ApplicationDidTerminateHandler interface {
	ApplicationDidTerminate(*ApplicationDidTerminateEvent)
}

// An ApplicationDidTerminateHandlerFunc responds to ApplicationDidTerminateEvents.
type ApplicationDidTerminateHandlerFunc func(*ApplicationDidTerminateEvent)

// ApplicationDidTerminate calls f(e).
func (f ApplicationDidTerminateHandlerFunc) ApplicationDidTerminate(e *ApplicationDidTerminateEvent) {
	f(e)
}

// An DeviceDidConnectHandler responds to DeviceDidConnectEvents.
type DeviceDidConnectHandler interface {
	DeviceDidConnect(*DeviceDidConnectEvent)
}

// A DeviceDidConnectHandlerFunc responds to DeviceDidConnectEvents.
type DeviceDidConnectHandlerFunc func(*DeviceDidConnectEvent)

// DeviceDidConnect calls f(e)
func (f DeviceDidConnectHandlerFunc) DeviceDidConnect(e *DeviceDidConnectEvent) {
	f(e)
}

// An DeviceDidDisconnectHandler responds to DeviceDidDisconnectEvents.
type DeviceDidDisconnectHandler interface {
	DeviceDidDisconnect(*DeviceDidDisconnectEvent)
}

// A DeviceDidDisconnectHandlerFunc responds to DeviceDidDisconnectEvents.
type DeviceDidDisconnectHandlerFunc func(*DeviceDidDisconnectEvent)

// DeviceDidDisconnect calls f(e).
func (f DeviceDidDisconnectHandlerFunc) DeviceDidDisconnect(e *DeviceDidDisconnectEvent) {
	f(e)
}

// An KeyDownHandler resopnds to KeyDownEvents.
type KeyDownHandler interface {
	KeyDown(*KeyDownEvent)
}

// A KeyDownHandlerFunc responds to KeyDownEvents.
type KeyDownHandlerFunc func(*KeyDownEvent)

// KeyDown calls f(e).
func (f KeyDownHandlerFunc) KeyDown(e *KeyDownEvent) {
	f(e)
}

// An KeyUpHandler responds to KeyUpEvents.
type KeyUpHandler interface {
	KeyUp(*KeyUpEvent)
}

// A KeyUpHandlerFunc responds to KeyUpEvents.
type KeyUpHandlerFunc func(*KeyUpEvent)

// KeyUp calls f(e).
func (f KeyUpHandlerFunc) KeyUp(e *KeyUpEvent) {
	f(e)
}

// An TitleParametersDidChangeHandler responds to TitleParametersDidChangeEvents.
type TitleParametersDidChangeHandler interface {
	TitleParametersDidChange(*TitleParametersDidChangeEvent)
}

// A TitleParametersDidChangeHandlerFunc responds to TitleParametersDidChangeEvents.
type TitleParametersDidChangeHandlerFunc func(*TitleParametersDidChangeEvent)

// TitleParametersDidChange calls f(e)
func (f TitleParametersDidChangeHandlerFunc) TitleParametersDidChange(e *TitleParametersDidChangeEvent) {
	f(e)
}

// An WillAppearHandler handles an "willAppear" event.
type WillAppearHandler interface {
	WillAppear(*WillAppearEvent)
}

// A WillAppearHandlerFunc handles a WillAppearEvent.
type WillAppearHandlerFunc func(*WillAppearEvent)

// WillAppear invokes the handler func.
func (f WillAppearHandlerFunc) WillAppear(e *WillAppearEvent) {
	f(e)
}

// An WillDisappearHandler responds to WillDisappearEvents.
type WillDisappearHandler interface {
	WillDisappear(*WillDisappearEvent)
}

// A WillDisappearHandlerFunc responds to WillDisappearEvents.
type WillDisappearHandlerFunc func(*WillDisappearEvent)

// WillDisappear calls f(e)
func (f WillDisappearHandlerFunc) WillDisappear(e *WillDisappearEvent) {
	f(e)
}
