package models

type Media struct {
	BaseModel

	// pdf, jpeg, etc... | TODO should be enum
	Type string

	Path string

	// s3, gc, minio, etc | TODO should be enum
	Store string

	Size uint
}

type CreateMediaData struct {
	Type  string
	Path  string
	Store string
	Size  uint
}

type MediaService interface {
	Create(d CreateMediaData)
}

type MediaRepository interface {
	//GetBy
}
