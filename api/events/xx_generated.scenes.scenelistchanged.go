// This file has been automatically generated. Don't edit it.

package events

import typedefs "github.com/onthegit/goobs/api/typedefs"

/*
Represents the event body for the SceneListChanged event.

The list of scenes has changed.

TODO: Make OBS fire this event when scenes are reordered.
*/
type SceneListChanged struct {
	// Updated array of scenes
	Scenes []*typedefs.Scene `json:"scenes,omitempty"`
}
