package hmac

import (
	auth "github.com/rinojoss/test-repo"
)

type HMAC struct {
	userRepo auth.UserRepository
}

func (h *HMAC) Authenticate() error {
	_ = h.userRepo.FindByID("user1")

	return nil
}
