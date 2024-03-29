// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package recommendation

import (
	json "encoding/json"
	constform "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
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

func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation(in *jlexer.Lexer, out *UserRecommend) {
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
		case "userId":
			out.UserId = uint(in.Uint())
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation(out *jwriter.Writer, in UserRecommend) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.UserId))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserRecommend) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserRecommend) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserRecommend) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserRecommend) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation(l, v)
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation1(in *jlexer.Lexer, out *Recommendation) {
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
		case "userId":
			out.Id = uint(in.Uint())
		case "name":
			out.Name = string(in.String())
		case "photos":
			if in.IsNull() {
				in.Skip()
				out.Photos = nil
			} else {
				in.Delim('[')
				if out.Photos == nil {
					if !in.IsDelim(']') {
						out.Photos = make([]uint, 0, 8)
					} else {
						out.Photos = []uint{}
					}
				} else {
					out.Photos = (out.Photos)[:0]
				}
				for !in.IsDelim(']') {
					var v1 uint
					v1 = uint(in.Uint())
					out.Photos = append(out.Photos, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "age":
			out.Age = int(in.Int())
		case "sex":
			out.Sex = constform.Sex(in.Int())
		case "description":
			out.Description = string(in.String())
		case "city":
			out.City = string(in.String())
		case "hashtags":
			if in.IsNull() {
				in.Skip()
				out.Hashtags = nil
			} else {
				in.Delim('[')
				if out.Hashtags == nil {
					if !in.IsDelim(']') {
						out.Hashtags = make([]string, 0, 4)
					} else {
						out.Hashtags = []string{}
					}
				} else {
					out.Hashtags = (out.Hashtags)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.Hashtags = append(out.Hashtags, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "zodiac":
			out.Zodiac = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "education":
			out.Education = string(in.String())
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation1(out *jwriter.Writer, in Recommendation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"photos\":"
		out.RawString(prefix)
		if in.Photos == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Photos {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.Uint(uint(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"age\":"
		out.RawString(prefix)
		out.Int(int(in.Age))
	}
	{
		const prefix string = ",\"sex\":"
		out.RawString(prefix)
		out.Int(int(in.Sex))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"hashtags\":"
		out.RawString(prefix)
		if in.Hashtags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Hashtags {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"zodiac\":"
		out.RawString(prefix)
		out.String(string(in.Zodiac))
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.String(string(in.Job))
	}
	{
		const prefix string = ",\"education\":"
		out.RawString(prefix)
		out.String(string(in.Education))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Recommendation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Recommendation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Recommendation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Recommendation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation1(l, v)
}
func easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation2(in *jlexer.Lexer, out *DBRecommendation) {
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
		case "userId":
			out.Id = uint(in.Uint())
		case "name":
			out.Name = string(in.String())
		case "birthDay":
			out.BirthDay = string(in.String())
		case "sex":
			out.Sex = constform.Sex(in.Int())
		case "description":
			out.Description = string(in.String())
		case "city":
			out.City = string(in.String())
		case "zodiac":
			out.Zodiac = string(in.String())
		case "job":
			out.Job = string(in.String())
		case "education":
			out.Education = string(in.String())
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
func easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation2(out *jwriter.Writer, in DBRecommendation) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix[1:])
		out.Uint(uint(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"birthDay\":"
		out.RawString(prefix)
		out.String(string(in.BirthDay))
	}
	{
		const prefix string = ",\"sex\":"
		out.RawString(prefix)
		out.Int(int(in.Sex))
	}
	{
		const prefix string = ",\"description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"city\":"
		out.RawString(prefix)
		out.String(string(in.City))
	}
	{
		const prefix string = ",\"zodiac\":"
		out.RawString(prefix)
		out.String(string(in.Zodiac))
	}
	{
		const prefix string = ",\"job\":"
		out.RawString(prefix)
		out.String(string(in.Job))
	}
	{
		const prefix string = ",\"education\":"
		out.RawString(prefix)
		out.String(string(in.Education))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v DBRecommendation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v DBRecommendation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6a975c40EncodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *DBRecommendation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *DBRecommendation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6a975c40DecodeGithubComGoParkMailRu20231MRGAGitInternalPkgRecommendation2(l, v)
}
