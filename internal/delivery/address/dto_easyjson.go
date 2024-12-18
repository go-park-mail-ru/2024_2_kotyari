// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package address

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

func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryAddress(in *jlexer.Lexer, out *UpdateAddressRequest) {
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
		case "address":
			out.Address = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryAddress(out *jwriter.Writer, in UpdateAddressRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"address\":"
		out.RawString(prefix[1:])
		out.String(string(in.Address))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateAddressRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryAddress(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateAddressRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryAddress(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateAddressRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryAddress(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateAddressRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryAddress(l, v)
}
