// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package repository

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson32ceb8acDecodeGithubComEvgen1067Hw12131415CalendarInternalRepository(in *jlexer.Lexer, out *Notice) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "eventId":
			out.EventID = EventID(in.Int64())
		case "title":
			out.Title = string(in.String())
		case "datetime":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Datetime).UnmarshalJSON(data))
			}
		case "ownerId":
			out.OwnerID = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson32ceb8acEncodeGithubComEvgen1067Hw12131415CalendarInternalRepository(out *jwriter.Writer, in Notice) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"eventId\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.EventID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"datetime\":"
		out.RawString(prefix)
		out.Raw((in.Datetime).MarshalJSON())
	}
	{
		const prefix string = ",\"ownerId\":"
		out.RawString(prefix)
		out.Int64(int64(in.OwnerID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Notice) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson32ceb8acEncodeGithubComEvgen1067Hw12131415CalendarInternalRepository(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Notice) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson32ceb8acEncodeGithubComEvgen1067Hw12131415CalendarInternalRepository(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Notice) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson32ceb8acDecodeGithubComEvgen1067Hw12131415CalendarInternalRepository(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Notice) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson32ceb8acDecodeGithubComEvgen1067Hw12131415CalendarInternalRepository(l, v)
}
func easyjson32ceb8acDecodeGithubComEvgen1067Hw12131415CalendarInternalRepository1(in *jlexer.Lexer, out *Event) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = EventID(in.Int64())
		case "title":
			out.Title = string(in.String())
		case "description":
			out.Description = string(in.String())
		case "dateStart":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateStart).UnmarshalJSON(data))
			}
		case "dateEnd":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.DateEnd).UnmarshalJSON(data))
			}
		case "notifyIn":
			out.NotifyIn = int64(in.Int64())
		case "ownerId":
			out.OwnerID = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson32ceb8acEncodeGithubComEvgen1067Hw12131415CalendarInternalRepository1(out *jwriter.Writer, in Event) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"title\":"
		out.RawString(prefix)
		out.String(string(in.Title))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"dateStart\":"
		out.RawString(prefix)
		out.Raw((in.DateStart).MarshalJSON())
	}
	{
		const prefix string = ",\"dateEnd\":"
		out.RawString(prefix)
		out.Raw((in.DateEnd).MarshalJSON())
	}
	{
		const prefix string = ",\"notifyIn\":"
		out.RawString(prefix)
		out.Int64(int64(in.NotifyIn))
	}
	{
		const prefix string = ",\"ownerId\":"
		out.RawString(prefix)
		out.Int64(int64(in.OwnerID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Event) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson32ceb8acEncodeGithubComEvgen1067Hw12131415CalendarInternalRepository1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Event) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson32ceb8acEncodeGithubComEvgen1067Hw12131415CalendarInternalRepository1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Event) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson32ceb8acDecodeGithubComEvgen1067Hw12131415CalendarInternalRepository1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Event) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson32ceb8acDecodeGithubComEvgen1067Hw12131415CalendarInternalRepository1(l, v)
}