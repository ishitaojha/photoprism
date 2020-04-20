package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism/internal/event"
	"github.com/photoprism/photoprism/internal/form"
	"github.com/photoprism/photoprism/internal/query"
)

type EntityEvent string

const (
	EntityUpdated EntityEvent = "updated"
	EntityCreated EntityEvent = "created"
	EntityDeleted EntityEvent = "deleted"
)

func PublishPhotoEvent(e EntityEvent, uuid string, c *gin.Context, q *query.Query) {
	f := form.PhotoSearch{ID: uuid}
	result, _, err := q.Photos(f)

	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrUnexpectedError)
		return
	}

	event.PublishEntities("photos", string(e), result)
}

func PublishAlbumEvent(e EntityEvent, uuid string, c *gin.Context, q *query.Query) {
	f := form.AlbumSearch{ID: uuid}
	result, err := q.Albums(f)

	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrUnexpectedError)
		return
	}

	event.PublishEntities("albums", string(e), result)
}

func PublishLabelEvent(e EntityEvent, uuid string, c *gin.Context, q *query.Query) {
	f := form.LabelSearch{ID: uuid}
	result, err := q.Labels(f)

	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrUnexpectedError)
		return
	}

	event.PublishEntities("labels", string(e), result)
}
