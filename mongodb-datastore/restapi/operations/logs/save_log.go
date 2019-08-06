// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// SaveLogHandlerFunc turns a function with the right signature into a save log handler
type SaveLogHandlerFunc func(SaveLogParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveLogHandlerFunc) Handle(params SaveLogParams) middleware.Responder {
	return fn(params)
}

// SaveLogHandler interface for that can handle valid save log params
type SaveLogHandler interface {
	Handle(SaveLogParams) middleware.Responder
}

// NewSaveLog creates a new http.Handler for the save log operation
func NewSaveLog(ctx *middleware.Context, handler SaveLogHandler) *SaveLog {
	return &SaveLog{Context: ctx, Handler: handler}
}

/*SaveLog swagger:route POST /logs logs saveLog

Saves a log to the datastore

*/
type SaveLog struct {
	Context *middleware.Context
	Handler SaveLogHandler
}

func (o *SaveLog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSaveLogParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// SaveLogDefaultBody save log default body
// swagger:model SaveLogDefaultBody
type SaveLogDefaultBody struct {

	// code
	Code int64 `json:"code,omitempty"`

	// fields
	Fields string `json:"fields,omitempty"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this save log default body
func (o *SaveLogDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SaveLogDefaultBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("saveLog default"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SaveLogDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SaveLogDefaultBody) UnmarshalBinary(b []byte) error {
	var res SaveLogDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// SaveLogParamsBodyItems0 save log params body items0
// swagger:model SaveLogParamsBodyItems0
type SaveLogParamsBodyItems0 struct {

	// event Id
	EventID string `json:"eventId,omitempty"`

	// keptn context
	KeptnContext string `json:"keptnContext,omitempty"`

	// keptn service
	KeptnService string `json:"keptnService,omitempty"`

	// log level
	LogLevel string `json:"logLevel,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this save log params body items0
func (o *SaveLogParamsBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SaveLogParamsBodyItems0) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SaveLogParamsBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SaveLogParamsBodyItems0) UnmarshalBinary(b []byte) error {
	var res SaveLogParamsBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
