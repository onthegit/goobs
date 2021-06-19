// This file has been automatically generated. Don't edit it.

package sceneitems

import (
	requests "github.com/andreykaipov/goobs/api/requests"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
)

/*
GetSceneItemPropertiesParams represents the params body for the "GetSceneItemProperties" request.
Gets the scene specific properties of the specified source item.
Coordinates are relative to the item's parent (the scene or group it belongs to).
Since 4.3.0.
*/
type GetSceneItemPropertiesParams struct {
	requests.ParamsBasic

	// The item specification for this object.
	Item typedefs.Item `json:"item"`

	// Name of the scene the scene item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// GetSelfName just returns "GetSceneItemProperties".
func (o *GetSceneItemPropertiesParams) GetSelfName() string {
	return "GetSceneItemProperties"
}

/*
GetSceneItemPropertiesResponse represents the response body for the "GetSceneItemProperties" request.
Gets the scene specific properties of the specified source item.
Coordinates are relative to the item's parent (the scene or group it belongs to).
Since v4.3.0.
*/
type GetSceneItemPropertiesResponse struct {
	requests.ResponseBasic

	// The bounding box of the object (source, scene item, etc).
	Bounds typedefs.Bounds `json:"bounds"`

	// The crop specification for the object (source, scene item, etc).
	Crop typedefs.Crop `json:"crop"`

	// List of children (if this item is a group)
	GroupChildren []typedefs.SceneItemTransform `json:"groupChildren"`

	// Scene item height (base source height multiplied by the vertical scaling factor)
	Height float64 `json:"height"`

	// Scene Item ID.
	ItemId int `json:"itemId"`

	// If the source's transform is locked.
	Locked bool `json:"locked"`

	// If the source is muted.
	Muted bool `json:"muted"`

	// Scene Item name.
	Name string `json:"name"`

	// Name of the item's parent (if this item belongs to a group)
	ParentGroupName string `json:"parentGroupName"`

	// The position of the object (source, scene item, etc).
	Position typedefs.Position `json:"position"`

	// The clockwise rotation of the item in degrees around the point of alignment.
	Rotation float64 `json:"rotation"`

	// The scaling specification for the object (source, scene item, etc).
	Scale typedefs.Scale `json:"scale"`

	// Base source (without scaling) of the source
	SourceHeight int `json:"sourceHeight"`

	// Base width (without scaling) of the source
	SourceWidth int `json:"sourceWidth"`

	// If the source is visible.
	Visible bool `json:"visible"`

	// Scene item width (base source width multiplied by the horizontal scaling factor)
	Width float64 `json:"width"`
}

// GetSceneItemProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemProperties(params *GetSceneItemPropertiesParams) (*GetSceneItemPropertiesResponse, error) {
	data := &GetSceneItemPropertiesResponse{}
	if err := c.SendRequest(params, data); err != nil {
		return nil, err
	}
	return data, nil
}
