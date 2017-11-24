package coa

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *Account) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Number":
			z.Number, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Tags":
			var zb0002 uint32
			zb0002, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Tags) >= int(zb0002) {
				z.Tags = (z.Tags)[:zb0002]
			} else {
				z.Tags = make([]string, zb0002)
			}
			for za0001 := range z.Tags {
				z.Tags[za0001], err = dc.ReadString()
				if err != nil {
					return
				}
			}
		case "Parent":
			z.Parent, err = dc.ReadString()
			if err != nil {
				return
			}
		case "User":
			z.User, err = dc.ReadString()
			if err != nil {
				return
			}
		case "AsOf":
			z.AsOf, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "Created":
			z.Created, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "Removed":
			z.Removed, err = dc.ReadTime()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *Account) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 9
	// write "Id"
	err = en.Append(0x89, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Id)
	if err != nil {
		return
	}
	// write "Number"
	err = en.Append(0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Number)
	if err != nil {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "Tags"
	err = en.Append(0xa4, 0x54, 0x61, 0x67, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Tags)))
	if err != nil {
		return
	}
	for za0001 := range z.Tags {
		err = en.WriteString(z.Tags[za0001])
		if err != nil {
			return
		}
	}
	// write "Parent"
	err = en.Append(0xa6, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Parent)
	if err != nil {
		return
	}
	// write "User"
	err = en.Append(0xa4, 0x55, 0x73, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.User)
	if err != nil {
		return
	}
	// write "AsOf"
	err = en.Append(0xa4, 0x41, 0x73, 0x4f, 0x66)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.AsOf)
	if err != nil {
		return
	}
	// write "Created"
	err = en.Append(0xa7, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.Created)
	if err != nil {
		return
	}
	// write "Removed"
	err = en.Append(0xa7, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.Removed)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *Account) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 9
	// string "Id"
	o = append(o, 0x89, 0xa2, 0x49, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "Number"
	o = append(o, 0xa6, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72)
	o = msgp.AppendString(o, z.Number)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "Tags"
	o = append(o, 0xa4, 0x54, 0x61, 0x67, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Tags)))
	for za0001 := range z.Tags {
		o = msgp.AppendString(o, z.Tags[za0001])
	}
	// string "Parent"
	o = append(o, 0xa6, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74)
	o = msgp.AppendString(o, z.Parent)
	// string "User"
	o = append(o, 0xa4, 0x55, 0x73, 0x65, 0x72)
	o = msgp.AppendString(o, z.User)
	// string "AsOf"
	o = append(o, 0xa4, 0x41, 0x73, 0x4f, 0x66)
	o = msgp.AppendTime(o, z.AsOf)
	// string "Created"
	o = append(o, 0xa7, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendTime(o, z.Created)
	// string "Removed"
	o = append(o, 0xa7, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64)
	o = msgp.AppendTime(o, z.Removed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Account) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Number":
			z.Number, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Tags":
			var zb0002 uint32
			zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Tags) >= int(zb0002) {
				z.Tags = (z.Tags)[:zb0002]
			} else {
				z.Tags = make([]string, zb0002)
			}
			for za0001 := range z.Tags {
				z.Tags[za0001], bts, err = msgp.ReadStringBytes(bts)
				if err != nil {
					return
				}
			}
		case "Parent":
			z.Parent, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "User":
			z.User, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AsOf":
			z.AsOf, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "Created":
			z.Created, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "Removed":
			z.Removed, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *Account) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.Id) + 7 + msgp.StringPrefixSize + len(z.Number) + 5 + msgp.StringPrefixSize + len(z.Name) + 5 + msgp.ArrayHeaderSize
	for za0001 := range z.Tags {
		s += msgp.StringPrefixSize + len(z.Tags[za0001])
	}
	s += 7 + msgp.StringPrefixSize + len(z.Parent) + 5 + msgp.StringPrefixSize + len(z.User) + 5 + msgp.TimeSize + 8 + msgp.TimeSize + 8 + msgp.TimeSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *Accounts) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Accounts, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Account)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z Accounts) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z Accounts) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *Accounts) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(Accounts, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(Account)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z Accounts) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ChartOfAccounts) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, err = dc.ReadString()
			if err != nil {
				return
			}
		case "Name":
			z.Name, err = dc.ReadString()
			if err != nil {
				return
			}
		case "RetainedEarningsAccount":
			z.RetainedEarningsAccount, err = dc.ReadString()
			if err != nil {
				return
			}
		case "User":
			z.User, err = dc.ReadString()
			if err != nil {
				return
			}
		case "AsOf":
			z.AsOf, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "Created":
			z.Created, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "Removed":
			z.Removed, err = dc.ReadTime()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ChartOfAccounts) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "Id"
	err = en.Append(0x87, 0xa2, 0x49, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Id)
	if err != nil {
		return
	}
	// write "Name"
	err = en.Append(0xa4, 0x4e, 0x61, 0x6d, 0x65)
	if err != nil {
		return err
	}
	err = en.WriteString(z.Name)
	if err != nil {
		return
	}
	// write "RetainedEarningsAccount"
	err = en.Append(0xb7, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x45, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteString(z.RetainedEarningsAccount)
	if err != nil {
		return
	}
	// write "User"
	err = en.Append(0xa4, 0x55, 0x73, 0x65, 0x72)
	if err != nil {
		return err
	}
	err = en.WriteString(z.User)
	if err != nil {
		return
	}
	// write "AsOf"
	err = en.Append(0xa4, 0x41, 0x73, 0x4f, 0x66)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.AsOf)
	if err != nil {
		return
	}
	// write "Created"
	err = en.Append(0xa7, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.Created)
	if err != nil {
		return
	}
	// write "Removed"
	err = en.Append(0xa7, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64)
	if err != nil {
		return err
	}
	err = en.WriteTime(z.Removed)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ChartOfAccounts) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "Id"
	o = append(o, 0x87, 0xa2, 0x49, 0x64)
	o = msgp.AppendString(o, z.Id)
	// string "Name"
	o = append(o, 0xa4, 0x4e, 0x61, 0x6d, 0x65)
	o = msgp.AppendString(o, z.Name)
	// string "RetainedEarningsAccount"
	o = append(o, 0xb7, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x64, 0x45, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74)
	o = msgp.AppendString(o, z.RetainedEarningsAccount)
	// string "User"
	o = append(o, 0xa4, 0x55, 0x73, 0x65, 0x72)
	o = msgp.AppendString(o, z.User)
	// string "AsOf"
	o = append(o, 0xa4, 0x41, 0x73, 0x4f, 0x66)
	o = msgp.AppendTime(o, z.AsOf)
	// string "Created"
	o = append(o, 0xa7, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64)
	o = msgp.AppendTime(o, z.Created)
	// string "Removed"
	o = append(o, 0xa7, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64)
	o = msgp.AppendTime(o, z.Removed)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ChartOfAccounts) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "Id":
			z.Id, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "Name":
			z.Name, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "RetainedEarningsAccount":
			z.RetainedEarningsAccount, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "User":
			z.User, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "AsOf":
			z.AsOf, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "Created":
			z.Created, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "Removed":
			z.Removed, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ChartOfAccounts) Msgsize() (s int) {
	s = 1 + 3 + msgp.StringPrefixSize + len(z.Id) + 5 + msgp.StringPrefixSize + len(z.Name) + 24 + msgp.StringPrefixSize + len(z.RetainedEarningsAccount) + 5 + msgp.StringPrefixSize + len(z.User) + 5 + msgp.TimeSize + 8 + msgp.TimeSize + 8 + msgp.TimeSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ChartsOfAccounts) DecodeMsg(dc *msgp.Reader) (err error) {
	var zb0002 uint32
	zb0002, err = dc.ReadArrayHeader()
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(ChartsOfAccounts, zb0002)
	}
	for zb0001 := range *z {
		if dc.IsNil() {
			err = dc.ReadNil()
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(ChartOfAccounts)
			}
			err = (*z)[zb0001].DecodeMsg(dc)
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z ChartsOfAccounts) EncodeMsg(en *msgp.Writer) (err error) {
	err = en.WriteArrayHeader(uint32(len(z)))
	if err != nil {
		return
	}
	for zb0003 := range z {
		if z[zb0003] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z[zb0003].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z ChartsOfAccounts) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	o = msgp.AppendArrayHeader(o, uint32(len(z)))
	for zb0003 := range z {
		if z[zb0003] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z[zb0003].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ChartsOfAccounts) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var zb0002 uint32
	zb0002, bts, err = msgp.ReadArrayHeaderBytes(bts)
	if err != nil {
		return
	}
	if cap((*z)) >= int(zb0002) {
		(*z) = (*z)[:zb0002]
	} else {
		(*z) = make(ChartsOfAccounts, zb0002)
	}
	for zb0001 := range *z {
		if msgp.IsNil(bts) {
			bts, err = msgp.ReadNilBytes(bts)
			if err != nil {
				return
			}
			(*z)[zb0001] = nil
		} else {
			if (*z)[zb0001] == nil {
				(*z)[zb0001] = new(ChartOfAccounts)
			}
			bts, err = (*z)[zb0001].UnmarshalMsg(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z ChartsOfAccounts) Msgsize() (s int) {
	s = msgp.ArrayHeaderSize
	for zb0003 := range z {
		if z[zb0003] == nil {
			s += msgp.NilSize
		} else {
			s += z[zb0003].Msgsize()
		}
	}
	return
}
