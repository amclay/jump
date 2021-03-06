package permissions

import "github.com/gdbu/dbl"

func newResource(key string) (r Resource) {
	r.Key = key
	r.Groups = make(Groups)
	return
}

// Resource represents a permissions resource entry
type Resource struct {
	dbl.Entry

	Key    string `json:"key"`
	Groups `json:"groups"`
}

// core.Value interface methods below

// GetID will get the message ID
func (r *Resource) GetID() (id string) { return r.ID }

// GetCreatedAt will get the created at timestamp
func (r *Resource) GetCreatedAt() (createdAt int64) { return r.CreatedAt }

// GetUpdatedAt will get the updated at timestamp
func (r *Resource) GetUpdatedAt() (updatedAt int64) { return r.UpdatedAt }

// GetRelationships will get the associated relationship IDs
func (r *Resource) GetRelationships() (rs dbl.Relationships) {
	rs.Append(r.Key)
	return
}

// SetID will get the message ID
func (r *Resource) SetID(id string) { r.ID = id }

// SetCreatedAt will get the created at timestamp
func (r *Resource) SetCreatedAt(createdAt int64) { r.CreatedAt = createdAt }

// SetUpdatedAt will get the updated at timestamp
func (r *Resource) SetUpdatedAt(updatedAt int64) { r.UpdatedAt = updatedAt }
