package payloads

type UpdateMemberRoles struct {
	GuildID string
	UserID  string
	RoleID  string
}
