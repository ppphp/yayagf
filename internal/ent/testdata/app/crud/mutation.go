// Code generated by entc, DO NOT EDIT.

package crud

import (
	"context"
	"fmt"
	"sync"
	"testdata/app/crud/a"
	"testdata/app/crud/predicate"

	"github.com/facebook/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeA = "A"
)

// AMutation represents an operation that mutate the As
// nodes in the graph.
type AMutation struct {
	config
	op            Op
	typ           string
	id            *int
	a             *int
	adda          *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*A, error)
	predicates    []predicate.A
}

var _ ent.Mutation = (*AMutation)(nil)

// aOption allows to manage the mutation configuration using functional options.
type aOption func(*AMutation)

// newAMutation creates new mutation for A.
func newAMutation(c config, op Op, opts ...aOption) *AMutation {
	m := &AMutation{
		config:        c,
		op:            op,
		typ:           TypeA,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAID sets the id field of the mutation.
func withAID(id int) aOption {
	return func(m *AMutation) {
		var (
			err   error
			once  sync.Once
			value *A
		)
		m.oldValue = func(ctx context.Context) (*A, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().A.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withA sets the old A of the mutation.
func withA(node *A) aOption {
	return func(m *AMutation) {
		m.oldValue = func(context.Context) (*A, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("crud: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *AMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetA sets the a field.
func (m *AMutation) SetA(i int) {
	m.a = &i
	m.adda = nil
}

// A returns the a value in the mutation.
func (m *AMutation) A() (r int, exists bool) {
	v := m.a
	if v == nil {
		return
	}
	return *v, true
}

// OldA returns the old a value of the A.
// If the A object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *AMutation) OldA(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldA is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldA requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldA: %w", err)
	}
	return oldValue.A, nil
}

// AddA adds i to a.
func (m *AMutation) AddA(i int) {
	if m.adda != nil {
		*m.adda += i
	} else {
		m.adda = &i
	}
}

// AddedA returns the value that was added to the a field in this mutation.
func (m *AMutation) AddedA() (r int, exists bool) {
	v := m.adda
	if v == nil {
		return
	}
	return *v, true
}

// ResetA reset all changes of the "a" field.
func (m *AMutation) ResetA() {
	m.a = nil
	m.adda = nil
}

// Op returns the operation name.
func (m *AMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (A).
func (m *AMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *AMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.a != nil {
		fields = append(fields, a.FieldA)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *AMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case a.FieldA:
		return m.A()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *AMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case a.FieldA:
		return m.OldA(ctx)
	}
	return nil, fmt.Errorf("unknown A field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *AMutation) SetField(name string, value ent.Value) error {
	switch name {
	case a.FieldA:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetA(v)
		return nil
	}
	return fmt.Errorf("unknown A field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *AMutation) AddedFields() []string {
	var fields []string
	if m.adda != nil {
		fields = append(fields, a.FieldA)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *AMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case a.FieldA:
		return m.AddedA()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *AMutation) AddField(name string, value ent.Value) error {
	switch name {
	case a.FieldA:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddA(v)
		return nil
	}
	return fmt.Errorf("unknown A numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *AMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *AMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *AMutation) ClearField(name string) error {
	return fmt.Errorf("unknown A nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *AMutation) ResetField(name string) error {
	switch name {
	case a.FieldA:
		m.ResetA()
		return nil
	}
	return fmt.Errorf("unknown A field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *AMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *AMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *AMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *AMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *AMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *AMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *AMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown A unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *AMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown A edge %s", name)
}
