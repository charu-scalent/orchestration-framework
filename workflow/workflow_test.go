package workflow

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func TestWorkflow(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	idempotentMockInst := NewMockIdempotent(ctrl)

	type args struct {
		idempotentKey  string
		idempotentInst Idempotent
	}
	tests := []struct {
		name string
		args args
		want *Workflow
	}{
		{
			name: "CreateWorkflow",
			args: args{
				idempotentKey:  "idem_key",
				idempotentInst: idempotentMockInst,
			},
			want: &Workflow{
				Steps:          []Step{},
				IdempotentKey:  "idem_key",
				IdempotentInst: idempotentMockInst,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateWorkflow(tt.args.idempotentKey, tt.args.idempotentInst)
			if !cmp.Equal(got.IdempotentKey, tt.want.IdempotentKey) {
				t.Errorf("CreateWorkflow() = %v, want %v", got.IdempotentKey, tt.want.IdempotentKey)
			}
		})
	}
}

func TestWorkflow_Register(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	idempotentMockInst := NewMockIdempotent(ctrl)

	orderMockInstance := NewMockOrderInterface(ctrl)

	type fields struct {
		Steps          []Step
		IdempotentKey  string
		IdempotentInst Idempotent
	}
	type args struct {
		instance        interface{}
		method          string
		rollbackMethod  string
		isMandatoryStep bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "Register Workflow Test",
			fields: fields{
				Steps:          []Step{},
				IdempotentKey:  "idem_key",
				IdempotentInst: idempotentMockInst,
			},
			args:   args{
				instance:        orderMockInstance,
				method:          "CreateOrder",
				rollbackMethod:  "RollbackCreateOrder",
				isMandatoryStep: false,
			},
		},
	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Workflow{
				Steps:          tt.fields.Steps,
				IdempotentKey:  tt.fields.IdempotentKey,
				IdempotentInst: tt.fields.IdempotentInst,
			}
			w.Register(tt.args.instance, tt.args.method, tt.args.rollbackMethod, tt.args.isMandatoryStep)
		})
	}
}

func TestWorkflow_Execute(t *testing.T) {
	type fields struct {
		Steps          []Step
		IdempotentKey  string
		IdempotentInst Idempotent
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Workflow{
				Steps:          tt.fields.Steps,
				IdempotentKey:  tt.fields.IdempotentKey,
				IdempotentInst: tt.fields.IdempotentInst,
			}
			if err := w.Execute(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Workflow.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWorkflow_executeStep(t *testing.T) {
	type fields struct {
		Steps          []Step
		IdempotentKey  string
		IdempotentInst Idempotent
	}
	type args struct {
		ctx           context.Context
		idempotentKey string
		step          Step
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Workflow{
				Steps:          tt.fields.Steps,
				IdempotentKey:  tt.fields.IdempotentKey,
				IdempotentInst: tt.fields.IdempotentInst,
			}
			if err := w.executeStep(tt.args.ctx, tt.args.idempotentKey, tt.args.step); (err != nil) != tt.wantErr {
				t.Errorf("Workflow.executeStep() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
