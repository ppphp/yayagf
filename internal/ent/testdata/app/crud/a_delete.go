// Code generated by entc, DO NOT EDIT.

package crud

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"test.com/testdata/app/crud/a"
	"test.com/testdata/app/crud/predicate"
)

// ADelete is the builder for deleting a A entity.
type ADelete struct {
	config
	hooks    []Hook
	mutation *AMutation
}

// Where adds a new predicate to the ADelete builder.
func (a *ADelete) Where(ps ...predicate.A) *ADelete {
	a.mutation.predicates = append(a.mutation.predicates, ps...)
	return a
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (a *ADelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(a.hooks) == 0 {
		affected, err = a.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			a.mutation = mutation
			affected, err = a.sqlExec(ctx)
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

// ExecX is like Exec, but panics if an error occurs.
func (a *ADelete) ExecX(ctx context.Context) int {
	n, err := a.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (a *ADelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: a.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: a.FieldID,
			},
		},
	}
	if ps := a.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, a.driver, _spec)
}

// ADeleteOne is the builder for deleting a single A entity.
type ADeleteOne struct {
	a *ADelete
}

// Exec executes the deletion query.
func (ao *ADeleteOne) Exec(ctx context.Context) error {
	n, err := ao.a.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{a.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ao *ADeleteOne) ExecX(ctx context.Context) {
	ao.a.ExecX(ctx)
}
