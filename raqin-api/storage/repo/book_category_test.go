// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repo

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testBookCategories(t *testing.T) {
	t.Parallel()

	query := BookCategories()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBookCategoriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookCategoriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BookCategories().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookCategoriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookCategorySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBookCategoriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BookCategoryExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if BookCategory exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BookCategoryExists to return true, but got false.")
	}
}

func testBookCategoriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bookCategoryFound, err := FindBookCategory(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if bookCategoryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBookCategoriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BookCategories().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBookCategoriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BookCategories().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBookCategoriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bookCategoryOne := &BookCategory{}
	bookCategoryTwo := &BookCategory{}
	if err = randomize.Struct(seed, bookCategoryOne, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, bookCategoryTwo, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBookCategoriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bookCategoryOne := &BookCategory{}
	bookCategoryTwo := &BookCategory{}
	if err = randomize.Struct(seed, bookCategoryOne, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}
	if err = randomize.Struct(seed, bookCategoryTwo, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bookCategoryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bookCategoryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bookCategoryBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func bookCategoryAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BookCategory) error {
	*o = BookCategory{}
	return nil
}

func testBookCategoriesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BookCategory{}
	o := &BookCategory{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookCategory object: %s", err)
	}

	AddBookCategoryHook(boil.BeforeInsertHook, bookCategoryBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bookCategoryBeforeInsertHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterInsertHook, bookCategoryAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterInsertHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterSelectHook, bookCategoryAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterSelectHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.BeforeUpdateHook, bookCategoryBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bookCategoryBeforeUpdateHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterUpdateHook, bookCategoryAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterUpdateHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.BeforeDeleteHook, bookCategoryBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bookCategoryBeforeDeleteHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterDeleteHook, bookCategoryAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterDeleteHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.BeforeUpsertHook, bookCategoryBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bookCategoryBeforeUpsertHooks = []BookCategoryHook{}

	AddBookCategoryHook(boil.AfterUpsertHook, bookCategoryAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bookCategoryAfterUpsertHooks = []BookCategoryHook{}
}

func testBookCategoriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookCategoriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bookCategoryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBookCategoryToOneBookUsingBook(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BookCategory
	var foreign Book

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, bookDBTypes, false, bookColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Book struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.BookID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Book().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BookCategorySlice{&local}
	if err = local.L.LoadBook(ctx, tx, false, (*[]*BookCategory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Book == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Book = nil
	if err = local.L.LoadBook(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Book == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBookCategoryToOneCategoryUsingCategory(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local BookCategory
	var foreign Category

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, bookCategoryDBTypes, false, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, categoryDBTypes, false, categoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Category struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CategoryID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Category().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := BookCategorySlice{&local}
	if err = local.L.LoadCategory(ctx, tx, false, (*[]*BookCategory)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Category == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Category = nil
	if err = local.L.LoadCategory(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Category == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testBookCategoryToOneSetOpBookUsingBook(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c Book

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, bookDBTypes, false, strmangle.SetComplement(bookPrimaryKeyColumns, bookColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Book{&b, &c} {
		err = a.SetBook(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Book != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BookCategories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.BookID != x.ID {
			t.Error("foreign key was wrong value", a.BookID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.BookID))
		reflect.Indirect(reflect.ValueOf(&a.BookID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.BookID != x.ID {
			t.Error("foreign key was wrong value", a.BookID, x.ID)
		}
	}
}
func testBookCategoryToOneSetOpCategoryUsingCategory(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a BookCategory
	var b, c Category

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, bookCategoryDBTypes, false, strmangle.SetComplement(bookCategoryPrimaryKeyColumns, bookCategoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, categoryDBTypes, false, strmangle.SetComplement(categoryPrimaryKeyColumns, categoryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Category{&b, &c} {
		err = a.SetCategory(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Category != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.BookCategories[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CategoryID != x.ID {
			t.Error("foreign key was wrong value", a.CategoryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CategoryID))
		reflect.Indirect(reflect.ValueOf(&a.CategoryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CategoryID != x.ID {
			t.Error("foreign key was wrong value", a.CategoryID, x.ID)
		}
	}
}

func testBookCategoriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookCategoriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BookCategorySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBookCategoriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BookCategories().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bookCategoryDBTypes = map[string]string{`ID`: `int`, `BookID`: `int`, `CategoryID`: `int`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`, `DeletedAt`: `timestamp`}
	_                   = bytes.MinRead
)

func testBookCategoriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bookCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bookCategoryAllColumns) == len(bookCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBookCategoriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bookCategoryAllColumns) == len(bookCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BookCategory{}
	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bookCategoryDBTypes, true, bookCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bookCategoryAllColumns, bookCategoryPrimaryKeyColumns) {
		fields = bookCategoryAllColumns
	} else {
		fields = strmangle.SetComplement(
			bookCategoryAllColumns,
			bookCategoryPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := BookCategorySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBookCategoriesUpsert(t *testing.T) {
	t.Parallel()

	if len(bookCategoryAllColumns) == len(bookCategoryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLBookCategoryUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BookCategory{}
	if err = randomize.Struct(seed, &o, bookCategoryDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookCategory: %s", err)
	}

	count, err := BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bookCategoryDBTypes, false, bookCategoryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BookCategory struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BookCategory: %s", err)
	}

	count, err = BookCategories().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
