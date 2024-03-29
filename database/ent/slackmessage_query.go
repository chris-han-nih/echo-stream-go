// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/echo-stream/database/ent/predicate"
	"github.com/echo-stream/database/ent/slackmessage"
	"github.com/echo-stream/database/ent/slackmessageattachment"
)

// SlackMessageQuery is the builder for querying SlackMessage entities.
type SlackMessageQuery struct {
	config
	ctx                         *QueryContext
	order                       []slackmessage.OrderOption
	inters                      []Interceptor
	predicates                  []predicate.SlackMessage
	withSlackMessageAttachments *SlackMessageAttachmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SlackMessageQuery builder.
func (smq *SlackMessageQuery) Where(ps ...predicate.SlackMessage) *SlackMessageQuery {
	smq.predicates = append(smq.predicates, ps...)
	return smq
}

// Limit the number of records to be returned by this query.
func (smq *SlackMessageQuery) Limit(limit int) *SlackMessageQuery {
	smq.ctx.Limit = &limit
	return smq
}

// Offset to start from.
func (smq *SlackMessageQuery) Offset(offset int) *SlackMessageQuery {
	smq.ctx.Offset = &offset
	return smq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smq *SlackMessageQuery) Unique(unique bool) *SlackMessageQuery {
	smq.ctx.Unique = &unique
	return smq
}

// Order specifies how the records should be ordered.
func (smq *SlackMessageQuery) Order(o ...slackmessage.OrderOption) *SlackMessageQuery {
	smq.order = append(smq.order, o...)
	return smq
}

// QuerySlackMessageAttachments chains the current query on the "slack_message_attachments" edge.
func (smq *SlackMessageQuery) QuerySlackMessageAttachments() *SlackMessageAttachmentQuery {
	query := (&SlackMessageAttachmentClient{config: smq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(slackmessage.Table, slackmessage.FieldID, selector),
			sqlgraph.To(slackmessageattachment.Table, slackmessageattachment.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, slackmessage.SlackMessageAttachmentsTable, slackmessage.SlackMessageAttachmentsColumn),
		)
		fromU = sqlgraph.SetNeighbors(smq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SlackMessage entity from the query.
// Returns a *NotFoundError when no SlackMessage was found.
func (smq *SlackMessageQuery) First(ctx context.Context) (*SlackMessage, error) {
	nodes, err := smq.Limit(1).All(setContextOp(ctx, smq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{slackmessage.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smq *SlackMessageQuery) FirstX(ctx context.Context) *SlackMessage {
	node, err := smq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SlackMessage ID from the query.
// Returns a *NotFoundError when no SlackMessage ID was found.
func (smq *SlackMessageQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = smq.Limit(1).IDs(setContextOp(ctx, smq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{slackmessage.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smq *SlackMessageQuery) FirstIDX(ctx context.Context) int {
	id, err := smq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SlackMessage entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SlackMessage entity is found.
// Returns a *NotFoundError when no SlackMessage entities are found.
func (smq *SlackMessageQuery) Only(ctx context.Context) (*SlackMessage, error) {
	nodes, err := smq.Limit(2).All(setContextOp(ctx, smq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{slackmessage.Label}
	default:
		return nil, &NotSingularError{slackmessage.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smq *SlackMessageQuery) OnlyX(ctx context.Context) *SlackMessage {
	node, err := smq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SlackMessage ID in the query.
// Returns a *NotSingularError when more than one SlackMessage ID is found.
// Returns a *NotFoundError when no entities are found.
func (smq *SlackMessageQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = smq.Limit(2).IDs(setContextOp(ctx, smq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{slackmessage.Label}
	default:
		err = &NotSingularError{slackmessage.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smq *SlackMessageQuery) OnlyIDX(ctx context.Context) int {
	id, err := smq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SlackMessages.
func (smq *SlackMessageQuery) All(ctx context.Context) ([]*SlackMessage, error) {
	ctx = setContextOp(ctx, smq.ctx, "All")
	if err := smq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SlackMessage, *SlackMessageQuery]()
	return withInterceptors[[]*SlackMessage](ctx, smq, qr, smq.inters)
}

// AllX is like All, but panics if an error occurs.
func (smq *SlackMessageQuery) AllX(ctx context.Context) []*SlackMessage {
	nodes, err := smq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SlackMessage IDs.
func (smq *SlackMessageQuery) IDs(ctx context.Context) (ids []int, err error) {
	if smq.ctx.Unique == nil && smq.path != nil {
		smq.Unique(true)
	}
	ctx = setContextOp(ctx, smq.ctx, "IDs")
	if err = smq.Select(slackmessage.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smq *SlackMessageQuery) IDsX(ctx context.Context) []int {
	ids, err := smq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smq *SlackMessageQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, smq.ctx, "Count")
	if err := smq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, smq, querierCount[*SlackMessageQuery](), smq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (smq *SlackMessageQuery) CountX(ctx context.Context) int {
	count, err := smq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smq *SlackMessageQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, smq.ctx, "Exist")
	switch _, err := smq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (smq *SlackMessageQuery) ExistX(ctx context.Context) bool {
	exist, err := smq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SlackMessageQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smq *SlackMessageQuery) Clone() *SlackMessageQuery {
	if smq == nil {
		return nil
	}
	return &SlackMessageQuery{
		config:                      smq.config,
		ctx:                         smq.ctx.Clone(),
		order:                       append([]slackmessage.OrderOption{}, smq.order...),
		inters:                      append([]Interceptor{}, smq.inters...),
		predicates:                  append([]predicate.SlackMessage{}, smq.predicates...),
		withSlackMessageAttachments: smq.withSlackMessageAttachments.Clone(),
		// clone intermediate query.
		sql:  smq.sql.Clone(),
		path: smq.path,
	}
}

// WithSlackMessageAttachments tells the query-builder to eager-load the nodes that are connected to
// the "slack_message_attachments" edge. The optional arguments are used to configure the query builder of the edge.
func (smq *SlackMessageQuery) WithSlackMessageAttachments(opts ...func(*SlackMessageAttachmentQuery)) *SlackMessageQuery {
	query := (&SlackMessageAttachmentClient{config: smq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	smq.withSlackMessageAttachments = query
	return smq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ApplicationId uuid.UUID `json:"ApplicationId,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SlackMessage.Query().
//		GroupBy(slackmessage.FieldApplicationId).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (smq *SlackMessageQuery) GroupBy(field string, fields ...string) *SlackMessageGroupBy {
	smq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SlackMessageGroupBy{build: smq}
	grbuild.flds = &smq.ctx.Fields
	grbuild.label = slackmessage.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ApplicationId uuid.UUID `json:"ApplicationId,omitempty"`
//	}
//
//	client.SlackMessage.Query().
//		Select(slackmessage.FieldApplicationId).
//		Scan(ctx, &v)
func (smq *SlackMessageQuery) Select(fields ...string) *SlackMessageSelect {
	smq.ctx.Fields = append(smq.ctx.Fields, fields...)
	sbuild := &SlackMessageSelect{SlackMessageQuery: smq}
	sbuild.label = slackmessage.Label
	sbuild.flds, sbuild.scan = &smq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SlackMessageSelect configured with the given aggregations.
func (smq *SlackMessageQuery) Aggregate(fns ...AggregateFunc) *SlackMessageSelect {
	return smq.Select().Aggregate(fns...)
}

func (smq *SlackMessageQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range smq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, smq); err != nil {
				return err
			}
		}
	}
	for _, f := range smq.ctx.Fields {
		if !slackmessage.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if smq.path != nil {
		prev, err := smq.path(ctx)
		if err != nil {
			return err
		}
		smq.sql = prev
	}
	return nil
}

func (smq *SlackMessageQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SlackMessage, error) {
	var (
		nodes       = []*SlackMessage{}
		_spec       = smq.querySpec()
		loadedTypes = [1]bool{
			smq.withSlackMessageAttachments != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SlackMessage).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SlackMessage{config: smq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, smq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := smq.withSlackMessageAttachments; query != nil {
		if err := smq.loadSlackMessageAttachments(ctx, query, nodes,
			func(n *SlackMessage) { n.Edges.SlackMessageAttachments = []*SlackMessageAttachment{} },
			func(n *SlackMessage, e *SlackMessageAttachment) {
				n.Edges.SlackMessageAttachments = append(n.Edges.SlackMessageAttachments, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (smq *SlackMessageQuery) loadSlackMessageAttachments(ctx context.Context, query *SlackMessageAttachmentQuery, nodes []*SlackMessage, init func(*SlackMessage), assign func(*SlackMessage, *SlackMessageAttachment)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*SlackMessage)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.SlackMessageAttachment(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(slackmessage.SlackMessageAttachmentsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.message_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "message_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "message_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (smq *SlackMessageQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smq.querySpec()
	_spec.Node.Columns = smq.ctx.Fields
	if len(smq.ctx.Fields) > 0 {
		_spec.Unique = smq.ctx.Unique != nil && *smq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, smq.driver, _spec)
}

func (smq *SlackMessageQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(slackmessage.Table, slackmessage.Columns, sqlgraph.NewFieldSpec(slackmessage.FieldID, field.TypeInt))
	_spec.From = smq.sql
	if unique := smq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if smq.path != nil {
		_spec.Unique = true
	}
	if fields := smq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, slackmessage.FieldID)
		for i := range fields {
			if fields[i] != slackmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := smq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smq *SlackMessageQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smq.driver.Dialect())
	t1 := builder.Table(slackmessage.Table)
	columns := smq.ctx.Fields
	if len(columns) == 0 {
		columns = slackmessage.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if smq.sql != nil {
		selector = smq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if smq.ctx.Unique != nil && *smq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range smq.predicates {
		p(selector)
	}
	for _, p := range smq.order {
		p(selector)
	}
	if offset := smq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SlackMessageGroupBy is the group-by builder for SlackMessage entities.
type SlackMessageGroupBy struct {
	selector
	build *SlackMessageQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smgb *SlackMessageGroupBy) Aggregate(fns ...AggregateFunc) *SlackMessageGroupBy {
	smgb.fns = append(smgb.fns, fns...)
	return smgb
}

// Scan applies the selector query and scans the result into the given value.
func (smgb *SlackMessageGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, smgb.build.ctx, "GroupBy")
	if err := smgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SlackMessageQuery, *SlackMessageGroupBy](ctx, smgb.build, smgb, smgb.build.inters, v)
}

func (smgb *SlackMessageGroupBy) sqlScan(ctx context.Context, root *SlackMessageQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(smgb.fns))
	for _, fn := range smgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*smgb.flds)+len(smgb.fns))
		for _, f := range *smgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*smgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SlackMessageSelect is the builder for selecting fields of SlackMessage entities.
type SlackMessageSelect struct {
	*SlackMessageQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sms *SlackMessageSelect) Aggregate(fns ...AggregateFunc) *SlackMessageSelect {
	sms.fns = append(sms.fns, fns...)
	return sms
}

// Scan applies the selector query and scans the result into the given value.
func (sms *SlackMessageSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sms.ctx, "Select")
	if err := sms.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SlackMessageQuery, *SlackMessageSelect](ctx, sms.SlackMessageQuery, sms, sms.inters, v)
}

func (sms *SlackMessageSelect) sqlScan(ctx context.Context, root *SlackMessageQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sms.fns))
	for _, fn := range sms.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sms.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
