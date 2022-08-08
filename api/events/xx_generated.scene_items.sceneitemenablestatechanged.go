// This file has been automatically generated. Don't edit it.

package events

/*
SceneItemEnableStateChanged represents the event body for the "SceneItemEnableStateChanged" event.
Since v5.0.0.
*/
type SceneItemEnableStateChanged struct {
	// Whether the scene item is enabled (visible)
	SceneItemEnabled bool `json:"sceneItemEnabled,omitempty"`

	// Numeric ID of the scene item
	SceneItemId float64 `json:"sceneItemId,omitempty"`

	// Name of the scene the item is in
	SceneName string `json:"sceneName,omitempty"`
}