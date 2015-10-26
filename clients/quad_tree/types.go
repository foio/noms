// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.TypeRef{
		types.MakeStructTypeRef("Node",
			[]types.Field{
				types.Field{"Geoposition", types.MakeTypeRef(ref.Parse("sha1-fb09d21d144c518467325465327d46489cff7c47"), 0), false},
				types.Field{"Reference", types.MakeCompoundTypeRef(types.RefKind, types.MakePrimitiveTypeRef(types.ValueKind)), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("QuadTree",
			[]types.Field{
				types.Field{"Nodes", types.MakeCompoundTypeRef(types.ListKind, types.MakeTypeRef(ref.Ref{}, 0)), false},
				types.Field{"Tiles", types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef(ref.Ref{}, 1)), false},
				types.Field{"Depth", types.MakePrimitiveTypeRef(types.UInt8Kind), false},
				types.Field{"NumDescendents", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
				types.Field{"Path", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Georectangle", types.MakeTypeRef(ref.Parse("sha1-fb09d21d144c518467325465327d46489cff7c47"), 1), false},
			},
			types.Choices{},
		),
		types.MakeStructTypeRef("SQuadTree",
			[]types.Field{
				types.Field{"Nodes", types.MakeCompoundTypeRef(types.ListKind, types.MakeCompoundTypeRef(types.RefKind, types.MakePrimitiveTypeRef(types.ValueKind))), false},
				types.Field{"Tiles", types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(ref.Ref{}, 2))), false},
				types.Field{"Depth", types.MakePrimitiveTypeRef(types.UInt8Kind), false},
				types.Field{"NumDescendents", types.MakePrimitiveTypeRef(types.UInt32Kind), false},
				types.Field{"Path", types.MakePrimitiveTypeRef(types.StringKind), false},
				types.Field{"Georectangle", types.MakeTypeRef(ref.Parse("sha1-fb09d21d144c518467325465327d46489cff7c47"), 1), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-fb09d21d144c518467325465327d46489cff7c47"),
	})
	__mainPackageInFile_types_CachedRef = types.RegisterPackage(&p)
}

// Node

type Node struct {
	m   types.Map
	ref *ref.Ref
}

func NewNode() Node {
	return Node{types.NewMap(
		types.NewString("Geoposition"), NewGeoposition(),
		types.NewString("Reference"), NewRefOfValue(ref.Ref{}),
	), &ref.Ref{}}
}

type NodeDef struct {
	Geoposition GeopositionDef
	Reference   ref.Ref
}

func (def NodeDef) New() Node {
	return Node{
		types.NewMap(
			types.NewString("Geoposition"), def.Geoposition.New(),
			types.NewString("Reference"), NewRefOfValue(def.Reference),
		), &ref.Ref{}}
}

func (s Node) Def() (d NodeDef) {
	d.Geoposition = s.m.Get(types.NewString("Geoposition")).(Geoposition).Def()
	d.Reference = s.m.Get(types.NewString("Reference")).Ref()
	return
}

var __typeRefForNode types.TypeRef

func (m Node) TypeRef() types.TypeRef {
	return __typeRefForNode
}

func init() {
	__typeRefForNode = types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0)
	types.RegisterFromValFunction(__typeRefForNode, func(v types.Value) types.Value {
		return NodeFromVal(v)
	})
}

func NodeFromVal(val types.Value) Node {
	// TODO: Do we still need FromVal?
	if val, ok := val.(Node); ok {
		return val
	}
	// TODO: Validate here
	return Node{val.(types.Map), &ref.Ref{}}
}

func (s Node) InternalImplementation() types.Map {
	return s.m
}

func (s Node) Equals(other types.Value) bool {
	if other, ok := other.(Node); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s Node) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Node) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s Node) Geoposition() Geoposition {
	return s.m.Get(types.NewString("Geoposition")).(Geoposition)
}

func (s Node) SetGeoposition(val Geoposition) Node {
	return Node{s.m.Set(types.NewString("Geoposition"), val), &ref.Ref{}}
}

func (s Node) Reference() RefOfValue {
	return s.m.Get(types.NewString("Reference")).(RefOfValue)
}

func (s Node) SetReference(val RefOfValue) Node {
	return Node{s.m.Set(types.NewString("Reference"), val), &ref.Ref{}}
}

// QuadTree

type QuadTree struct {
	m   types.Map
	ref *ref.Ref
}

func NewQuadTree() QuadTree {
	return QuadTree{types.NewMap(
		types.NewString("Nodes"), NewListOfNode(),
		types.NewString("Tiles"), NewMapOfStringToQuadTree(),
		types.NewString("Depth"), types.UInt8(0),
		types.NewString("NumDescendents"), types.UInt32(0),
		types.NewString("Path"), types.NewString(""),
		types.NewString("Georectangle"), NewGeorectangle(),
	), &ref.Ref{}}
}

type QuadTreeDef struct {
	Nodes          ListOfNodeDef
	Tiles          MapOfStringToQuadTreeDef
	Depth          uint8
	NumDescendents uint32
	Path           string
	Georectangle   GeorectangleDef
}

func (def QuadTreeDef) New() QuadTree {
	return QuadTree{
		types.NewMap(
			types.NewString("Nodes"), def.Nodes.New(),
			types.NewString("Tiles"), def.Tiles.New(),
			types.NewString("Depth"), types.UInt8(def.Depth),
			types.NewString("NumDescendents"), types.UInt32(def.NumDescendents),
			types.NewString("Path"), types.NewString(def.Path),
			types.NewString("Georectangle"), def.Georectangle.New(),
		), &ref.Ref{}}
}

func (s QuadTree) Def() (d QuadTreeDef) {
	d.Nodes = s.m.Get(types.NewString("Nodes")).(ListOfNode).Def()
	d.Tiles = s.m.Get(types.NewString("Tiles")).(MapOfStringToQuadTree).Def()
	d.Depth = uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
	d.NumDescendents = uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
	d.Path = s.m.Get(types.NewString("Path")).(types.String).String()
	d.Georectangle = s.m.Get(types.NewString("Georectangle")).(Georectangle).Def()
	return
}

var __typeRefForQuadTree types.TypeRef

func (m QuadTree) TypeRef() types.TypeRef {
	return __typeRefForQuadTree
}

func init() {
	__typeRefForQuadTree = types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 1)
	types.RegisterFromValFunction(__typeRefForQuadTree, func(v types.Value) types.Value {
		return QuadTreeFromVal(v)
	})
}

func QuadTreeFromVal(val types.Value) QuadTree {
	// TODO: Do we still need FromVal?
	if val, ok := val.(QuadTree); ok {
		return val
	}
	// TODO: Validate here
	return QuadTree{val.(types.Map), &ref.Ref{}}
}

func (s QuadTree) InternalImplementation() types.Map {
	return s.m
}

func (s QuadTree) Equals(other types.Value) bool {
	if other, ok := other.(QuadTree); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s QuadTree) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s QuadTree) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s QuadTree) Nodes() ListOfNode {
	return s.m.Get(types.NewString("Nodes")).(ListOfNode)
}

func (s QuadTree) SetNodes(val ListOfNode) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Nodes"), val), &ref.Ref{}}
}

func (s QuadTree) Tiles() MapOfStringToQuadTree {
	return s.m.Get(types.NewString("Tiles")).(MapOfStringToQuadTree)
}

func (s QuadTree) SetTiles(val MapOfStringToQuadTree) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Tiles"), val), &ref.Ref{}}
}

func (s QuadTree) Depth() uint8 {
	return uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
}

func (s QuadTree) SetDepth(val uint8) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Depth"), types.UInt8(val)), &ref.Ref{}}
}

func (s QuadTree) NumDescendents() uint32 {
	return uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
}

func (s QuadTree) SetNumDescendents(val uint32) QuadTree {
	return QuadTree{s.m.Set(types.NewString("NumDescendents"), types.UInt32(val)), &ref.Ref{}}
}

func (s QuadTree) Path() string {
	return s.m.Get(types.NewString("Path")).(types.String).String()
}

func (s QuadTree) SetPath(val string) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Path"), types.NewString(val)), &ref.Ref{}}
}

func (s QuadTree) Georectangle() Georectangle {
	return s.m.Get(types.NewString("Georectangle")).(Georectangle)
}

func (s QuadTree) SetGeorectangle(val Georectangle) QuadTree {
	return QuadTree{s.m.Set(types.NewString("Georectangle"), val), &ref.Ref{}}
}

// SQuadTree

type SQuadTree struct {
	m   types.Map
	ref *ref.Ref
}

func NewSQuadTree() SQuadTree {
	return SQuadTree{types.NewMap(
		types.NewString("Nodes"), NewListOfRefOfValue(),
		types.NewString("Tiles"), NewMapOfStringToRefOfSQuadTree(),
		types.NewString("Depth"), types.UInt8(0),
		types.NewString("NumDescendents"), types.UInt32(0),
		types.NewString("Path"), types.NewString(""),
		types.NewString("Georectangle"), NewGeorectangle(),
	), &ref.Ref{}}
}

type SQuadTreeDef struct {
	Nodes          ListOfRefOfValueDef
	Tiles          MapOfStringToRefOfSQuadTreeDef
	Depth          uint8
	NumDescendents uint32
	Path           string
	Georectangle   GeorectangleDef
}

func (def SQuadTreeDef) New() SQuadTree {
	return SQuadTree{
		types.NewMap(
			types.NewString("Nodes"), def.Nodes.New(),
			types.NewString("Tiles"), def.Tiles.New(),
			types.NewString("Depth"), types.UInt8(def.Depth),
			types.NewString("NumDescendents"), types.UInt32(def.NumDescendents),
			types.NewString("Path"), types.NewString(def.Path),
			types.NewString("Georectangle"), def.Georectangle.New(),
		), &ref.Ref{}}
}

func (s SQuadTree) Def() (d SQuadTreeDef) {
	d.Nodes = s.m.Get(types.NewString("Nodes")).(ListOfRefOfValue).Def()
	d.Tiles = s.m.Get(types.NewString("Tiles")).(MapOfStringToRefOfSQuadTree).Def()
	d.Depth = uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
	d.NumDescendents = uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
	d.Path = s.m.Get(types.NewString("Path")).(types.String).String()
	d.Georectangle = s.m.Get(types.NewString("Georectangle")).(Georectangle).Def()
	return
}

var __typeRefForSQuadTree types.TypeRef

func (m SQuadTree) TypeRef() types.TypeRef {
	return __typeRefForSQuadTree
}

func init() {
	__typeRefForSQuadTree = types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 2)
	types.RegisterFromValFunction(__typeRefForSQuadTree, func(v types.Value) types.Value {
		return SQuadTreeFromVal(v)
	})
}

func SQuadTreeFromVal(val types.Value) SQuadTree {
	// TODO: Do we still need FromVal?
	if val, ok := val.(SQuadTree); ok {
		return val
	}
	// TODO: Validate here
	return SQuadTree{val.(types.Map), &ref.Ref{}}
}

func (s SQuadTree) InternalImplementation() types.Map {
	return s.m
}

func (s SQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(SQuadTree); ok {
		return s.Ref() == other.Ref()
	}
	return false
}

func (s SQuadTree) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s SQuadTree) Chunks() (futures []types.Future) {
	futures = append(futures, s.TypeRef().Chunks()...)
	futures = append(futures, s.m.Chunks()...)
	return
}

func (s SQuadTree) Nodes() ListOfRefOfValue {
	return s.m.Get(types.NewString("Nodes")).(ListOfRefOfValue)
}

func (s SQuadTree) SetNodes(val ListOfRefOfValue) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Nodes"), val), &ref.Ref{}}
}

func (s SQuadTree) Tiles() MapOfStringToRefOfSQuadTree {
	return s.m.Get(types.NewString("Tiles")).(MapOfStringToRefOfSQuadTree)
}

func (s SQuadTree) SetTiles(val MapOfStringToRefOfSQuadTree) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Tiles"), val), &ref.Ref{}}
}

func (s SQuadTree) Depth() uint8 {
	return uint8(s.m.Get(types.NewString("Depth")).(types.UInt8))
}

func (s SQuadTree) SetDepth(val uint8) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Depth"), types.UInt8(val)), &ref.Ref{}}
}

func (s SQuadTree) NumDescendents() uint32 {
	return uint32(s.m.Get(types.NewString("NumDescendents")).(types.UInt32))
}

func (s SQuadTree) SetNumDescendents(val uint32) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("NumDescendents"), types.UInt32(val)), &ref.Ref{}}
}

func (s SQuadTree) Path() string {
	return s.m.Get(types.NewString("Path")).(types.String).String()
}

func (s SQuadTree) SetPath(val string) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Path"), types.NewString(val)), &ref.Ref{}}
}

func (s SQuadTree) Georectangle() Georectangle {
	return s.m.Get(types.NewString("Georectangle")).(Georectangle)
}

func (s SQuadTree) SetGeorectangle(val Georectangle) SQuadTree {
	return SQuadTree{s.m.Set(types.NewString("Georectangle"), val), &ref.Ref{}}
}

// RefOfValue

type RefOfValue struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfValue(target ref.Ref) RefOfValue {
	return RefOfValue{target, &ref.Ref{}}
}

func (r RefOfValue) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfValue) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfValue) Equals(other types.Value) bool {
	if other, ok := other.(RefOfValue); ok {
		return r.Ref() == other.Ref()
	}
	return false
}

func (r RefOfValue) Chunks() []types.Future {
	return r.TypeRef().Chunks()
}

func RefOfValueFromVal(val types.Value) RefOfValue {
	// TODO: Do we still need FromVal?
	if val, ok := val.(RefOfValue); ok {
		return val
	}
	return NewRefOfValue(val.(types.Ref).TargetRef())
}

// A Noms Value that describes RefOfValue.
var __typeRefForRefOfValue types.TypeRef

func (m RefOfValue) TypeRef() types.TypeRef {
	return __typeRefForRefOfValue
}

func init() {
	__typeRefForRefOfValue = types.MakeCompoundTypeRef(types.RefKind, types.MakePrimitiveTypeRef(types.ValueKind))
	types.RegisterFromValFunction(__typeRefForRefOfValue, func(v types.Value) types.Value {
		return RefOfValueFromVal(v)
	})
}

func (r RefOfValue) TargetValue(cs chunks.ChunkSource) types.Value {
	return types.ReadValue(r.target, cs)
}

func (r RefOfValue) SetTargetValue(val types.Value, cs chunks.ChunkSink) RefOfValue {
	return NewRefOfValue(types.WriteValue(val, cs))
}

// ListOfNode

type ListOfNode struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfNode() ListOfNode {
	return ListOfNode{types.NewList(), &ref.Ref{}}
}

type ListOfNodeDef []NodeDef

func (def ListOfNodeDef) New() ListOfNode {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New()
	}
	return ListOfNode{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfNode) Def() ListOfNodeDef {
	d := make([]NodeDef, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).(Node).Def()
	}
	return d
}

func ListOfNodeFromVal(val types.Value) ListOfNode {
	// TODO: Do we still need FromVal?
	if val, ok := val.(ListOfNode); ok {
		return val
	}
	// TODO: Validate here
	return ListOfNode{val.(types.List), &ref.Ref{}}
}

func (l ListOfNode) InternalImplementation() types.List {
	return l.l
}

func (l ListOfNode) Equals(other types.Value) bool {
	if other, ok := other.(ListOfNode); ok {
		return l.Ref() == other.Ref()
	}
	return false
}

func (l ListOfNode) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfNode) Chunks() (futures []types.Future) {
	futures = append(futures, l.TypeRef().Chunks()...)
	futures = append(futures, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfNode.
var __typeRefForListOfNode types.TypeRef

func (m ListOfNode) TypeRef() types.TypeRef {
	return __typeRefForListOfNode
}

func init() {
	__typeRefForListOfNode = types.MakeCompoundTypeRef(types.ListKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 0))
	types.RegisterFromValFunction(__typeRefForListOfNode, func(v types.Value) types.Value {
		return ListOfNodeFromVal(v)
	})
}

func (l ListOfNode) Len() uint64 {
	return l.l.Len()
}

func (l ListOfNode) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfNode) Get(i uint64) Node {
	return l.l.Get(i).(Node)
}

func (l ListOfNode) Slice(idx uint64, end uint64) ListOfNode {
	return ListOfNode{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfNode) Set(i uint64, val Node) ListOfNode {
	return ListOfNode{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfNode) Append(v ...Node) ListOfNode {
	return ListOfNode{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfNode) Insert(idx uint64, v ...Node) ListOfNode {
	return ListOfNode{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfNode) Remove(idx uint64, end uint64) ListOfNode {
	return ListOfNode{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfNode) RemoveAt(idx uint64) ListOfNode {
	return ListOfNode{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfNode) fromElemSlice(p []Node) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfNodeIterCallback func(v Node, i uint64) (stop bool)

func (l ListOfNode) Iter(cb ListOfNodeIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(Node), i)
	})
}

type ListOfNodeIterAllCallback func(v Node, i uint64)

func (l ListOfNode) IterAll(cb ListOfNodeIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(Node), i)
	})
}

type ListOfNodeFilterCallback func(v Node, i uint64) (keep bool)

func (l ListOfNode) Filter(cb ListOfNodeFilterCallback) ListOfNode {
	nl := NewListOfNode()
	l.IterAll(func(v Node, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// MapOfStringToQuadTree

type MapOfStringToQuadTree struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToQuadTree() MapOfStringToQuadTree {
	return MapOfStringToQuadTree{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToQuadTreeDef map[string]QuadTreeDef

func (def MapOfStringToQuadTreeDef) New() MapOfStringToQuadTree {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New())
	}
	return MapOfStringToQuadTree{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToQuadTree) Def() MapOfStringToQuadTreeDef {
	def := make(map[string]QuadTreeDef)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(QuadTree).Def()
		return false
	})
	return def
}

func MapOfStringToQuadTreeFromVal(val types.Value) MapOfStringToQuadTree {
	// TODO: Do we still need FromVal?
	if val, ok := val.(MapOfStringToQuadTree); ok {
		return val
	}
	// TODO: Validate here
	return MapOfStringToQuadTree{val.(types.Map), &ref.Ref{}}
}

func (m MapOfStringToQuadTree) InternalImplementation() types.Map {
	return m.m
}

func (m MapOfStringToQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToQuadTree); ok {
		return m.Ref() == other.Ref()
	}
	return false
}

func (m MapOfStringToQuadTree) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToQuadTree) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToQuadTree.
var __typeRefForMapOfStringToQuadTree types.TypeRef

func (m MapOfStringToQuadTree) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToQuadTree
}

func init() {
	__typeRefForMapOfStringToQuadTree = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 1))
	types.RegisterFromValFunction(__typeRefForMapOfStringToQuadTree, func(v types.Value) types.Value {
		return MapOfStringToQuadTreeFromVal(v)
	})
}

func (m MapOfStringToQuadTree) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToQuadTree) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToQuadTree) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToQuadTree) Get(p string) QuadTree {
	return m.m.Get(types.NewString(p)).(QuadTree)
}

func (m MapOfStringToQuadTree) MaybeGet(p string) (QuadTree, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewQuadTree(), false
	}
	return v.(QuadTree), ok
}

func (m MapOfStringToQuadTree) Set(k string, v QuadTree) MapOfStringToQuadTree {
	return MapOfStringToQuadTree{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToQuadTree) Remove(p string) MapOfStringToQuadTree {
	return MapOfStringToQuadTree{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToQuadTreeIterCallback func(k string, v QuadTree) (stop bool)

func (m MapOfStringToQuadTree) Iter(cb MapOfStringToQuadTreeIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(QuadTree))
	})
}

type MapOfStringToQuadTreeIterAllCallback func(k string, v QuadTree)

func (m MapOfStringToQuadTree) IterAll(cb MapOfStringToQuadTreeIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(QuadTree))
	})
}

type MapOfStringToQuadTreeFilterCallback func(k string, v QuadTree) (keep bool)

func (m MapOfStringToQuadTree) Filter(cb MapOfStringToQuadTreeFilterCallback) MapOfStringToQuadTree {
	nm := NewMapOfStringToQuadTree()
	m.IterAll(func(k string, v QuadTree) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// ListOfRefOfValue

type ListOfRefOfValue struct {
	l   types.List
	ref *ref.Ref
}

func NewListOfRefOfValue() ListOfRefOfValue {
	return ListOfRefOfValue{types.NewList(), &ref.Ref{}}
}

type ListOfRefOfValueDef []ref.Ref

func (def ListOfRefOfValueDef) New() ListOfRefOfValue {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = NewRefOfValue(d)
	}
	return ListOfRefOfValue{types.NewList(l...), &ref.Ref{}}
}

func (l ListOfRefOfValue) Def() ListOfRefOfValueDef {
	d := make([]ref.Ref, l.Len())
	for i := uint64(0); i < l.Len(); i++ {
		d[i] = l.l.Get(i).Ref()
	}
	return d
}

func ListOfRefOfValueFromVal(val types.Value) ListOfRefOfValue {
	// TODO: Do we still need FromVal?
	if val, ok := val.(ListOfRefOfValue); ok {
		return val
	}
	// TODO: Validate here
	return ListOfRefOfValue{val.(types.List), &ref.Ref{}}
}

func (l ListOfRefOfValue) InternalImplementation() types.List {
	return l.l
}

func (l ListOfRefOfValue) Equals(other types.Value) bool {
	if other, ok := other.(ListOfRefOfValue); ok {
		return l.Ref() == other.Ref()
	}
	return false
}

func (l ListOfRefOfValue) Ref() ref.Ref {
	return types.EnsureRef(l.ref, l)
}

func (l ListOfRefOfValue) Chunks() (futures []types.Future) {
	futures = append(futures, l.TypeRef().Chunks()...)
	futures = append(futures, l.l.Chunks()...)
	return
}

// A Noms Value that describes ListOfRefOfValue.
var __typeRefForListOfRefOfValue types.TypeRef

func (m ListOfRefOfValue) TypeRef() types.TypeRef {
	return __typeRefForListOfRefOfValue
}

func init() {
	__typeRefForListOfRefOfValue = types.MakeCompoundTypeRef(types.ListKind, types.MakeCompoundTypeRef(types.RefKind, types.MakePrimitiveTypeRef(types.ValueKind)))
	types.RegisterFromValFunction(__typeRefForListOfRefOfValue, func(v types.Value) types.Value {
		return ListOfRefOfValueFromVal(v)
	})
}

func (l ListOfRefOfValue) Len() uint64 {
	return l.l.Len()
}

func (l ListOfRefOfValue) Empty() bool {
	return l.Len() == uint64(0)
}

func (l ListOfRefOfValue) Get(i uint64) RefOfValue {
	return l.l.Get(i).(RefOfValue)
}

func (l ListOfRefOfValue) Slice(idx uint64, end uint64) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Slice(idx, end), &ref.Ref{}}
}

func (l ListOfRefOfValue) Set(i uint64, val RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Set(i, val), &ref.Ref{}}
}

func (l ListOfRefOfValue) Append(v ...RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Append(l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfRefOfValue) Insert(idx uint64, v ...RefOfValue) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Insert(idx, l.fromElemSlice(v)...), &ref.Ref{}}
}

func (l ListOfRefOfValue) Remove(idx uint64, end uint64) ListOfRefOfValue {
	return ListOfRefOfValue{l.l.Remove(idx, end), &ref.Ref{}}
}

func (l ListOfRefOfValue) RemoveAt(idx uint64) ListOfRefOfValue {
	return ListOfRefOfValue{(l.l.RemoveAt(idx)), &ref.Ref{}}
}

func (l ListOfRefOfValue) fromElemSlice(p []RefOfValue) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v
	}
	return r
}

type ListOfRefOfValueIterCallback func(v RefOfValue, i uint64) (stop bool)

func (l ListOfRefOfValue) Iter(cb ListOfRefOfValueIterCallback) {
	l.l.Iter(func(v types.Value, i uint64) bool {
		return cb(v.(RefOfValue), i)
	})
}

type ListOfRefOfValueIterAllCallback func(v RefOfValue, i uint64)

func (l ListOfRefOfValue) IterAll(cb ListOfRefOfValueIterAllCallback) {
	l.l.IterAll(func(v types.Value, i uint64) {
		cb(v.(RefOfValue), i)
	})
}

type ListOfRefOfValueFilterCallback func(v RefOfValue, i uint64) (keep bool)

func (l ListOfRefOfValue) Filter(cb ListOfRefOfValueFilterCallback) ListOfRefOfValue {
	nl := NewListOfRefOfValue()
	l.IterAll(func(v RefOfValue, i uint64) {
		if cb(v, i) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// MapOfStringToRefOfSQuadTree

type MapOfStringToRefOfSQuadTree struct {
	m   types.Map
	ref *ref.Ref
}

func NewMapOfStringToRefOfSQuadTree() MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{types.NewMap(), &ref.Ref{}}
}

type MapOfStringToRefOfSQuadTreeDef map[string]ref.Ref

func (def MapOfStringToRefOfSQuadTreeDef) New() MapOfStringToRefOfSQuadTree {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), NewRefOfSQuadTree(v))
	}
	return MapOfStringToRefOfSQuadTree{types.NewMap(kv...), &ref.Ref{}}
}

func (m MapOfStringToRefOfSQuadTree) Def() MapOfStringToRefOfSQuadTreeDef {
	def := make(map[string]ref.Ref)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.Ref()
		return false
	})
	return def
}

func MapOfStringToRefOfSQuadTreeFromVal(val types.Value) MapOfStringToRefOfSQuadTree {
	// TODO: Do we still need FromVal?
	if val, ok := val.(MapOfStringToRefOfSQuadTree); ok {
		return val
	}
	// TODO: Validate here
	return MapOfStringToRefOfSQuadTree{val.(types.Map), &ref.Ref{}}
}

func (m MapOfStringToRefOfSQuadTree) InternalImplementation() types.Map {
	return m.m
}

func (m MapOfStringToRefOfSQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(MapOfStringToRefOfSQuadTree); ok {
		return m.Ref() == other.Ref()
	}
	return false
}

func (m MapOfStringToRefOfSQuadTree) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToRefOfSQuadTree) Chunks() (futures []types.Future) {
	futures = append(futures, m.TypeRef().Chunks()...)
	futures = append(futures, m.m.Chunks()...)
	return
}

// A Noms Value that describes MapOfStringToRefOfSQuadTree.
var __typeRefForMapOfStringToRefOfSQuadTree types.TypeRef

func (m MapOfStringToRefOfSQuadTree) TypeRef() types.TypeRef {
	return __typeRefForMapOfStringToRefOfSQuadTree
}

func init() {
	__typeRefForMapOfStringToRefOfSQuadTree = types.MakeCompoundTypeRef(types.MapKind, types.MakePrimitiveTypeRef(types.StringKind), types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 2)))
	types.RegisterFromValFunction(__typeRefForMapOfStringToRefOfSQuadTree, func(v types.Value) types.Value {
		return MapOfStringToRefOfSQuadTreeFromVal(v)
	})
}

func (m MapOfStringToRefOfSQuadTree) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToRefOfSQuadTree) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToRefOfSQuadTree) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToRefOfSQuadTree) Get(p string) RefOfSQuadTree {
	return m.m.Get(types.NewString(p)).(RefOfSQuadTree)
}

func (m MapOfStringToRefOfSQuadTree) MaybeGet(p string) (RefOfSQuadTree, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewRefOfSQuadTree(ref.Ref{}), false
	}
	return v.(RefOfSQuadTree), ok
}

func (m MapOfStringToRefOfSQuadTree) Set(k string, v RefOfSQuadTree) MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{m.m.Set(types.NewString(k), v), &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToRefOfSQuadTree) Remove(p string) MapOfStringToRefOfSQuadTree {
	return MapOfStringToRefOfSQuadTree{m.m.Remove(types.NewString(p)), &ref.Ref{}}
}

type MapOfStringToRefOfSQuadTreeIterCallback func(k string, v RefOfSQuadTree) (stop bool)

func (m MapOfStringToRefOfSQuadTree) Iter(cb MapOfStringToRefOfSQuadTreeIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfSQuadTree))
	})
}

type MapOfStringToRefOfSQuadTreeIterAllCallback func(k string, v RefOfSQuadTree)

func (m MapOfStringToRefOfSQuadTree) IterAll(cb MapOfStringToRefOfSQuadTreeIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfSQuadTree))
	})
}

type MapOfStringToRefOfSQuadTreeFilterCallback func(k string, v RefOfSQuadTree) (keep bool)

func (m MapOfStringToRefOfSQuadTree) Filter(cb MapOfStringToRefOfSQuadTreeFilterCallback) MapOfStringToRefOfSQuadTree {
	nm := NewMapOfStringToRefOfSQuadTree()
	m.IterAll(func(k string, v RefOfSQuadTree) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// RefOfSQuadTree

type RefOfSQuadTree struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfSQuadTree(target ref.Ref) RefOfSQuadTree {
	return RefOfSQuadTree{target, &ref.Ref{}}
}

func (r RefOfSQuadTree) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfSQuadTree) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfSQuadTree) Equals(other types.Value) bool {
	if other, ok := other.(RefOfSQuadTree); ok {
		return r.Ref() == other.Ref()
	}
	return false
}

func (r RefOfSQuadTree) Chunks() []types.Future {
	return r.TypeRef().Chunks()
}

func RefOfSQuadTreeFromVal(val types.Value) RefOfSQuadTree {
	// TODO: Do we still need FromVal?
	if val, ok := val.(RefOfSQuadTree); ok {
		return val
	}
	return NewRefOfSQuadTree(val.(types.Ref).TargetRef())
}

// A Noms Value that describes RefOfSQuadTree.
var __typeRefForRefOfSQuadTree types.TypeRef

func (m RefOfSQuadTree) TypeRef() types.TypeRef {
	return __typeRefForRefOfSQuadTree
}

func init() {
	__typeRefForRefOfSQuadTree = types.MakeCompoundTypeRef(types.RefKind, types.MakeTypeRef(__mainPackageInFile_types_CachedRef, 2))
	types.RegisterFromValFunction(__typeRefForRefOfSQuadTree, func(v types.Value) types.Value {
		return RefOfSQuadTreeFromVal(v)
	})
}

func (r RefOfSQuadTree) TargetValue(cs chunks.ChunkSource) SQuadTree {
	return types.ReadValue(r.target, cs).(SQuadTree)
}

func (r RefOfSQuadTree) SetTargetValue(val SQuadTree, cs chunks.ChunkSink) RefOfSQuadTree {
	return NewRefOfSQuadTree(types.WriteValue(val, cs))
}
