// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package reviews

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

func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews(in *jlexer.Lexer, out *UpdateReviewRequestDTO) {
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
		case "rating":
			out.Rating = uint8(in.Uint8())
		case "text":
			out.Text = string(in.String())
		case "is_private":
			out.IsPrivate = bool(in.Bool())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews(out *jwriter.Writer, in UpdateReviewRequestDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix[1:])
		out.Uint8(uint8(in.Rating))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"is_private\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsPrivate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UpdateReviewRequestDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UpdateReviewRequestDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UpdateReviewRequestDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UpdateReviewRequestDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews1(in *jlexer.Lexer, out *AddReviewRequestDTO) {
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
		case "rating":
			out.Rating = uint8(in.Uint8())
		case "text":
			out.Text = string(in.String())
		case "is_private":
			out.IsPrivate = bool(in.Bool())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews1(out *jwriter.Writer, in AddReviewRequestDTO) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"rating\":"
		out.RawString(prefix[1:])
		out.Uint8(uint8(in.Rating))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"is_private\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsPrivate))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AddReviewRequestDTO) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AddReviewRequestDTO) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AddReviewRequestDTO) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AddReviewRequestDTO) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryReviews1(l, v)
}