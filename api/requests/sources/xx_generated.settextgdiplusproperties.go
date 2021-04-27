// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
SetTextGDIPlusPropertiesParams represents the params body for the "SetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextGDIPlusProperties.
*/
type SetTextGDIPlusPropertiesParams struct {
	requests.Params

	// Text Alignment ("left", "center", "right").
	Align string `json:"align"`

	// Background color.
	BkColor int `json:"bk-color"`

	// Background opacity (0-100).
	BkOpacity int `json:"bk-opacity"`

	// Chat log.
	Chatlog bool `json:"chatlog"`

	// Chat log lines.
	ChatlogLines int `json:"chatlog_lines"`

	// Text color.
	Color int `json:"color"`

	// Extents wrap.
	Extents bool `json:"extents"`

	// Extents cx.
	ExtentsCx int `json:"extents_cx"`

	// Extents cy.
	ExtentsCy int `json:"extents_cy"`

	// File path name.
	File string `json:"file"`

	Font struct {
		// Font face.
		Face string `json:"face"`

		// Font text styling flag. `Bold=1, Italic=2, Bold Italic=3, Underline=5, Strikeout=8`
		Flags int `json:"flags"`

		// Font text size.
		Size int `json:"size"`

		// Font Style (unknown function).
		Style string `json:"style"`
	} `json:"font"`

	// Gradient enabled.
	Gradient bool `json:"gradient"`

	// Gradient color.
	GradientColor int `json:"gradient_color"`

	// Gradient direction.
	GradientDir float64 `json:"gradient_dir"`

	// Gradient opacity (0-100).
	GradientOpacity int `json:"gradient_opacity"`

	// Outline.
	Outline bool `json:"outline"`

	// Outline color.
	OutlineColor int `json:"outline_color"`

	// Outline opacity (0-100).
	OutlineOpacity int `json:"outline_opacity"`

	// Outline size.
	OutlineSize int `json:"outline_size"`

	// Read text from the specified file.
	ReadFromFile bool `json:"read_from_file"`

	// Visibility of the scene item.
	Render bool `json:"render"`

	// Name of the source.
	Source string `json:"source"`

	// Text content to be displayed.
	Text string `json:"text"`

	// Text vertical alignment ("top", "center", "bottom").
	Valign string `json:"valign"`

	// Vertical text enabled.
	Vertical bool `json:"vertical"`
}

// GetRequestType returns the RequestType of SetTextGDIPlusPropertiesParams
func (o *SetTextGDIPlusPropertiesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetTextGDIPlusPropertiesParams
func (o *SetTextGDIPlusPropertiesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetTextGDIPlusPropertiesParams
func (o *SetTextGDIPlusPropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetTextGDIPlusPropertiesResponse represents the response body for the "SetTextGDIPlusProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetTextGDIPlusProperties.
*/
type SetTextGDIPlusPropertiesResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetTextGDIPlusPropertiesResponse
func (o *SetTextGDIPlusPropertiesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetTextGDIPlusPropertiesResponse
func (o *SetTextGDIPlusPropertiesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetTextGDIPlusPropertiesResponse
func (o *SetTextGDIPlusPropertiesResponse) GetError() string {
	return o.Error
}

// SetTextGDIPlusProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetTextGDIPlusProperties(
	params *SetTextGDIPlusPropertiesParams,
) (*SetTextGDIPlusPropertiesResponse, error) {
	params.RequestType = "SetTextGDIPlusProperties"
	data := &SetTextGDIPlusPropertiesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
