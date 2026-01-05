package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Message holds the schema definition for the Message entity.
type Message struct {
    ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
    return []ent.Field{
        field.UUID("id", uuid.UUID{}).
            Default(uuid.New),

        field.String("username").
            NotEmpty(),

        field.String("text").
            NotEmpty(),

        field.Time("created_at").
            Default(time.Now),
    }
}


// Edges of the Message.
func (Message) Edges() []ent.Edge {
    return nil
}
    