// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package data

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

func easyjsonC80ae7adDecodeGithubComTheZionMatrixCoreAppUserServiceInternalData(in *jlexer.Lexer, out *Profile) {
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
		case "createdAt":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CreatedAt).UnmarshalJSON(data))
			}
		case "updated":
			out.Updated = int64(in.Int64())
		case "uuid":
			out.Uuid = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "school":
			out.School = string(in.String())
		case "company":
			out.Company = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "homepage":
			out.Homepage = string(in.String())
		case "introduce":
			out.Introduce = string(in.String())
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
func easyjsonC80ae7adEncodeGithubComTheZionMatrixCoreAppUserServiceInternalData(out *jwriter.Writer, in Profile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"createdAt\":"
		out.RawString(prefix[1:])
		out.Raw((in.CreatedAt).MarshalJSON())
	}
	{
		const prefix string = ",\"updated\":"
		out.RawString(prefix)
		out.Int64(int64(in.Updated))
	}
	{
		const prefix string = ",\"uuid\":"
		out.RawString(prefix)
		out.String(string(in.Uuid))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	{
		const prefix string = ",\"school\":"
		out.RawString(prefix)
		out.String(string(in.School))
	}
	{
		const prefix string = ",\"company\":"
		out.RawString(prefix)
		out.String(string(in.Company))
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.String(string(in.Job))
	}
	{
		const prefix string = ",\"homepage\":"
		out.RawString(prefix)
		out.String(string(in.Homepage))
	}
	{
		const prefix string = ",\"introduce\":"
		out.RawString(prefix)
		out.String(string(in.Introduce))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Profile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComTheZionMatrixCoreAppUserServiceInternalData(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Profile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComTheZionMatrixCoreAppUserServiceInternalData(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Profile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComTheZionMatrixCoreAppUserServiceInternalData(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Profile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComTheZionMatrixCoreAppUserServiceInternalData(l, v)
}
