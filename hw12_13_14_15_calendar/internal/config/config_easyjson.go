// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package config

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

func easyjson6615c02eDecodeGithubComEvgen1067Hw12131415CalendarInternalConfig(in *jlexer.Lexer, out *Config) {
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
		case "logger":
			easyjson6615c02eDecode(in, &out.Logger)
		case "http":
			easyjson6615c02eDecode1(in, &out.HTTP)
		case "grpc":
			easyjson6615c02eDecode1(in, &out.GRPC)
		case "sql":
			out.SQL = bool(in.Bool())
		case "db":
			easyjson6615c02eDecode2(in, &out.DB)
		case "amqp":
			easyjson6615c02eDecode3(in, &out.AMQP)
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
func easyjson6615c02eEncodeGithubComEvgen1067Hw12131415CalendarInternalConfig(out *jwriter.Writer, in Config) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"logger\":"
		out.RawString(prefix[1:])
		easyjson6615c02eEncode(out, in.Logger)
	}
	{
		const prefix string = ",\"http\":"
		out.RawString(prefix)
		easyjson6615c02eEncode1(out, in.HTTP)
	}
	{
		const prefix string = ",\"grpc\":"
		out.RawString(prefix)
		easyjson6615c02eEncode1(out, in.GRPC)
	}
	{
		const prefix string = ",\"sql\":"
		out.RawString(prefix)
		out.Bool(bool(in.SQL))
	}
	{
		const prefix string = ",\"db\":"
		out.RawString(prefix)
		easyjson6615c02eEncode2(out, in.DB)
	}
	{
		const prefix string = ",\"amqp\":"
		out.RawString(prefix)
		easyjson6615c02eEncode3(out, in.AMQP)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Config) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6615c02eEncodeGithubComEvgen1067Hw12131415CalendarInternalConfig(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Config) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6615c02eEncodeGithubComEvgen1067Hw12131415CalendarInternalConfig(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Config) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6615c02eDecodeGithubComEvgen1067Hw12131415CalendarInternalConfig(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Config) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6615c02eDecodeGithubComEvgen1067Hw12131415CalendarInternalConfig(l, v)
}
func easyjson6615c02eDecode3(in *jlexer.Lexer, out *struct {
	URI   string `json:"uri"`
	Queue string `json:"queue"`
}) {
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
		case "uri":
			out.URI = string(in.String())
		case "queue":
			out.Queue = string(in.String())
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
func easyjson6615c02eEncode3(out *jwriter.Writer, in struct {
	URI   string `json:"uri"`
	Queue string `json:"queue"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"uri\":"
		out.RawString(prefix[1:])
		out.String(string(in.URI))
	}
	{
		const prefix string = ",\"queue\":"
		out.RawString(prefix)
		out.String(string(in.Queue))
	}
	out.RawByte('}')
}
func easyjson6615c02eDecode2(in *jlexer.Lexer, out *struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslMode"`
}) {
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
		case "host":
			out.Host = string(in.String())
		case "port":
			out.Port = string(in.String())
		case "user":
			out.User = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "database":
			out.Database = string(in.String())
		case "sslMode":
			out.SSLMode = string(in.String())
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
func easyjson6615c02eEncode2(out *jwriter.Writer, in struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslMode"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"host\":"
		out.RawString(prefix[1:])
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.String(string(in.Port))
	}
	{
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		out.String(string(in.User))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	{
		const prefix string = ",\"database\":"
		out.RawString(prefix)
		out.String(string(in.Database))
	}
	{
		const prefix string = ",\"sslMode\":"
		out.RawString(prefix)
		out.String(string(in.SSLMode))
	}
	out.RawByte('}')
}
func easyjson6615c02eDecode1(in *jlexer.Lexer, out *struct {
	Host string `json:"host"`
	Port string `json:"port"`
}) {
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
		case "host":
			out.Host = string(in.String())
		case "port":
			out.Port = string(in.String())
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
func easyjson6615c02eEncode1(out *jwriter.Writer, in struct {
	Host string `json:"host"`
	Port string `json:"port"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"host\":"
		out.RawString(prefix[1:])
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.String(string(in.Port))
	}
	out.RawByte('}')
}
func easyjson6615c02eDecode(in *jlexer.Lexer, out *struct {
	Level string `json:"level"`
	File  string `json:"file"`
}) {
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
		case "level":
			out.Level = string(in.String())
		case "file":
			out.File = string(in.String())
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
func easyjson6615c02eEncode(out *jwriter.Writer, in struct {
	Level string `json:"level"`
	File  string `json:"file"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"level\":"
		out.RawString(prefix[1:])
		out.String(string(in.Level))
	}
	{
		const prefix string = ",\"file\":"
		out.RawString(prefix)
		out.String(string(in.File))
	}
	out.RawByte('}')
}
