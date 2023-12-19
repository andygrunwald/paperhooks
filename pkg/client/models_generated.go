// Code generated by "generate_models --output models_generated.go"; DO NOT EDIT.
package client

import "encoding/json"
import "time"

type Correspondent struct {
	ID                 int64             `json:"id"`
	Slug               string            `json:"slug"`
	Name               string            `json:"name"`
	Match              string            `json:"match"`
	MatchingAlgorithm  MatchingAlgorithm `json:"matching_algorithm"`
	IsInsensitive      bool              `json:"is_insensitive"`
	DocumentCount      int64             `json:"document_count"`
	LastCorrespondence *time.Time        `json:"last_correspondence"`

	// Object owner; objects without owner can be viewed and edited by all users.
	Owner *int64 `json:"owner"`
}

type CorrespondentFields struct {
	objectFields
}

var _ json.Marshaler = (*CorrespondentFields)(nil)

func NewCorrespondentFields() *CorrespondentFields {
	return &CorrespondentFields{objectFields{}}
}

// SetName sets the "name" field.
func (f *CorrespondentFields) SetName(name string) *CorrespondentFields {
	f.set("name", name)
	return f
}

// SetMatch sets the "match" field.
func (f *CorrespondentFields) SetMatch(match string) *CorrespondentFields {
	f.set("match", match)
	return f
}

// SetMatchingAlgorithm sets the "matching_algorithm" field.
func (f *CorrespondentFields) SetMatchingAlgorithm(matchingAlgorithm MatchingAlgorithm) *CorrespondentFields {
	f.set("matching_algorithm", matchingAlgorithm)
	return f
}

// SetIsInsensitive sets the "is_insensitive" field.
func (f *CorrespondentFields) SetIsInsensitive(isInsensitive bool) *CorrespondentFields {
	f.set("is_insensitive", isInsensitive)
	return f
}

// SetOwner sets the "owner" field.
//
// Object owner; objects without owner can be viewed and edited by all users.
func (f *CorrespondentFields) SetOwner(owner *int64) *CorrespondentFields {
	f.set("owner", owner)
	return f
}

type CustomField struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	DataType string `json:"data_type"`

	// Object owner; objects without owner can be viewed and edited by all users.
	Owner *int64 `json:"owner"`
}

type CustomFieldFields struct {
	objectFields
}

var _ json.Marshaler = (*CustomFieldFields)(nil)

func NewCustomFieldFields() *CustomFieldFields {
	return &CustomFieldFields{objectFields{}}
}

// SetName sets the "name" field.
func (f *CustomFieldFields) SetName(name string) *CustomFieldFields {
	f.set("name", name)
	return f
}

// SetDataType sets the "data_type" field.
func (f *CustomFieldFields) SetDataType(dataType string) *CustomFieldFields {
	f.set("data_type", dataType)
	return f
}

// SetOwner sets the "owner" field.
//
// Object owner; objects without owner can be viewed and edited by all users.
func (f *CustomFieldFields) SetOwner(owner *int64) *CustomFieldFields {
	f.set("owner", owner)
	return f
}

type Document struct {
	// ID of the document.
	ID int64 `json:"id"`

	// Title of the document.
	Title string `json:"title"`

	// Plain-text content of the document.
	Content string `json:"content"`

	// List of tag IDs assigned to this document, or empty list.
	Tags []int64 `json:"tags"`

	// Document type of this document or nil.
	DocumentType *int64 `json:"document_type"`

	// Correspondent of this document or nil.
	Correspondent *int64 `json:"correspondent"`

	// Storage path of this document or nil.
	StoragePath *int64 `json:"storage_path"`

	// The date time at which this document was created.
	Created time.Time `json:"created"`

	// The date at which this document was last edited in paperless.
	Modified time.Time `json:"modified"`

	// The date at which this document was added to paperless.
	Added time.Time `json:"added"`

	// The identifier of this document in a physical document archive.
	ArchiveSerialNumber *int64 `json:"archive_serial_number"`

	// Verbose filename of the original document.
	OriginalFileName string `json:"original_file_name"`

	// Verbose filename of the archived document. Nil if no archived document is available.
	ArchivedFileName *string `json:"archived_file_name"`

	// Custom fields on the document.
	CustomFields []CustomFieldInstance `json:"custom_fields"`

	// Object owner; objects without owner can be viewed and edited by all users.
	Owner *int64 `json:"owner"`
}

type DocumentFields struct {
	objectFields
}

var _ json.Marshaler = (*DocumentFields)(nil)

func NewDocumentFields() *DocumentFields {
	return &DocumentFields{objectFields{}}
}

// SetTitle sets the "title" field.
//
// Title of the document.
func (f *DocumentFields) SetTitle(title string) *DocumentFields {
	f.set("title", title)
	return f
}

// SetContent sets the "content" field.
//
// Plain-text content of the document.
func (f *DocumentFields) SetContent(content string) *DocumentFields {
	f.set("content", content)
	return f
}

// SetTags sets the "tags" field.
//
// List of tag IDs assigned to this document, or empty list.
func (f *DocumentFields) SetTags(tags []int64) *DocumentFields {
	f.set("tags", tags)
	return f
}

// SetDocumentType sets the "document_type" field.
//
// Document type of this document or nil.
func (f *DocumentFields) SetDocumentType(documentType *int64) *DocumentFields {
	f.set("document_type", documentType)
	return f
}

// SetCorrespondent sets the "correspondent" field.
//
// Correspondent of this document or nil.
func (f *DocumentFields) SetCorrespondent(correspondent *int64) *DocumentFields {
	f.set("correspondent", correspondent)
	return f
}

// SetStoragePath sets the "storage_path" field.
//
// Storage path of this document or nil.
func (f *DocumentFields) SetStoragePath(storagePath *int64) *DocumentFields {
	f.set("storage_path", storagePath)
	return f
}

// SetCreated sets the "created" field.
//
// The date time at which this document was created.
func (f *DocumentFields) SetCreated(created time.Time) *DocumentFields {
	f.set("created", created)
	return f
}

// SetArchiveSerialNumber sets the "archive_serial_number" field.
//
// The identifier of this document in a physical document archive.
func (f *DocumentFields) SetArchiveSerialNumber(archiveSerialNumber *int64) *DocumentFields {
	f.set("archive_serial_number", archiveSerialNumber)
	return f
}

// SetCustomFields sets the "custom_fields" field.
//
// Custom fields on the document.
func (f *DocumentFields) SetCustomFields(customFields []CustomFieldInstance) *DocumentFields {
	f.set("custom_fields", customFields)
	return f
}

// SetOwner sets the "owner" field.
//
// Object owner; objects without owner can be viewed and edited by all users.
func (f *DocumentFields) SetOwner(owner *int64) *DocumentFields {
	f.set("owner", owner)
	return f
}

type DocumentType struct {
	ID                int64             `json:"id"`
	Slug              string            `json:"slug"`
	Name              string            `json:"name"`
	Match             string            `json:"match"`
	MatchingAlgorithm MatchingAlgorithm `json:"matching_algorithm"`
	IsInsensitive     bool              `json:"is_insensitive"`
	DocumentCount     int64             `json:"document_count"`

	// Object owner; objects without owner can be viewed and edited by all users.
	Owner *int64 `json:"owner"`
}

type DocumentTypeFields struct {
	objectFields
}

var _ json.Marshaler = (*DocumentTypeFields)(nil)

func NewDocumentTypeFields() *DocumentTypeFields {
	return &DocumentTypeFields{objectFields{}}
}

// SetName sets the "name" field.
func (f *DocumentTypeFields) SetName(name string) *DocumentTypeFields {
	f.set("name", name)
	return f
}

// SetMatch sets the "match" field.
func (f *DocumentTypeFields) SetMatch(match string) *DocumentTypeFields {
	f.set("match", match)
	return f
}

// SetMatchingAlgorithm sets the "matching_algorithm" field.
func (f *DocumentTypeFields) SetMatchingAlgorithm(matchingAlgorithm MatchingAlgorithm) *DocumentTypeFields {
	f.set("matching_algorithm", matchingAlgorithm)
	return f
}

// SetIsInsensitive sets the "is_insensitive" field.
func (f *DocumentTypeFields) SetIsInsensitive(isInsensitive bool) *DocumentTypeFields {
	f.set("is_insensitive", isInsensitive)
	return f
}

// SetOwner sets the "owner" field.
//
// Object owner; objects without owner can be viewed and edited by all users.
func (f *DocumentTypeFields) SetOwner(owner *int64) *DocumentTypeFields {
	f.set("owner", owner)
	return f
}

type Group struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type GroupFields struct {
	objectFields
}

var _ json.Marshaler = (*GroupFields)(nil)

func NewGroupFields() *GroupFields {
	return &GroupFields{objectFields{}}
}

// SetName sets the "name" field.
func (f *GroupFields) SetName(name string) *GroupFields {
	f.set("name", name)
	return f
}

type StoragePath struct {
	ID                int64             `json:"id"`
	Slug              string            `json:"slug"`
	Name              string            `json:"name"`
	Match             string            `json:"match"`
	MatchingAlgorithm MatchingAlgorithm `json:"matching_algorithm"`
	IsInsensitive     bool              `json:"is_insensitive"`
	DocumentCount     int64             `json:"document_count"`

	// Object owner; objects without owner can be viewed and edited by all users.
	Owner *int64 `json:"owner"`
}

type StoragePathFields struct {
	objectFields
}

var _ json.Marshaler = (*StoragePathFields)(nil)

func NewStoragePathFields() *StoragePathFields {
	return &StoragePathFields{objectFields{}}
}

// SetName sets the "name" field.
func (f *StoragePathFields) SetName(name string) *StoragePathFields {
	f.set("name", name)
	return f
}

// SetMatch sets the "match" field.
func (f *StoragePathFields) SetMatch(match string) *StoragePathFields {
	f.set("match", match)
	return f
}

// SetMatchingAlgorithm sets the "matching_algorithm" field.
func (f *StoragePathFields) SetMatchingAlgorithm(matchingAlgorithm MatchingAlgorithm) *StoragePathFields {
	f.set("matching_algorithm", matchingAlgorithm)
	return f
}

// SetIsInsensitive sets the "is_insensitive" field.
func (f *StoragePathFields) SetIsInsensitive(isInsensitive bool) *StoragePathFields {
	f.set("is_insensitive", isInsensitive)
	return f
}

// SetOwner sets the "owner" field.
//
// Object owner; objects without owner can be viewed and edited by all users.
func (f *StoragePathFields) SetOwner(owner *int64) *StoragePathFields {
	f.set("owner", owner)
	return f
}

type Tag struct {
	ID                int64             `json:"id"`
	Slug              string            `json:"slug"`
	Name              string            `json:"name"`
	Color             Color             `json:"color"`
	TextColor         Color             `json:"text_color"`
	Match             string            `json:"match"`
	MatchingAlgorithm MatchingAlgorithm `json:"matching_algorithm"`
	IsInsensitive     bool              `json:"is_insensitive"`
	IsInboxTag        bool              `json:"is_inbox_tag"`
	DocumentCount     int64             `json:"document_count"`

	// Object owner; objects without owner can be viewed and edited by all users.
	Owner *int64 `json:"owner"`
}

type TagFields struct {
	objectFields
}

var _ json.Marshaler = (*TagFields)(nil)

func NewTagFields() *TagFields {
	return &TagFields{objectFields{}}
}

// SetName sets the "name" field.
func (f *TagFields) SetName(name string) *TagFields {
	f.set("name", name)
	return f
}

// SetColor sets the "color" field.
func (f *TagFields) SetColor(color Color) *TagFields {
	f.set("color", color)
	return f
}

// SetTextColor sets the "text_color" field.
func (f *TagFields) SetTextColor(textColor Color) *TagFields {
	f.set("text_color", textColor)
	return f
}

// SetMatch sets the "match" field.
func (f *TagFields) SetMatch(match string) *TagFields {
	f.set("match", match)
	return f
}

// SetMatchingAlgorithm sets the "matching_algorithm" field.
func (f *TagFields) SetMatchingAlgorithm(matchingAlgorithm MatchingAlgorithm) *TagFields {
	f.set("matching_algorithm", matchingAlgorithm)
	return f
}

// SetIsInsensitive sets the "is_insensitive" field.
func (f *TagFields) SetIsInsensitive(isInsensitive bool) *TagFields {
	f.set("is_insensitive", isInsensitive)
	return f
}

// SetIsInboxTag sets the "is_inbox_tag" field.
func (f *TagFields) SetIsInboxTag(isInboxTag bool) *TagFields {
	f.set("is_inbox_tag", isInboxTag)
	return f
}

// SetOwner sets the "owner" field.
//
// Object owner; objects without owner can be viewed and edited by all users.
func (f *TagFields) SetOwner(owner *int64) *TagFields {
	f.set("owner", owner)
	return f
}

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	IsActive    bool   `json:"is_active"`
	IsStaff     bool   `json:"is_staff"`
	IsSuperuser bool   `json:"is_superuser"`
}

type UserFields struct {
	objectFields
}

var _ json.Marshaler = (*UserFields)(nil)

func NewUserFields() *UserFields {
	return &UserFields{objectFields{}}
}

// SetUsername sets the "username" field.
func (f *UserFields) SetUsername(username string) *UserFields {
	f.set("username", username)
	return f
}

// SetEmail sets the "email" field.
func (f *UserFields) SetEmail(email string) *UserFields {
	f.set("email", email)
	return f
}

// SetFirstName sets the "first_name" field.
func (f *UserFields) SetFirstName(firstName string) *UserFields {
	f.set("first_name", firstName)
	return f
}

// SetLastName sets the "last_name" field.
func (f *UserFields) SetLastName(lastName string) *UserFields {
	f.set("last_name", lastName)
	return f
}

// SetIsActive sets the "is_active" field.
func (f *UserFields) SetIsActive(isActive bool) *UserFields {
	f.set("is_active", isActive)
	return f
}

// SetIsStaff sets the "is_staff" field.
func (f *UserFields) SetIsStaff(isStaff bool) *UserFields {
	f.set("is_staff", isStaff)
	return f
}

// SetIsSuperuser sets the "is_superuser" field.
func (f *UserFields) SetIsSuperuser(isSuperuser bool) *UserFields {
	f.set("is_superuser", isSuperuser)
	return f
}
