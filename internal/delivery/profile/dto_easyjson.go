// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package profile

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

func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile(in *jlexer.Lexer, out *UpdateProfile) {
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
		case "email":
			out.Email = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "gender":
			out.Gender = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile(out *jwriter.Writer, in UpdateProfile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"gender\":"
		out.RawString(prefix)
		out.String(string(in.Gender))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateProfile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateProfile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateProfile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateProfile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile1(in *jlexer.Lexer, out *UpdatePasswordRequest) {
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
		case "old_password":
			out.OldPassword = string(in.String())
		case "new_password":
			out.NewPassword = string(in.String())
		case "repeat_password":
			out.RepeatPassword = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile1(out *jwriter.Writer, in UpdatePasswordRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"old_password\":"
		out.RawString(prefix[1:])
		out.String(string(in.OldPassword))
	}
	{
		const prefix string = ",\"new_password\":"
		out.RawString(prefix)
		out.String(string(in.NewPassword))
	}
	{
		const prefix string = ",\"repeat_password\":"
		out.RawString(prefix)
		out.String(string(in.RepeatPassword))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdatePasswordRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdatePasswordRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdatePasswordRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdatePasswordRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryProfile1(l, v)
}