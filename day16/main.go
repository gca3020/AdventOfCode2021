package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	text := input2slice("day16/input")
	decoder := NewDecoder(text[0])
	p := decoder.NextPacket()
	fmt.Println("Got a total version count of", p.GetVersion())
	fmt.Println("The value of this operation is", p.GetValue())
}

type TypeId uint

const (
	Sum TypeId = iota
	Product
	Minimum
	Maximum
	Literal
	GreaterThan
	LessThan
	EqualTo
)

type Packet struct {
	version    uint
	typeId     TypeId
	literal    uint
	subPackets []*Packet
}

func (p *Packet) GetVersion() uint {
	subVersion := uint(0)
	for _, sp := range p.subPackets {
		subVersion += sp.GetVersion()
	}
	return p.version + subVersion
}

func (p *Packet) GetValue() uint {
	switch p.typeId {
	case Sum:
		sum := uint(0)
		for _, sp := range p.subPackets {
			sum += sp.GetValue()
		}
		return sum
	case Product:
		prod := uint(1)
		for _, sp := range p.subPackets {
			prod *= sp.GetValue()
		}
		return prod
	case Minimum:
		value := uint(math.MaxUint)
		for _, sp := range p.subPackets {
			v := sp.GetValue()
			if v < value {
				value = v
			}
		}
		return value
	case Maximum:
		value := uint(0)
		for _, sp := range p.subPackets {
			v := sp.GetValue()
			if v > value {
				value = v
			}
		}
		return value
	case Literal:
		return p.literal
	case GreaterThan:
		if p.subPackets[0].GetValue() > p.subPackets[1].GetValue() {
			return 1
		}
		return 0
	case LessThan:
		if p.subPackets[0].GetValue() < p.subPackets[1].GetValue() {
			return 1
		}
		return 0
	case EqualTo:
		if p.subPackets[0].GetValue() == p.subPackets[1].GetValue() {
			return 1
		}
		return 0
	default:
		fmt.Println("Unhandled Operation:", p.typeId)
		return 0
	}
}

type Decoder struct {
	bits   []uint8
	offset uint
}

func NewDecoder(line string) Decoder {
	bytes, _ := hex.DecodeString(line)
	bits := make([]uint8, 0, len(bytes)*8)
	for _, b := range bytes {
		for i := 7; i >= 0; i-- {
			bits = append(bits, (b>>i)&0x01)
		}
	}
	return Decoder{bits, 0}
}

func (d *Decoder) NextPacket() *Packet {
	if len(d.bits)-int(d.offset) < 8 {
		return nil
	}

	vers := d.GetBits(3)
	typeId := TypeId(d.GetBits(3))

	// This is a literal, so read until the number is full
	if typeId == 4 {
		var val uint = 0
		for done := false; !done; {
			done = d.GetBits(1) == 0
			val <<= 4
			val |= d.GetBits(4)
		}
		return &Packet{vers, typeId, val, nil}
	} else {
		lengthType := d.GetBits(1)
		subPackets := make([]*Packet, 0)
		if lengthType == 0 {
			subpacketBits := d.GetBits(15)
			endOffset := d.offset + subpacketBits
			for d.offset < endOffset {
				subPackets = append(subPackets, d.NextPacket())
			}
			return &Packet{vers, typeId, 0, subPackets}
		} else {
			subpacketCount := d.GetBits(11)
			for i := uint(0); i < subpacketCount; i++ {
				subPackets = append(subPackets, d.NextPacket())
			}
			return &Packet{vers, typeId, 0, subPackets}
		}
	}
}

func (d *Decoder) GetBits(count uint) uint {
	var result uint
	for i := uint(0); i < count; i++ {
		result <<= 1
		result |= uint(d.bits[d.offset+i])
	}
	d.offset += count
	return result
}

func input2slice(fileName string) []string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var text []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		text = append(text, s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	return text
}
