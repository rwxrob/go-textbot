package textbot

// ReponseMethod is a function that takes a natural language string and
// returns a natural language string suitable for passing to another
// ResponseMethod. Context must be the only means of passing data and
// parameters besides the natural language strings themselves.
type ResponseMethod func(txt string, ctxt *State) string

// Responders should generally be constructed through declaration rather
// than any constructor method to allow them to modularly be added
// through declared composition.
type Responder struct {

	// UUID must be version 4 random and are used to reference
	// responder sessions and other responder-specific information
	// in the textbot and state.
	UUID string `json:"uuid"`

	// RespondTo is the main method to be called. It is implemented
	// this way to allow it to be changed dynamically at run-time as
	// more intelligent textbots detect new conditions and want to
	// safely swap out how a reponse is handled. There are many
	// reasons for swapping out the methods including processing
	// different languages or AI adaptation to a given session.
	RespondTo ResponseMethod `json:"-"`
}
