// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/degenerat3/meteor/core/ent/action"
	"github.com/degenerat3/meteor/core/ent/bot"
	"github.com/degenerat3/meteor/core/ent/host"
	"github.com/degenerat3/meteor/core/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	actionFields := schema.Action{}.Fields()
	_ = actionFields
	// actionDescQueued is the schema descriptor for queued field.
	actionDescQueued := actionFields[3].Descriptor()
	// action.DefaultQueued holds the default value on creation for the queued field.
	action.DefaultQueued = actionDescQueued.Default.(bool)
	// actionDescResponded is the schema descriptor for responded field.
	actionDescResponded := actionFields[4].Descriptor()
	// action.DefaultResponded holds the default value on creation for the responded field.
	action.DefaultResponded = actionDescResponded.Default.(bool)
	// actionDescResult is the schema descriptor for result field.
	actionDescResult := actionFields[5].Descriptor()
	// action.DefaultResult holds the default value on creation for the result field.
	action.DefaultResult = actionDescResult.Default.(string)
	botFields := schema.Bot{}.Fields()
	_ = botFields
	// botDescLastSeen is the schema descriptor for lastSeen field.
	botDescLastSeen := botFields[3].Descriptor()
	// bot.DefaultLastSeen holds the default value on creation for the lastSeen field.
	bot.DefaultLastSeen = botDescLastSeen.Default.(int)
	hostFields := schema.Host{}.Fields()
	_ = hostFields
	// hostDescLastSeen is the schema descriptor for lastSeen field.
	hostDescLastSeen := hostFields[2].Descriptor()
	// host.DefaultLastSeen holds the default value on creation for the lastSeen field.
	host.DefaultLastSeen = hostDescLastSeen.Default.(int)
}
