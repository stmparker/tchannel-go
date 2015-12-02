// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package test

import (
	"bytes"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var GoUnusedProtection__ int

// Attributes:
//  - B1
//  - S2
//  - I3
type Data struct {
	B1 bool   `thrift:"b1,1" db:"b1" json:"b1"`
	S2 string `thrift:"s2,2" db:"s2" json:"s2"`
	I3 int32  `thrift:"i3,3" db:"i3" json:"i3"`
}

func NewData() *Data {
	return &Data{}
}

func (p *Data) GetB1() bool {
	return p.B1
}

func (p *Data) GetS2() string {
	return p.S2
}

func (p *Data) GetI3() int32 {
	return p.I3
}
func (p *Data) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *Data) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.B1 = v
	}
	return nil
}

func (p *Data) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 2: ", err)
	} else {
		p.S2 = v
	}
	return nil
}

func (p *Data) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return thrift.PrependError("error reading field 3: ", err)
	} else {
		p.I3 = v
	}
	return nil
}

func (p *Data) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Data"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *Data) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("b1", thrift.BOOL, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:b1: ", p), err)
	}
	if err := oprot.WriteBool(bool(p.B1)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.b1 (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:b1: ", p), err)
	}
	return err
}

func (p *Data) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("s2", thrift.STRING, 2); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 2:s2: ", p), err)
	}
	if err := oprot.WriteString(string(p.S2)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.s2 (2) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 2:s2: ", p), err)
	}
	return err
}

func (p *Data) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("i3", thrift.I32, 3); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 3:i3: ", p), err)
	}
	if err := oprot.WriteI32(int32(p.I3)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.i3 (3) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 3:i3: ", p), err)
	}
	return err
}

func (p *Data) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Data(%+v)", *p)
}

// Attributes:
//  - Message
type SimpleErr struct {
	Message string `thrift:"message,1" db:"message" json:"message"`
}

func NewSimpleErr() *SimpleErr {
	return &SimpleErr{}
}

func (p *SimpleErr) GetMessage() string {
	return p.Message
}
func (p *SimpleErr) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *SimpleErr) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Message = v
	}
	return nil
}

func (p *SimpleErr) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("SimpleErr"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *SimpleErr) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("message", thrift.STRING, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:message: ", p), err)
	}
	if err := oprot.WriteString(string(p.Message)); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T.message (1) field write error: ", p), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:message: ", p), err)
	}
	return err
}

func (p *SimpleErr) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("SimpleErr(%+v)", *p)
}

func (p *SimpleErr) Error() string {
	return p.String()
}
