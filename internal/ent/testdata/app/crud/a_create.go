// Code generated by entc, DO NOT EDIT.

package crud

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ACreate is the builder for creating a A entity.
type ACreate struct {
	config
	mutation *AMutation
	hooks    []Hook
}

// SetA sets the "a" field.
func (a *ACreate) SetA(i int) *ACreate {
	a.mutation.SetA(i)
	return a
}

// Mutation returns the AMutation object of the builder.
func (a *ACreate) Mutation() *AMutation {
	return a.mutation
}

// Save creates the A in the database.
func (a *ACreate) Save(ctx context.Context) (*A, error) {
	var (
		err  error
		node *A
	)
	if len(a.hooks) == 0 {
		if err = a.check(); err != nil {
			return nil, err
		}
		node, err = a.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = a.check(); err != nil {
				return nil, err
			}
			a.mutation = mutation
			node, err = a.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(a.hooks) - 1; i >= 0; i-- {
			mut = a.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, a.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (a *ACreate) SaveX(ctx context.Context) *A {
	v, err := a.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (a *ACreate) check() error {
	if _, ok := a.mutation.A(); !ok {
		return &ValidationError{Name: "a", err: errors.New("crud: missing required field \"a\"")}
	}
	return nil
}

func (a *ACreate) sqlSave(ctx context.Context) (*A, error) {
	_node, _spec := a.createSpec()
	if err := sqlgraph.CreateNode(ctx, a.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (a *ACreate) createSpec() (*A, *sqlgraph.CreateSpec) {
	var (
		_node = &A{config: a.config}
		_spec = &sqlgraph.CreateSpec{
			Table: a.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: a.FieldID,
			},
		}
	)
	if value, ok := a.mutation.A(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: a.FieldA,
		})
		_node.A = value
	}
	return _node, _spec
}

// ACreateBulk is the builder for creating many A entities in bulk.
type ACreateBulk struct {
	config
	builders []*ACreate
}

// Save creates the A entities in the database.
func (ab *ACreateBulk) Save(ctx context.Context) ([]*A, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ab.builders))
	nodes := make([]*A, len(ab.builders))
	mutators := make([]Mutator, len(ab.builders))
	for i := range ab.builders {
		func(i int, root context.Context) {
			builder := ab.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ab.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ab.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ab.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ab *ACreateBulk) SaveX(ctx context.Context) []*A {
	v, err := ab.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
