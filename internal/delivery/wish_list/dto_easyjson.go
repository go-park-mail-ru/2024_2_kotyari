// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package wish_list

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

func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList(in *jlexer.Lexer, out *renameWishlistRequest) {
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
		case "new_name":
			out.NewName = string(in.String())
		case "link":
			out.Link = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList(out *jwriter.Writer, in renameWishlistRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"new_name\":"
		out.RawString(prefix[1:])
		out.String(string(in.NewName))
	}
	{
		const prefix string = ",\"link\":"
		out.RawString(prefix)
		out.String(string(in.Link))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v renameWishlistRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v renameWishlistRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *renameWishlistRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *renameWishlistRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList1(in *jlexer.Lexer, out *removeFromWishlistRequest) {
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
		case "links":
			if in.IsNull() {
				in.Skip()
				out.Links = nil
			} else {
				in.Delim('[')
				if out.Links == nil {
					if !in.IsDelim(']') {
						out.Links = make([]string, 0, 4)
					} else {
						out.Links = []string{}
					}
				} else {
					out.Links = (out.Links)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Links = append(out.Links, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "product_id":
			out.ProductId = uint32(in.Uint32())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList1(out *jwriter.Writer, in removeFromWishlistRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"links\":"
		out.RawString(prefix[1:])
		if in.Links == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Links {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"product_id\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.ProductId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v removeFromWishlistRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v removeFromWishlistRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *removeFromWishlistRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *removeFromWishlistRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList1(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList2(in *jlexer.Lexer, out *deleteWishlistsRequest) {
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
		case "link":
			out.Link = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList2(out *jwriter.Writer, in deleteWishlistsRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"link\":"
		out.RawString(prefix[1:])
		out.String(string(in.Link))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v deleteWishlistsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v deleteWishlistsRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *deleteWishlistsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *deleteWishlistsRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList2(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList3(in *jlexer.Lexer, out *createWishlistRequest) {
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
		case "name":
			out.Name = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList3(out *jwriter.Writer, in createWishlistRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v createWishlistRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v createWishlistRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *createWishlistRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *createWishlistRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList3(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList4(in *jlexer.Lexer, out *copyWishlistsRequest) {
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
		case "link":
			out.Link = string(in.String())
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList4(out *jwriter.Writer, in copyWishlistsRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"link\":"
		out.RawString(prefix[1:])
		out.String(string(in.Link))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v copyWishlistsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v copyWishlistsRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *copyWishlistsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *copyWishlistsRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList4(l, v)
}
func easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList5(in *jlexer.Lexer, out *addProductToWishlistsRequest) {
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
		case "product_id":
			out.ProductId = uint32(in.Uint32())
		case "links":
			if in.IsNull() {
				in.Skip()
				out.Links = nil
			} else {
				in.Delim('[')
				if out.Links == nil {
					if !in.IsDelim(']') {
						out.Links = make([]string, 0, 4)
					} else {
						out.Links = []string{}
					}
				} else {
					out.Links = (out.Links)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Links = append(out.Links, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList5(out *jwriter.Writer, in addProductToWishlistsRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"product_id\":"
		out.RawString(prefix[1:])
		out.Uint32(uint32(in.ProductId))
	}
	{
		const prefix string = ",\"links\":"
		out.RawString(prefix)
		if in.Links == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Links {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v addProductToWishlistsRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v addProductToWishlistsRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson56de76c1EncodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *addProductToWishlistsRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *addProductToWishlistsRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson56de76c1DecodeGithubComGoParkMailRu20242KotyariInternalDeliveryWishList5(l, v)
}
