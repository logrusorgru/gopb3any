// lifo storage
package lis

import (
	"github.com/golang/protobuf/proto"
	"github.com/logrusorgru/gopb3any/msg"

	"encoding/binary"
	"errors"
	"github.com/logrusorgru/lifo"
	"reflect"
)

var ErrNoOne = errors.New("lis.TypeRegister.Get: no one")

// TypeRegister - type register
type TypeRegister map[string]reflect.Type

// Ser registers new type
func (t TypeRegister) Set(i interface{}) {
	if reflect.ValueOf(i).Kind() != reflect.Ptr {
		panic(errors.New("TypeRegister.Set() argument must to be a pointer"))
	}
	t[reflect.TypeOf(i).String()] = reflect.TypeOf(i)
}

// Get element of type, if no one - err will be ErrNoOne
func (t TypeRegister) Get(name string) (interface{}, error) {
	if typ, ok := t[name]; ok {
		return reflect.New(typ.Elem()).Elem().Addr().Interface(), nil
	}
	return nil, ErrNoOne
}

// shared type register
var TypeReg = make(TypeRegister)

// Repo - is LIFO pb storage
type Repo struct {
	lifo.Buffer
}

// Push values to storage (string and proto.Mesage)
func (r *Repo) Push(key string, m proto.Message) error {
	type_usl := reflect.TypeOf(m).String()
	value, err := proto.Marshal(m)
	if err != nil {
		return err
	}
	msg := &msg.Sih{
		Key: key,
		Value: &msg.Any{
			TypeUrl: type_usl,
			Value:   value,
		},
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	n, err := r.Write(data)
	if err != nil {
		return err
	}
	err = binary.Write(r, binary.LittleEndian, int32(n))
	return err
}

// Pop values from storage
func (r *Repo) Pop() (string, proto.Message, error) {
	var n int32
	err := binary.Read(r, binary.LittleEndian, &n)
	if err != nil {
		return "", nil, err
	}
	data := r.Next(int(n))
	msg := new(msg.Sih)
	err = proto.Unmarshal(data, msg)
	if err != nil {
		return "", nil, err
	}
	if msg.Value == nil {
		return msg.Key, nil, nil
	}
	any, err := TypeReg.Get(msg.Value.TypeUrl)
	if err != nil {
		return "", nil, err
	}
	pm := any.(proto.Message)
	err = proto.Unmarshal(msg.Value.Value, pm)
	return msg.Key, pm, err
}
