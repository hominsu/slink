// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/hominsu/slink/app/slink/service/internal/data/ent/predicate"
	"github.com/hominsu/slink/app/slink/service/internal/data/ent/shortlink"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeShortLink = "ShortLink"
)

// ShortLinkMutation represents an operation that mutates the ShortLink nodes in the graph.
type ShortLinkMutation struct {
	config
	op            Op
	typ           string
	id            *int
	created_at    *time.Time
	updated_at    *time.Time
	key           *string
	link          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*ShortLink, error)
	predicates    []predicate.ShortLink
}

var _ ent.Mutation = (*ShortLinkMutation)(nil)

// shortlinkOption allows management of the mutation configuration using functional options.
type shortlinkOption func(*ShortLinkMutation)

// newShortLinkMutation creates new mutation for the ShortLink entity.
func newShortLinkMutation(c config, op Op, opts ...shortlinkOption) *ShortLinkMutation {
	m := &ShortLinkMutation{
		config:        c,
		op:            op,
		typ:           TypeShortLink,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withShortLinkID sets the ID field of the mutation.
func withShortLinkID(id int) shortlinkOption {
	return func(m *ShortLinkMutation) {
		var (
			err   error
			once  sync.Once
			value *ShortLink
		)
		m.oldValue = func(ctx context.Context) (*ShortLink, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().ShortLink.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withShortLink sets the old ShortLink of the mutation.
func withShortLink(node *ShortLink) shortlinkOption {
	return func(m *ShortLinkMutation) {
		m.oldValue = func(context.Context) (*ShortLink, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ShortLinkMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ShortLinkMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ShortLinkMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *ShortLinkMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().ShortLink.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *ShortLinkMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *ShortLinkMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the ShortLink entity.
// If the ShortLink object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ShortLinkMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *ShortLinkMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *ShortLinkMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *ShortLinkMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the ShortLink entity.
// If the ShortLink object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ShortLinkMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *ShortLinkMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetKey sets the "key" field.
func (m *ShortLinkMutation) SetKey(s string) {
	m.key = &s
}

// Key returns the value of the "key" field in the mutation.
func (m *ShortLinkMutation) Key() (r string, exists bool) {
	v := m.key
	if v == nil {
		return
	}
	return *v, true
}

// OldKey returns the old "key" field's value of the ShortLink entity.
// If the ShortLink object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ShortLinkMutation) OldKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKey: %w", err)
	}
	return oldValue.Key, nil
}

// ResetKey resets all changes to the "key" field.
func (m *ShortLinkMutation) ResetKey() {
	m.key = nil
}

// SetLink sets the "link" field.
func (m *ShortLinkMutation) SetLink(s string) {
	m.link = &s
}

// Link returns the value of the "link" field in the mutation.
func (m *ShortLinkMutation) Link() (r string, exists bool) {
	v := m.link
	if v == nil {
		return
	}
	return *v, true
}

// OldLink returns the old "link" field's value of the ShortLink entity.
// If the ShortLink object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ShortLinkMutation) OldLink(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLink is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLink requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLink: %w", err)
	}
	return oldValue.Link, nil
}

// ResetLink resets all changes to the "link" field.
func (m *ShortLinkMutation) ResetLink() {
	m.link = nil
}

// Where appends a list predicates to the ShortLinkMutation builder.
func (m *ShortLinkMutation) Where(ps ...predicate.ShortLink) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the ShortLinkMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *ShortLinkMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.ShortLink, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *ShortLinkMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *ShortLinkMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (ShortLink).
func (m *ShortLinkMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ShortLinkMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.created_at != nil {
		fields = append(fields, shortlink.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, shortlink.FieldUpdatedAt)
	}
	if m.key != nil {
		fields = append(fields, shortlink.FieldKey)
	}
	if m.link != nil {
		fields = append(fields, shortlink.FieldLink)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ShortLinkMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case shortlink.FieldCreatedAt:
		return m.CreatedAt()
	case shortlink.FieldUpdatedAt:
		return m.UpdatedAt()
	case shortlink.FieldKey:
		return m.Key()
	case shortlink.FieldLink:
		return m.Link()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ShortLinkMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case shortlink.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case shortlink.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case shortlink.FieldKey:
		return m.OldKey(ctx)
	case shortlink.FieldLink:
		return m.OldLink(ctx)
	}
	return nil, fmt.Errorf("unknown ShortLink field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ShortLinkMutation) SetField(name string, value ent.Value) error {
	switch name {
	case shortlink.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case shortlink.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case shortlink.FieldKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKey(v)
		return nil
	case shortlink.FieldLink:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLink(v)
		return nil
	}
	return fmt.Errorf("unknown ShortLink field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ShortLinkMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ShortLinkMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ShortLinkMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown ShortLink numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ShortLinkMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ShortLinkMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ShortLinkMutation) ClearField(name string) error {
	return fmt.Errorf("unknown ShortLink nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ShortLinkMutation) ResetField(name string) error {
	switch name {
	case shortlink.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case shortlink.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case shortlink.FieldKey:
		m.ResetKey()
		return nil
	case shortlink.FieldLink:
		m.ResetLink()
		return nil
	}
	return fmt.Errorf("unknown ShortLink field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ShortLinkMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ShortLinkMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ShortLinkMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ShortLinkMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ShortLinkMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ShortLinkMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ShortLinkMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown ShortLink unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ShortLinkMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown ShortLink edge %s", name)
}
