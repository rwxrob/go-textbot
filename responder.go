package textbot

type Responder interface {

	// Consistently returns the same UUID for this specific
	// responder. Usually this is a hard-coded constant. The UUID is
	// always associated with this specific responder in this
	// specific package. If the package location or name changes in
	// any way the UUID needs to be reset.
	UUID() string

	// Keys are the names of the context keys used in the RespondTo
	// context State. This provides visibility into the keys for
	// other responder classes or textbots when managing the context
	// state. In some sense, this allows context garbage collection.
	// When a responder is removed from a bot the bot can clear the
	// context state keys for that bot so long as no other responder
	// also needs those keys.
	Keys() []string

	// RespondTo processes a natural language text input in a given
	// context represented by the State (which is usually the state
	// of the TextBot itself). It returns a blank string unless
	// there is something significant to say/respond.
	RespondTo(text string, context *State) string

	// Must return a JSON representation of the current responder
	// state (not to be confused with the context of the RespondTo()
	// method). Can be empty ({}).
	String() string
}
