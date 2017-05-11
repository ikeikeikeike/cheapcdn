// IMPORTANT! This is auto generated code by https://github.com/src-d/go-kallax
// Please, do not touch the code below, and if you do, do it under your own
// risk. Take into account that all the code you write here will be completely
// erased from earth the next time you generate the kallax models.
package models

import (
	"database/sql"
	"fmt"
	"time"

	"gopkg.in/src-d/go-kallax.v1"
	"gopkg.in/src-d/go-kallax.v1/types"
)

var _ types.SQLType
var _ fmt.Formatter

// NewNode returns a new instance of Node.
func NewNode() (record *Node) {
	return new(Node)
}

// GetID returns the primary key of the model.
func (r *Node) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Node) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "url":
		return &r.URL, nil
	case "alive":
		return &r.Alive, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Node: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Node) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "url":
		return r.URL, nil
	case "alive":
		return r.Alive, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Node: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Node) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Objects":
		return new(Object), nil

	}
	return nil, fmt.Errorf("kallax: model Node has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Node) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Objects":
		records, ok := rel.([]kallax.Record)
		if !ok {
			return fmt.Errorf("kallax: relationship field %s needs a collection of records, not %T", field, rel)
		}

		r.Objects = make([]*Object, len(records))
		for i, record := range records {
			rel, ok := record.(*Object)
			if !ok {
				return fmt.Errorf("kallax: element of type %T cannot be added to relationship %s", record, field)
			}
			r.Objects[i] = rel
		}
		return nil

	}
	return fmt.Errorf("kallax: model Node has no relationship %s", field)
}

// NodeStore is the entity to access the records of the type Node
// in the database.
type NodeStore struct {
	*kallax.Store
}

// NewNodeStore creates a new instance of NodeStore
// using a SQL database.
func NewNodeStore(db *sql.DB) *NodeStore {
	return &NodeStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *NodeStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *NodeStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *NodeStore) Debug() *NodeStore {
	return &NodeStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *NodeStore) DebugWith(logger kallax.LoggerFunc) *NodeStore {
	return &NodeStore{s.Store.DebugWith(logger)}
}

func (s *NodeStore) relationshipRecords(record *Node) []kallax.RecordWithSchema {
	var records []kallax.RecordWithSchema

	for _, rec := range record.Objects {
		rec.ClearVirtualColumns()
		rec.AddVirtualColumn("node_id", record.GetID())
		records = append(records, kallax.RecordWithSchema{
			Schema: Schema.Object.BaseSchema,
			Record: rec,
		})
	}

	return records
}

// Insert inserts a Node in the database. A non-persisted object is
// required for this operation.
func (s *NodeStore) Insert(record *Node) error {

	if err := record.BeforeSave(); err != nil {
		return err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {

			if err := s.Insert(Schema.Node.BaseSchema, record); err != nil {
				return err
			}

			for _, r := range records {
				if err := kallax.ApplyBeforeEvents(r.Record); err != nil {
					return err
				}
				persisted := r.Record.IsPersisted()

				if _, err := s.Save(r.Schema, r.Record); err != nil {
					return err
				}

				if err := kallax.ApplyAfterEvents(r.Record, persisted); err != nil {
					return err
				}
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Node.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *NodeStore) Update(record *Node, cols ...kallax.SchemaField) (updated int64, err error) {

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	records := s.relationshipRecords(record)

	if len(records) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {

			updated, err = s.Update(Schema.Node.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			for _, r := range records {
				if err := kallax.ApplyBeforeEvents(r.Record); err != nil {
					return err
				}
				persisted := r.Record.IsPersisted()

				if _, err := s.Save(r.Schema, r.Record); err != nil {
					return err
				}

				if err := kallax.ApplyAfterEvents(r.Record, persisted); err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Node.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *NodeStore) Save(record *Node) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *NodeStore) Delete(record *Node) error {

	return s.Store.Delete(Schema.Node.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *NodeStore) Find(q *NodeQuery) (*NodeResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewNodeResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *NodeStore) MustFind(q *NodeQuery) *NodeResultSet {
	return NewNodeResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *NodeStore) Count(q *NodeQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *NodeStore) MustCount(q *NodeQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *NodeStore) FindOne(q *NodeQuery) (*Node, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *NodeStore) MustFindOne(q *NodeQuery) *Node {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Node with the data in the database and
// makes it writable.
func (s *NodeStore) Reload(record *Node) error {
	return s.Store.Reload(Schema.Node.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *NodeStore) Transaction(callback func(*NodeStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&NodeStore{store})
	})
}

// RemoveObjects removes the given items of the Objects field of the
// model. If no items are given, it removes all of them.
// The items will also be removed from the passed record inside this method.
func (s *NodeStore) RemoveObjects(record *Node, deleted ...*Object) error {
	var updated []*Object
	var clear bool
	if len(deleted) == 0 {
		clear = true
		deleted = record.Objects
		if len(deleted) == 0 {
			return nil
		}
	}

	if len(deleted) > 1 {
		err := s.Store.Transaction(func(s *kallax.Store) error {
			for _, d := range deleted {
				var r kallax.Record = d

				if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
					if err := beforeDeleter.BeforeDelete(); err != nil {
						return err
					}
				}

				if err := s.Delete(Schema.Object.BaseSchema, d); err != nil {
					return err
				}

				if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
					if err := afterDeleter.AfterDelete(); err != nil {
						return err
					}
				}
			}
			return nil
		})

		if err != nil {
			return err
		}

		if clear {
			record.Objects = nil
			return nil
		}
	} else {
		var r kallax.Record = deleted[0]
		if beforeDeleter, ok := r.(kallax.BeforeDeleter); ok {
			if err := beforeDeleter.BeforeDelete(); err != nil {
				return err
			}
		}

		var err error
		if afterDeleter, ok := r.(kallax.AfterDeleter); ok {
			err = s.Store.Transaction(func(s *kallax.Store) error {
				err := s.Delete(Schema.Object.BaseSchema, r)
				if err != nil {
					return err
				}

				return afterDeleter.AfterDelete()
			})
		} else {
			err = s.Store.Delete(Schema.Object.BaseSchema, deleted[0])
		}

		if err != nil {
			return err
		}
	}

	for _, r := range record.Objects {
		var found bool
		for _, d := range deleted {
			if d.GetID().Equals(r.GetID()) {
				found = true
				break
			}
		}
		if !found {
			updated = append(updated, r)
		}
	}
	record.Objects = updated
	return nil
}

// NodeQuery is the object used to create queries for the Node
// entity.
type NodeQuery struct {
	*kallax.BaseQuery
}

// NewNodeQuery returns a new instance of NodeQuery.
func NewNodeQuery() *NodeQuery {
	return &NodeQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Node.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *NodeQuery) Select(columns ...kallax.SchemaField) *NodeQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *NodeQuery) SelectNot(columns ...kallax.SchemaField) *NodeQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *NodeQuery) Copy() *NodeQuery {
	return &NodeQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *NodeQuery) Order(cols ...kallax.ColumnOrder) *NodeQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *NodeQuery) BatchSize(size uint64) *NodeQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *NodeQuery) Limit(n uint64) *NodeQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *NodeQuery) Offset(n uint64) *NodeQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *NodeQuery) Where(cond kallax.Condition) *NodeQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *NodeQuery) WithObjects(cond kallax.Condition) *NodeQuery {
	q.AddRelation(Schema.Object.BaseSchema, "Objects", kallax.OneToMany, cond)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *NodeQuery) FindByID(v ...int64) *NodeQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Node.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *NodeQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *NodeQuery {
	return q.Where(cond(Schema.Node.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *NodeQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *NodeQuery {
	return q.Where(cond(Schema.Node.UpdatedAt, v))
}

// FindByURL adds a new filter to the query that will require that
// the URL property is equal to the passed value.
func (q *NodeQuery) FindByURL(v string) *NodeQuery {
	return q.Where(kallax.Eq(Schema.Node.URL, v))
}

// FindByAlive adds a new filter to the query that will require that
// the Alive property is equal to the passed value.
func (q *NodeQuery) FindByAlive(v bool) *NodeQuery {
	return q.Where(kallax.Eq(Schema.Node.Alive, v))
}

// NodeResultSet is the set of results returned by a query to the
// database.
type NodeResultSet struct {
	ResultSet kallax.ResultSet
	last      *Node
	lastErr   error
}

// NewNodeResultSet creates a new result set for rows of the type
// Node.
func NewNodeResultSet(rs kallax.ResultSet) *NodeResultSet {
	return &NodeResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *NodeResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Node.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Node)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Node")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *NodeResultSet) Get() (*Node, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *NodeResultSet) ForEach(fn func(*Node) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *NodeResultSet) All() ([]*Node, error) {
	var result []*Node
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *NodeResultSet) One() (*Node, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *NodeResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *NodeResultSet) Close() error {
	return rs.ResultSet.Close()
}

// NewObject returns a new instance of Object.
func NewObject() (record *Object) {
	return new(Object)
}

// GetID returns the primary key of the model.
func (r *Object) GetID() kallax.Identifier {
	return (*kallax.NumericID)(&r.ID)
}

// ColumnAddress returns the pointer to the value of the given column.
func (r *Object) ColumnAddress(col string) (interface{}, error) {
	switch col {
	case "id":
		return (*kallax.NumericID)(&r.ID), nil
	case "created_at":
		return &r.Timestamps.CreatedAt, nil
	case "updated_at":
		return &r.Timestamps.UpdatedAt, nil
	case "node_id":
		return types.Nullable(kallax.VirtualColumn("node_id", r, new(kallax.NumericID))), nil
	case "url":
		return &r.URL, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Object: %s", col)
	}
}

// Value returns the value of the given column.
func (r *Object) Value(col string) (interface{}, error) {
	switch col {
	case "id":
		return r.ID, nil
	case "created_at":
		return r.Timestamps.CreatedAt, nil
	case "updated_at":
		return r.Timestamps.UpdatedAt, nil
	case "node_id":
		return r.Model.VirtualColumn(col), nil
	case "url":
		return r.URL, nil

	default:
		return nil, fmt.Errorf("kallax: invalid column in Object: %s", col)
	}
}

// NewRelationshipRecord returns a new record for the relatiobship in the given
// field.
func (r *Object) NewRelationshipRecord(field string) (kallax.Record, error) {
	switch field {
	case "Node":
		return new(Node), nil

	}
	return nil, fmt.Errorf("kallax: model Object has no relationship %s", field)
}

// SetRelationship sets the given relationship in the given field.
func (r *Object) SetRelationship(field string, rel interface{}) error {
	switch field {
	case "Node":
		val, ok := rel.(*Node)
		if !ok {
			return fmt.Errorf("kallax: record of type %t can't be assigned to relationship Node", rel)
		}
		if !val.GetID().IsEmpty() {
			r.Node = val
		}

		return nil

	}
	return fmt.Errorf("kallax: model Object has no relationship %s", field)
}

// ObjectStore is the entity to access the records of the type Object
// in the database.
type ObjectStore struct {
	*kallax.Store
}

// NewObjectStore creates a new instance of ObjectStore
// using a SQL database.
func NewObjectStore(db *sql.DB) *ObjectStore {
	return &ObjectStore{kallax.NewStore(db)}
}

// GenericStore returns the generic store of this store.
func (s *ObjectStore) GenericStore() *kallax.Store {
	return s.Store
}

// SetGenericStore changes the generic store of this store.
func (s *ObjectStore) SetGenericStore(store *kallax.Store) {
	s.Store = store
}

// Debug returns a new store that will print all SQL statements to stdout using
// the log.Printf function.
func (s *ObjectStore) Debug() *ObjectStore {
	return &ObjectStore{s.Store.Debug()}
}

// DebugWith returns a new store that will print all SQL statements using the
// given logger function.
func (s *ObjectStore) DebugWith(logger kallax.LoggerFunc) *ObjectStore {
	return &ObjectStore{s.Store.DebugWith(logger)}
}

func (s *ObjectStore) inverseRecords(record *Object) []kallax.RecordWithSchema {
	record.ClearVirtualColumns()
	var records []kallax.RecordWithSchema

	if record.Node != nil {
		record.AddVirtualColumn("node_id", record.Node.GetID())
		records = append(records, kallax.RecordWithSchema{
			Schema: Schema.Node.BaseSchema,
			Record: record.Node,
		})
	}

	return records
}

// Insert inserts a Object in the database. A non-persisted object is
// required for this operation.
func (s *ObjectStore) Insert(record *Object) error {

	if err := record.BeforeSave(); err != nil {
		return err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		return s.Store.Transaction(func(s *kallax.Store) error {

			for _, r := range inverseRecords {
				if err := kallax.ApplyBeforeEvents(r.Record); err != nil {
					return err
				}
				persisted := r.Record.IsPersisted()

				if _, err := s.Save(r.Schema, r.Record); err != nil {
					return err
				}

				if err := kallax.ApplyAfterEvents(r.Record, persisted); err != nil {
					return err
				}
			}

			if err := s.Insert(Schema.Object.BaseSchema, record); err != nil {
				return err
			}

			return nil
		})
	}

	return s.Store.Insert(Schema.Object.BaseSchema, record)

}

// Update updates the given record on the database. If the columns are given,
// only these columns will be updated. Otherwise all of them will be.
// Be very careful with this, as you will have a potentially different object
// in memory but not on the database.
// Only writable records can be updated. Writable objects are those that have
// been just inserted or retrieved using a query with no custom select fields.
func (s *ObjectStore) Update(record *Object, cols ...kallax.SchemaField) (updated int64, err error) {

	if err := record.BeforeSave(); err != nil {
		return 0, err
	}

	inverseRecords := s.inverseRecords(record)

	if len(inverseRecords) > 0 {
		err = s.Store.Transaction(func(s *kallax.Store) error {

			for _, r := range inverseRecords {
				if err := kallax.ApplyBeforeEvents(r.Record); err != nil {
					return err
				}
				persisted := r.Record.IsPersisted()

				if _, err := s.Save(r.Schema, r.Record); err != nil {
					return err
				}

				if err := kallax.ApplyAfterEvents(r.Record, persisted); err != nil {
					return err
				}
			}

			updated, err = s.Update(Schema.Object.BaseSchema, record, cols...)
			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			return 0, err
		}

		return updated, nil
	}

	return s.Store.Update(Schema.Object.BaseSchema, record, cols...)

}

// Save inserts the object if the record is not persisted, otherwise it updates
// it. Same rules of Update and Insert apply depending on the case.
func (s *ObjectStore) Save(record *Object) (updated bool, err error) {
	if !record.IsPersisted() {
		return false, s.Insert(record)
	}

	rowsUpdated, err := s.Update(record)
	if err != nil {
		return false, err
	}

	return rowsUpdated > 0, nil
}

// Delete removes the given record from the database.
func (s *ObjectStore) Delete(record *Object) error {

	return s.Store.Delete(Schema.Object.BaseSchema, record)

}

// Find returns the set of results for the given query.
func (s *ObjectStore) Find(q *ObjectQuery) (*ObjectResultSet, error) {
	rs, err := s.Store.Find(q)
	if err != nil {
		return nil, err
	}

	return NewObjectResultSet(rs), nil
}

// MustFind returns the set of results for the given query, but panics if there
// is any error.
func (s *ObjectStore) MustFind(q *ObjectQuery) *ObjectResultSet {
	return NewObjectResultSet(s.Store.MustFind(q))
}

// Count returns the number of rows that would be retrieved with the given
// query.
func (s *ObjectStore) Count(q *ObjectQuery) (int64, error) {
	return s.Store.Count(q)
}

// MustCount returns the number of rows that would be retrieved with the given
// query, but panics if there is an error.
func (s *ObjectStore) MustCount(q *ObjectQuery) int64 {
	return s.Store.MustCount(q)
}

// FindOne returns the first row returned by the given query.
// `ErrNotFound` is returned if there are no results.
func (s *ObjectStore) FindOne(q *ObjectQuery) (*Object, error) {
	q.Limit(1)
	q.Offset(0)
	rs, err := s.Find(q)
	if err != nil {
		return nil, err
	}

	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// MustFindOne returns the first row retrieved by the given query. It panics
// if there is an error or if there are no rows.
func (s *ObjectStore) MustFindOne(q *ObjectQuery) *Object {
	record, err := s.FindOne(q)
	if err != nil {
		panic(err)
	}
	return record
}

// Reload refreshes the Object with the data in the database and
// makes it writable.
func (s *ObjectStore) Reload(record *Object) error {
	return s.Store.Reload(Schema.Object.BaseSchema, record)
}

// Transaction executes the given callback in a transaction and rollbacks if
// an error is returned.
// The transaction is only open in the store passed as a parameter to the
// callback.
func (s *ObjectStore) Transaction(callback func(*ObjectStore) error) error {
	if callback == nil {
		return kallax.ErrInvalidTxCallback
	}

	return s.Store.Transaction(func(store *kallax.Store) error {
		return callback(&ObjectStore{store})
	})
}

// ObjectQuery is the object used to create queries for the Object
// entity.
type ObjectQuery struct {
	*kallax.BaseQuery
}

// NewObjectQuery returns a new instance of ObjectQuery.
func NewObjectQuery() *ObjectQuery {
	return &ObjectQuery{
		BaseQuery: kallax.NewBaseQuery(Schema.Object.BaseSchema),
	}
}

// Select adds columns to select in the query.
func (q *ObjectQuery) Select(columns ...kallax.SchemaField) *ObjectQuery {
	if len(columns) == 0 {
		return q
	}
	q.BaseQuery.Select(columns...)
	return q
}

// SelectNot excludes columns from being selected in the query.
func (q *ObjectQuery) SelectNot(columns ...kallax.SchemaField) *ObjectQuery {
	q.BaseQuery.SelectNot(columns...)
	return q
}

// Copy returns a new identical copy of the query. Remember queries are mutable
// so make a copy any time you need to reuse them.
func (q *ObjectQuery) Copy() *ObjectQuery {
	return &ObjectQuery{
		BaseQuery: q.BaseQuery.Copy(),
	}
}

// Order adds order clauses to the query for the given columns.
func (q *ObjectQuery) Order(cols ...kallax.ColumnOrder) *ObjectQuery {
	q.BaseQuery.Order(cols...)
	return q
}

// BatchSize sets the number of items to fetch per batch when there are 1:N
// relationships selected in the query.
func (q *ObjectQuery) BatchSize(size uint64) *ObjectQuery {
	q.BaseQuery.BatchSize(size)
	return q
}

// Limit sets the max number of items to retrieve.
func (q *ObjectQuery) Limit(n uint64) *ObjectQuery {
	q.BaseQuery.Limit(n)
	return q
}

// Offset sets the number of items to skip from the result set of items.
func (q *ObjectQuery) Offset(n uint64) *ObjectQuery {
	q.BaseQuery.Offset(n)
	return q
}

// Where adds a condition to the query. All conditions added are concatenated
// using a logical AND.
func (q *ObjectQuery) Where(cond kallax.Condition) *ObjectQuery {
	q.BaseQuery.Where(cond)
	return q
}

func (q *ObjectQuery) WithNode() *ObjectQuery {
	q.AddRelation(Schema.Node.BaseSchema, "Node", kallax.OneToOne, nil)
	return q
}

// FindByID adds a new filter to the query that will require that
// the ID property is equal to one of the passed values; if no passed values,
// it will do nothing.
func (q *ObjectQuery) FindByID(v ...int64) *ObjectQuery {
	if len(v) == 0 {
		return q
	}
	values := make([]interface{}, len(v))
	for i, val := range v {
		values[i] = val
	}
	return q.Where(kallax.In(Schema.Object.ID, values...))
}

// FindByCreatedAt adds a new filter to the query that will require that
// the CreatedAt property is equal to the passed value.
func (q *ObjectQuery) FindByCreatedAt(cond kallax.ScalarCond, v time.Time) *ObjectQuery {
	return q.Where(cond(Schema.Object.CreatedAt, v))
}

// FindByUpdatedAt adds a new filter to the query that will require that
// the UpdatedAt property is equal to the passed value.
func (q *ObjectQuery) FindByUpdatedAt(cond kallax.ScalarCond, v time.Time) *ObjectQuery {
	return q.Where(cond(Schema.Object.UpdatedAt, v))
}

// FindByNode adds a new filter to the query that will require that
// the foreign key of Node is equal to the passed value.
func (q *ObjectQuery) FindByNode(v int64) *ObjectQuery {
	return q.Where(kallax.Eq(Schema.Object.NodeFK, v))
}

// FindByURL adds a new filter to the query that will require that
// the URL property is equal to the passed value.
func (q *ObjectQuery) FindByURL(v string) *ObjectQuery {
	return q.Where(kallax.Eq(Schema.Object.URL, v))
}

// ObjectResultSet is the set of results returned by a query to the
// database.
type ObjectResultSet struct {
	ResultSet kallax.ResultSet
	last      *Object
	lastErr   error
}

// NewObjectResultSet creates a new result set for rows of the type
// Object.
func NewObjectResultSet(rs kallax.ResultSet) *ObjectResultSet {
	return &ObjectResultSet{ResultSet: rs}
}

// Next fetches the next item in the result set and returns true if there is
// a next item.
// The result set is closed automatically when there are no more items.
func (rs *ObjectResultSet) Next() bool {
	if !rs.ResultSet.Next() {
		rs.lastErr = rs.ResultSet.Close()
		rs.last = nil
		return false
	}

	var record kallax.Record
	record, rs.lastErr = rs.ResultSet.Get(Schema.Object.BaseSchema)
	if rs.lastErr != nil {
		rs.last = nil
	} else {
		var ok bool
		rs.last, ok = record.(*Object)
		if !ok {
			rs.lastErr = fmt.Errorf("kallax: unable to convert record to *Object")
			rs.last = nil
		}
	}

	return true
}

// Get retrieves the last fetched item from the result set and the last error.
func (rs *ObjectResultSet) Get() (*Object, error) {
	return rs.last, rs.lastErr
}

// ForEach iterates over the complete result set passing every record found to
// the given callback. It is possible to stop the iteration by returning
// `kallax.ErrStop` in the callback.
// Result set is always closed at the end.
func (rs *ObjectResultSet) ForEach(fn func(*Object) error) error {
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return err
		}

		if err := fn(record); err != nil {
			if err == kallax.ErrStop {
				return rs.Close()
			}

			return err
		}
	}
	return nil
}

// All returns all records on the result set and closes the result set.
func (rs *ObjectResultSet) All() ([]*Object, error) {
	var result []*Object
	for rs.Next() {
		record, err := rs.Get()
		if err != nil {
			return nil, err
		}
		result = append(result, record)
	}
	return result, nil
}

// One returns the first record on the result set and closes the result set.
func (rs *ObjectResultSet) One() (*Object, error) {
	if !rs.Next() {
		return nil, kallax.ErrNotFound
	}

	record, err := rs.Get()
	if err != nil {
		return nil, err
	}

	if err := rs.Close(); err != nil {
		return nil, err
	}

	return record, nil
}

// Err returns the last error occurred.
func (rs *ObjectResultSet) Err() error {
	return rs.lastErr
}

// Close closes the result set.
func (rs *ObjectResultSet) Close() error {
	return rs.ResultSet.Close()
}

type schema struct {
	Node   *schemaNode
	Object *schemaObject
}

type schemaNode struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	URL       kallax.SchemaField
	Alive     kallax.SchemaField
}

type schemaObject struct {
	*kallax.BaseSchema
	ID        kallax.SchemaField
	CreatedAt kallax.SchemaField
	UpdatedAt kallax.SchemaField
	NodeFK    kallax.SchemaField
	URL       kallax.SchemaField
}

var Schema = &schema{
	Node: &schemaNode{
		BaseSchema: kallax.NewBaseSchema(
			"nodes",
			"__node",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Objects": kallax.NewForeignKey("node_id", false),
			},
			func() kallax.Record {
				return new(Node)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("url"),
			kallax.NewSchemaField("alive"),
		),
		ID:        kallax.NewSchemaField("id"),
		CreatedAt: kallax.NewSchemaField("created_at"),
		UpdatedAt: kallax.NewSchemaField("updated_at"),
		URL:       kallax.NewSchemaField("url"),
		Alive:     kallax.NewSchemaField("alive"),
	},
	Object: &schemaObject{
		BaseSchema: kallax.NewBaseSchema(
			"objects",
			"__object",
			kallax.NewSchemaField("id"),
			kallax.ForeignKeys{
				"Node": kallax.NewForeignKey("node_id", true),
			},
			func() kallax.Record {
				return new(Object)
			},
			true,
			kallax.NewSchemaField("id"),
			kallax.NewSchemaField("created_at"),
			kallax.NewSchemaField("updated_at"),
			kallax.NewSchemaField("node_id"),
			kallax.NewSchemaField("url"),
		),
		ID:        kallax.NewSchemaField("id"),
		CreatedAt: kallax.NewSchemaField("created_at"),
		UpdatedAt: kallax.NewSchemaField("updated_at"),
		NodeFK:    kallax.NewSchemaField("node_id"),
		URL:       kallax.NewSchemaField("url"),
	},
}
