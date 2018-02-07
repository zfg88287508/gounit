package gounit

import (
	"go/ast"
	"reflect"
	"testing"
)

func TestNewVisitor(t *testing.T) {
	type args struct {
		match matchFunc
	}

	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 *Visitor
	}{
		{
			name:  "success",
			args:  func(*testing.T) args { return args{} },
			want1: &Visitor{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)
			got1 := NewVisitor(tArgs.match)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("NewVisitor got1 = %v, want1: %v", got1, tt.want1)
			}

		})
	}
}

func TestVisitor_Visit(t *testing.T) {
	type args struct {
		node ast.Node
	}

	var notFoundVisitor = NewVisitor(func(*ast.FuncDecl) bool {
		return false
	})

	var foundVisitor = NewVisitor(func(*ast.FuncDecl) bool {
		return true
	})

	tests := []struct {
		name    string
		args    func(t *testing.T) args
		init    func(t *testing.T) *Visitor
		inspect func(r *Visitor, t *testing.T) //inspects receiver after method run

		want1 ast.Visitor
	}{
		{
			name:  "func not found",
			init:  func(*testing.T) *Visitor { return notFoundVisitor },
			args:  func(*testing.T) args { return args{} },
			want1: notFoundVisitor,
		},
		{
			name: "func found",
			init: func(*testing.T) *Visitor {
				return foundVisitor
			},
			args: func(*testing.T) args {
				return args{
					node: &ast.FuncDecl{},
				}
			},
			inspect: func(v *Visitor, t *testing.T) {
				if len(v.found) == 0 {
					t.Errorf("expected non-empty v.found")
				}
			},
			want1: foundVisitor,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)
			receiver := tt.init(t)
			got1 := receiver.Visit(tArgs.node)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Visitor.Visit got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

func TestVisitor_Funcs(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) *Visitor
		inspect func(r *Visitor, t *testing.T) //inspects receiver after test run

		want1 []*Func
	}{
		{
			name: "success",
			init: func(t *testing.T) *Visitor {
				return &Visitor{found: []*Func{{}}}
			},
			want1: []*Func{{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)
			got1 := receiver.Funcs()

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Visitor.Funcs got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}
