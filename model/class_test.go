//  Copyright (c) 2021 PingLeMe Team. All rights reserved.

package model

import (
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestClass_GetAllStudents(t *testing.T) {
	type fields struct {
		Model    gorm.Model
		Name     string
		Teachers []User
		Students []User
	}
	tests := []struct {
		name    string
		fields  fields
		want    []User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			class := &Class{
				Model:    tt.fields.Model,
				Name:     tt.fields.Name,
				Teachers: tt.fields.Teachers,
				Students: tt.fields.Students,
			}
			got, err := class.GetAllStudents()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllStudents() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllStudents() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClass_GetAllTeachers(t *testing.T) {
	type fields struct {
		Model    gorm.Model
		Name     string
		Teachers []User
		Students []User
	}
	tests := []struct {
		name    string
		fields  fields
		want    []User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			class := &Class{
				Model:    tt.fields.Model,
				Name:     tt.fields.Name,
				Teachers: tt.fields.Teachers,
				Students: tt.fields.Students,
			}
			got, err := class.GetAllTeachers()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllTeachers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllTeachers() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetClassByID(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		ID interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Class
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Repo := &Repository{
				DB: tt.fields.DB,
			}
			got, err := Repo.GetClassByID(tt.args.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClassByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClassByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
