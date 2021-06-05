package main

import (
	"strings"
)

type GoType uint8

const (
	TypeStruct GoType = iota
	TypePointer
	TypeChan
	TypeSlice
	TypeMap
	TypeString
	TypeComplex128
	TypeComplex64
	TypeFloat64
	TypeUint64
	TypeInt64
	TypeUint
	TypeInt
	TypeFloat32
	TypeUint32
	TypeInt32
	TypeRune
	TypeUint16
	TypeInt16
	TypeUint8
	TypeInt8
	TypeByte
	TypeBool
)

func GetGoType(t string) GoType {
	starts := func(s string) bool {
		return strings.HasPrefix(t, s)
	}

	if strings.HasPrefix(t, "*") {
		return TypePointer
	}

	if t == "rune" {
		return TypeRune
	}

	if starts("uint") {
		return uintType(t)
	}

	if starts("int") {
		return intType(t)
	}

	if starts("float") {
		return floatType(t)
	}

	if starts("complex") {
		return complexType(t)
	}

	if t == "string" {
		return TypeString
	}

	if starts("[") {
		return TypeSlice
	}

	if starts("chan") {
		return TypeChan
	}

	if starts("map") {
		return TypeMap
	}

	if t == "bool" {
		return TypeBool
	}

	if t == "byte" {
		return TypeByte
	}

	return TypeStruct
}

func uintType(t string) GoType {
	switch t {
	case "uint8":
		return TypeUint8
	case "uint16":
		return TypeUint16
	case "uint32":
		return TypeUint32
	case "uint64":
		return TypeUint64
	default:
		return TypeUint64
	}
}

func intType(t string) GoType {
	switch t {
	case "int8":
		return TypeInt8
	case "int16":
		return TypeInt16
	case "int32":
		return TypeInt32
	case "int64":
		return TypeInt64
	default:
		return TypeInt
	}
}

func complexType(t string) GoType {
	if t == "complex128" {
		return TypeComplex128
	}
	return TypeComplex64
}

func floatType(t string) GoType {
	if t == "float64" {
		return TypeFloat64
	}
	return TypeFloat32
}
