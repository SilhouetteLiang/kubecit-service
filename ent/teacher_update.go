// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"kubecit-service/ent/course"
	"kubecit-service/ent/predicate"
	"kubecit-service/ent/teacher"
	"kubecit-service/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TeacherUpdate is the builder for updating Teacher entities.
type TeacherUpdate struct {
	config
	hooks    []Hook
	mutation *TeacherMutation
}

// Where appends a list predicates to the TeacherUpdate builder.
func (tu *TeacherUpdate) Where(ps ...predicate.Teacher) *TeacherUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetDetail sets the "detail" field.
func (tu *TeacherUpdate) SetDetail(s string) *TeacherUpdate {
	tu.mutation.SetDetail(s)
	return tu
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableDetail(s *string) *TeacherUpdate {
	if s != nil {
		tu.SetDetail(*s)
	}
	return tu
}

// ClearDetail clears the value of the "detail" field.
func (tu *TeacherUpdate) ClearDetail() *TeacherUpdate {
	tu.mutation.ClearDetail()
	return tu
}

// SetCurriculumVitae sets the "curriculum_vitae" field.
func (tu *TeacherUpdate) SetCurriculumVitae(s string) *TeacherUpdate {
	tu.mutation.SetCurriculumVitae(s)
	return tu
}

// SetNillableCurriculumVitae sets the "curriculum_vitae" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableCurriculumVitae(s *string) *TeacherUpdate {
	if s != nil {
		tu.SetCurriculumVitae(*s)
	}
	return tu
}

// ClearCurriculumVitae clears the value of the "curriculum_vitae" field.
func (tu *TeacherUpdate) ClearCurriculumVitae() *TeacherUpdate {
	tu.mutation.ClearCurriculumVitae()
	return tu
}

// SetWorks sets the "works" field.
func (tu *TeacherUpdate) SetWorks(s string) *TeacherUpdate {
	tu.mutation.SetWorks(s)
	return tu
}

// SetNillableWorks sets the "works" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableWorks(s *string) *TeacherUpdate {
	if s != nil {
		tu.SetWorks(*s)
	}
	return tu
}

// ClearWorks clears the value of the "works" field.
func (tu *TeacherUpdate) ClearWorks() *TeacherUpdate {
	tu.mutation.ClearWorks()
	return tu
}

// SetSkills sets the "skills" field.
func (tu *TeacherUpdate) SetSkills(s string) *TeacherUpdate {
	tu.mutation.SetSkills(s)
	return tu
}

// SetNillableSkills sets the "skills" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableSkills(s *string) *TeacherUpdate {
	if s != nil {
		tu.SetSkills(*s)
	}
	return tu
}

// ClearSkills clears the value of the "skills" field.
func (tu *TeacherUpdate) ClearSkills() *TeacherUpdate {
	tu.mutation.ClearSkills()
	return tu
}

// SetName sets the "name" field.
func (tu *TeacherUpdate) SetName(s string) *TeacherUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetLevel sets the "level" field.
func (tu *TeacherUpdate) SetLevel(i int) *TeacherUpdate {
	tu.mutation.ResetLevel()
	tu.mutation.SetLevel(i)
	return tu
}

// AddLevel adds i to the "level" field.
func (tu *TeacherUpdate) AddLevel(i int) *TeacherUpdate {
	tu.mutation.AddLevel(i)
	return tu
}

// SetAvator sets the "avator" field.
func (tu *TeacherUpdate) SetAvator(s string) *TeacherUpdate {
	tu.mutation.SetAvator(s)
	return tu
}

// SetNillableAvator sets the "avator" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableAvator(s *string) *TeacherUpdate {
	if s != nil {
		tu.SetAvator(*s)
	}
	return tu
}

// ClearAvator clears the value of the "avator" field.
func (tu *TeacherUpdate) ClearAvator() *TeacherUpdate {
	tu.mutation.ClearAvator()
	return tu
}

// SetCreateAt sets the "create_at" field.
func (tu *TeacherUpdate) SetCreateAt(t time.Time) *TeacherUpdate {
	tu.mutation.SetCreateAt(t)
	return tu
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableCreateAt(t *time.Time) *TeacherUpdate {
	if t != nil {
		tu.SetCreateAt(*t)
	}
	return tu
}

// SetUpdateAt sets the "update_at" field.
func (tu *TeacherUpdate) SetUpdateAt(t time.Time) *TeacherUpdate {
	tu.mutation.SetUpdateAt(t)
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TeacherUpdate) SetUserID(i int) *TeacherUpdate {
	tu.mutation.SetUserID(i)
	return tu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tu *TeacherUpdate) SetNillableUserID(i *int) *TeacherUpdate {
	if i != nil {
		tu.SetUserID(*i)
	}
	return tu
}

// ClearUserID clears the value of the "user_id" field.
func (tu *TeacherUpdate) ClearUserID() *TeacherUpdate {
	tu.mutation.ClearUserID()
	return tu
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (tu *TeacherUpdate) AddCourseIDs(ids ...int) *TeacherUpdate {
	tu.mutation.AddCourseIDs(ids...)
	return tu
}

// AddCourses adds the "courses" edges to the Course entity.
func (tu *TeacherUpdate) AddCourses(c ...*Course) *TeacherUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.AddCourseIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (tu *TeacherUpdate) SetUser(u *User) *TeacherUpdate {
	return tu.SetUserID(u.ID)
}

// Mutation returns the TeacherMutation object of the builder.
func (tu *TeacherUpdate) Mutation() *TeacherMutation {
	return tu.mutation
}

// ClearCourses clears all "courses" edges to the Course entity.
func (tu *TeacherUpdate) ClearCourses() *TeacherUpdate {
	tu.mutation.ClearCourses()
	return tu
}

// RemoveCourseIDs removes the "courses" edge to Course entities by IDs.
func (tu *TeacherUpdate) RemoveCourseIDs(ids ...int) *TeacherUpdate {
	tu.mutation.RemoveCourseIDs(ids...)
	return tu
}

// RemoveCourses removes "courses" edges to Course entities.
func (tu *TeacherUpdate) RemoveCourses(c ...*Course) *TeacherUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tu.RemoveCourseIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (tu *TeacherUpdate) ClearUser() *TeacherUpdate {
	tu.mutation.ClearUser()
	return tu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeacherUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeacherUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeacherUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeacherUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TeacherUpdate) defaults() {
	if _, ok := tu.mutation.UpdateAt(); !ok {
		v := teacher.UpdateDefaultUpdateAt()
		tu.mutation.SetUpdateAt(v)
	}
}

func (tu *TeacherUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(teacher.Table, teacher.Columns, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Detail(); ok {
		_spec.SetField(teacher.FieldDetail, field.TypeString, value)
	}
	if tu.mutation.DetailCleared() {
		_spec.ClearField(teacher.FieldDetail, field.TypeString)
	}
	if value, ok := tu.mutation.CurriculumVitae(); ok {
		_spec.SetField(teacher.FieldCurriculumVitae, field.TypeString, value)
	}
	if tu.mutation.CurriculumVitaeCleared() {
		_spec.ClearField(teacher.FieldCurriculumVitae, field.TypeString)
	}
	if value, ok := tu.mutation.Works(); ok {
		_spec.SetField(teacher.FieldWorks, field.TypeString, value)
	}
	if tu.mutation.WorksCleared() {
		_spec.ClearField(teacher.FieldWorks, field.TypeString)
	}
	if value, ok := tu.mutation.Skills(); ok {
		_spec.SetField(teacher.FieldSkills, field.TypeString, value)
	}
	if tu.mutation.SkillsCleared() {
		_spec.ClearField(teacher.FieldSkills, field.TypeString)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(teacher.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Level(); ok {
		_spec.SetField(teacher.FieldLevel, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedLevel(); ok {
		_spec.AddField(teacher.FieldLevel, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Avator(); ok {
		_spec.SetField(teacher.FieldAvator, field.TypeString, value)
	}
	if tu.mutation.AvatorCleared() {
		_spec.ClearField(teacher.FieldAvator, field.TypeString)
	}
	if value, ok := tu.mutation.CreateAt(); ok {
		_spec.SetField(teacher.FieldCreateAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdateAt(); ok {
		_spec.SetField(teacher.FieldUpdateAt, field.TypeTime, value)
	}
	if tu.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.CoursesTable,
			Columns: []string{teacher.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !tu.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.CoursesTable,
			Columns: []string{teacher.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.CoursesTable,
			Columns: []string{teacher.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   teacher.UserTable,
			Columns: []string{teacher.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   teacher.UserTable,
			Columns: []string{teacher.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeacherUpdateOne is the builder for updating a single Teacher entity.
type TeacherUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TeacherMutation
}

// SetDetail sets the "detail" field.
func (tuo *TeacherUpdateOne) SetDetail(s string) *TeacherUpdateOne {
	tuo.mutation.SetDetail(s)
	return tuo
}

// SetNillableDetail sets the "detail" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableDetail(s *string) *TeacherUpdateOne {
	if s != nil {
		tuo.SetDetail(*s)
	}
	return tuo
}

// ClearDetail clears the value of the "detail" field.
func (tuo *TeacherUpdateOne) ClearDetail() *TeacherUpdateOne {
	tuo.mutation.ClearDetail()
	return tuo
}

// SetCurriculumVitae sets the "curriculum_vitae" field.
func (tuo *TeacherUpdateOne) SetCurriculumVitae(s string) *TeacherUpdateOne {
	tuo.mutation.SetCurriculumVitae(s)
	return tuo
}

// SetNillableCurriculumVitae sets the "curriculum_vitae" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableCurriculumVitae(s *string) *TeacherUpdateOne {
	if s != nil {
		tuo.SetCurriculumVitae(*s)
	}
	return tuo
}

// ClearCurriculumVitae clears the value of the "curriculum_vitae" field.
func (tuo *TeacherUpdateOne) ClearCurriculumVitae() *TeacherUpdateOne {
	tuo.mutation.ClearCurriculumVitae()
	return tuo
}

// SetWorks sets the "works" field.
func (tuo *TeacherUpdateOne) SetWorks(s string) *TeacherUpdateOne {
	tuo.mutation.SetWorks(s)
	return tuo
}

// SetNillableWorks sets the "works" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableWorks(s *string) *TeacherUpdateOne {
	if s != nil {
		tuo.SetWorks(*s)
	}
	return tuo
}

// ClearWorks clears the value of the "works" field.
func (tuo *TeacherUpdateOne) ClearWorks() *TeacherUpdateOne {
	tuo.mutation.ClearWorks()
	return tuo
}

// SetSkills sets the "skills" field.
func (tuo *TeacherUpdateOne) SetSkills(s string) *TeacherUpdateOne {
	tuo.mutation.SetSkills(s)
	return tuo
}

// SetNillableSkills sets the "skills" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableSkills(s *string) *TeacherUpdateOne {
	if s != nil {
		tuo.SetSkills(*s)
	}
	return tuo
}

// ClearSkills clears the value of the "skills" field.
func (tuo *TeacherUpdateOne) ClearSkills() *TeacherUpdateOne {
	tuo.mutation.ClearSkills()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TeacherUpdateOne) SetName(s string) *TeacherUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetLevel sets the "level" field.
func (tuo *TeacherUpdateOne) SetLevel(i int) *TeacherUpdateOne {
	tuo.mutation.ResetLevel()
	tuo.mutation.SetLevel(i)
	return tuo
}

// AddLevel adds i to the "level" field.
func (tuo *TeacherUpdateOne) AddLevel(i int) *TeacherUpdateOne {
	tuo.mutation.AddLevel(i)
	return tuo
}

// SetAvator sets the "avator" field.
func (tuo *TeacherUpdateOne) SetAvator(s string) *TeacherUpdateOne {
	tuo.mutation.SetAvator(s)
	return tuo
}

// SetNillableAvator sets the "avator" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableAvator(s *string) *TeacherUpdateOne {
	if s != nil {
		tuo.SetAvator(*s)
	}
	return tuo
}

// ClearAvator clears the value of the "avator" field.
func (tuo *TeacherUpdateOne) ClearAvator() *TeacherUpdateOne {
	tuo.mutation.ClearAvator()
	return tuo
}

// SetCreateAt sets the "create_at" field.
func (tuo *TeacherUpdateOne) SetCreateAt(t time.Time) *TeacherUpdateOne {
	tuo.mutation.SetCreateAt(t)
	return tuo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableCreateAt(t *time.Time) *TeacherUpdateOne {
	if t != nil {
		tuo.SetCreateAt(*t)
	}
	return tuo
}

// SetUpdateAt sets the "update_at" field.
func (tuo *TeacherUpdateOne) SetUpdateAt(t time.Time) *TeacherUpdateOne {
	tuo.mutation.SetUpdateAt(t)
	return tuo
}

// SetUserID sets the "user_id" field.
func (tuo *TeacherUpdateOne) SetUserID(i int) *TeacherUpdateOne {
	tuo.mutation.SetUserID(i)
	return tuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (tuo *TeacherUpdateOne) SetNillableUserID(i *int) *TeacherUpdateOne {
	if i != nil {
		tuo.SetUserID(*i)
	}
	return tuo
}

// ClearUserID clears the value of the "user_id" field.
func (tuo *TeacherUpdateOne) ClearUserID() *TeacherUpdateOne {
	tuo.mutation.ClearUserID()
	return tuo
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (tuo *TeacherUpdateOne) AddCourseIDs(ids ...int) *TeacherUpdateOne {
	tuo.mutation.AddCourseIDs(ids...)
	return tuo
}

// AddCourses adds the "courses" edges to the Course entity.
func (tuo *TeacherUpdateOne) AddCourses(c ...*Course) *TeacherUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.AddCourseIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (tuo *TeacherUpdateOne) SetUser(u *User) *TeacherUpdateOne {
	return tuo.SetUserID(u.ID)
}

// Mutation returns the TeacherMutation object of the builder.
func (tuo *TeacherUpdateOne) Mutation() *TeacherMutation {
	return tuo.mutation
}

// ClearCourses clears all "courses" edges to the Course entity.
func (tuo *TeacherUpdateOne) ClearCourses() *TeacherUpdateOne {
	tuo.mutation.ClearCourses()
	return tuo
}

// RemoveCourseIDs removes the "courses" edge to Course entities by IDs.
func (tuo *TeacherUpdateOne) RemoveCourseIDs(ids ...int) *TeacherUpdateOne {
	tuo.mutation.RemoveCourseIDs(ids...)
	return tuo
}

// RemoveCourses removes "courses" edges to Course entities.
func (tuo *TeacherUpdateOne) RemoveCourses(c ...*Course) *TeacherUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return tuo.RemoveCourseIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (tuo *TeacherUpdateOne) ClearUser() *TeacherUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// Where appends a list predicates to the TeacherUpdate builder.
func (tuo *TeacherUpdateOne) Where(ps ...predicate.Teacher) *TeacherUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeacherUpdateOne) Select(field string, fields ...string) *TeacherUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Teacher entity.
func (tuo *TeacherUpdateOne) Save(ctx context.Context) (*Teacher, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeacherUpdateOne) SaveX(ctx context.Context) *Teacher {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeacherUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeacherUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TeacherUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateAt(); !ok {
		v := teacher.UpdateDefaultUpdateAt()
		tuo.mutation.SetUpdateAt(v)
	}
}

func (tuo *TeacherUpdateOne) sqlSave(ctx context.Context) (_node *Teacher, err error) {
	_spec := sqlgraph.NewUpdateSpec(teacher.Table, teacher.Columns, sqlgraph.NewFieldSpec(teacher.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Teacher.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, teacher.FieldID)
		for _, f := range fields {
			if !teacher.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != teacher.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Detail(); ok {
		_spec.SetField(teacher.FieldDetail, field.TypeString, value)
	}
	if tuo.mutation.DetailCleared() {
		_spec.ClearField(teacher.FieldDetail, field.TypeString)
	}
	if value, ok := tuo.mutation.CurriculumVitae(); ok {
		_spec.SetField(teacher.FieldCurriculumVitae, field.TypeString, value)
	}
	if tuo.mutation.CurriculumVitaeCleared() {
		_spec.ClearField(teacher.FieldCurriculumVitae, field.TypeString)
	}
	if value, ok := tuo.mutation.Works(); ok {
		_spec.SetField(teacher.FieldWorks, field.TypeString, value)
	}
	if tuo.mutation.WorksCleared() {
		_spec.ClearField(teacher.FieldWorks, field.TypeString)
	}
	if value, ok := tuo.mutation.Skills(); ok {
		_spec.SetField(teacher.FieldSkills, field.TypeString, value)
	}
	if tuo.mutation.SkillsCleared() {
		_spec.ClearField(teacher.FieldSkills, field.TypeString)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(teacher.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Level(); ok {
		_spec.SetField(teacher.FieldLevel, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedLevel(); ok {
		_spec.AddField(teacher.FieldLevel, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Avator(); ok {
		_spec.SetField(teacher.FieldAvator, field.TypeString, value)
	}
	if tuo.mutation.AvatorCleared() {
		_spec.ClearField(teacher.FieldAvator, field.TypeString)
	}
	if value, ok := tuo.mutation.CreateAt(); ok {
		_spec.SetField(teacher.FieldCreateAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdateAt(); ok {
		_spec.SetField(teacher.FieldUpdateAt, field.TypeTime, value)
	}
	if tuo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.CoursesTable,
			Columns: []string{teacher.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !tuo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.CoursesTable,
			Columns: []string{teacher.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   teacher.CoursesTable,
			Columns: []string{teacher.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   teacher.UserTable,
			Columns: []string{teacher.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   teacher.UserTable,
			Columns: []string{teacher.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Teacher{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{teacher.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
