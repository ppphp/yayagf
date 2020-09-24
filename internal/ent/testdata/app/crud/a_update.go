// Code generated by entc, DO NOT EDIT.

package crud

import (
	"context"
	"fmt"
	"testdata/app/crud/a"
	"testdata/app/crud/predicate"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// AUpdate is the builder for updating A entities.
type AUpdate struct {
	config
	hooks      []Hook
	mutation   *AMutation
	predicates []predicate.A
}

// Where adds a new predicate for the builder.
func (a *AUpdate) Where(ps ...predicate.A) *AUpdate {
	a.predicates = append(a.predicates, ps...)
	return a
}

// SetA sets the a field.
func (a *AUpdate) SetA(i int) *AUpdate {
	a.mutation.ResetA()
	a.mutation.SetA(i)
	return a
}

// AddA adds i to a.
func (a *AUpdate) AddA(i int) *AUpdate {
	a.mutation.AddA(i)
	return a
}

// Mutation returns the AMutation object of the builder.
func (a *AUpdate) Mutation() *AMutation {
	return a.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (a *AUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(a.hooks) == 0 {
		affected, err = a.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			a.mutation = mutation
			affected, err = a.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(a.hooks) - 1; i >= 0; i-- {
			mut = a.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, a.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (a *AUpdate) SaveX(ctx context.Context) int {
	affected, err := a.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (a *AUpdate) Exec(ctx context.Context) error {
	_, err := a.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (a *AUpdate) ExecX(ctx context.Context) {
	if err := a.Exec(ctx); err != nil {
		panic(err)
	}
}

func (a *AUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   a.Table,
			Columns: a.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: a.FieldID,
			},
		},
	}
	if ps := a.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := a.mutation.A(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: a.FieldA,
		})
	}
	if value, ok := a.mutation.AddedA(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: a.FieldA,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, a.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{a.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AUpdateOne is the builder for updating a single A entity.
type AUpdateOne struct {
	config
	hooks    []Hook
	mutation *AMutation
}

// SetA sets the a field.
func (ao *AUpdateOne) SetA(i int) *AUpdateOne {
	ao.mutation.ResetA()
	ao.mutation.SetA(i)
	return ao
}

// AddA adds i to a.
func (ao *AUpdateOne) AddA(i int) *AUpdateOne {
	ao.mutation.AddA(i)
	return ao
}

// Mutation returns the AMutation object of the builder.
func (ao *AUpdateOne) Mutation() *AMutation {
	return ao.mutation
}

// Save executes the query and returns the updated entity.
func (ao *AUpdateOne) Save(ctx context.Context) (*A, error) {
	var (
		err  error
		node *A
	)
	if len(ao.hooks) == 0 {
		node, err = ao.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ao.mutation = mutation
			node, err = ao.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ao.hooks) - 1; i >= 0; i-- {
			mut = ao.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ao.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ao *AUpdateOne) SaveX(ctx context.Context) *A {
	node, err := ao.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ao *AUpdateOne) Exec(ctx context.Context) error {
	_, err := ao.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ao *AUpdateOne) ExecX(ctx context.Context) {
	if err := ao.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ao *AUpdateOne) sqlSave(ctx context.Context) (_node *A, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   a.Table,
			Columns: a.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: a.FieldID,
			},
		},
	}
	id, ok := ao.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing A.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ao.mutation.A(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: a.FieldA,
		})
	}
	if value, ok := ao.mutation.AddedA(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: a.FieldA,
		})
	}
	_node = &A{config: ao.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ao.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{a.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
