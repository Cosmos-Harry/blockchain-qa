package handlers

import (
	"context"

	"github.com/Cosmos-Harry/blockchain-qa/indexer/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// PollHandler handles poll-related HTTP requests
type PollHandler struct {
	db    *database.DB
	redis *redis.Client
}

// NewPollHandler creates a new poll handler
func NewPollHandler(db *database.DB, redis *redis.Client) *PollHandler {
	return &PollHandler{
		db:    db,
		redis: redis,
	}
}

// GetPoll retrieves a poll by contract address
// GET /api/polls/:address
func (h *PollHandler) GetPoll(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "poll address is required",
		})
	}

	ctx := context.Background()

	// Try to get from cache first
	cacheKey := "poll:" + address
	var poll *database.Poll

	// Check Redis cache
	if h.redis != nil {
		cached, err := h.redis.Get(ctx, cacheKey).Result()
		if err == nil && cached != "" {
			// TODO: Unmarshal from JSON
			_ = cached
		}
	}

	// If not in cache, query database
	var err error
	poll, err = h.db.GetPollByAddress(ctx, address)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve poll",
		})
	}

	if poll == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "poll not found",
		})
	}

	// Cache for 1 minute
	if h.redis != nil {
		// TODO: Marshal to JSON and cache
	}

	return c.JSON(poll)
}

// ListPolls retrieves all polls with optional filtering
// GET /api/polls?state=active&limit=10&offset=0
func (h *PollHandler) ListPolls(c *fiber.Ctx) error {
	state := c.Query("state", "")
	limit := c.QueryInt("limit", 20)
	offset := c.QueryInt("offset", 0)

	if limit > 100 {
		limit = 100
	}

	ctx := context.Background()
	polls, err := h.db.ListPolls(ctx, state, limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to list polls",
		})
	}

	return c.JSON(fiber.Map{
		"polls":  polls,
		"limit":  limit,
		"offset": offset,
		"count":  len(polls),
	})
}

// GetPollVotes retrieves all votes for a poll
// GET /api/polls/:address/votes?revealed_only=false
func (h *PollHandler) GetPollVotes(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "poll address is required",
		})
	}

	revealedOnly := c.QueryBool("revealed_only", false)

	ctx := context.Background()
	votes, err := h.db.ListVotesByPoll(ctx, address, revealedOnly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve votes",
		})
	}

	// If not revealing, hide sensitive data
	if !revealedOnly {
		for _, vote := range votes {
			if !vote.Revealed {
				vote.Choice = nil
				vote.Nonce = nil
			}
		}
	}

	return c.JSON(fiber.Map{
		"votes": votes,
		"count": len(votes),
	})
}

// GetPollResults retrieves the tallied results for a poll
// GET /api/polls/:address/results
func (h *PollHandler) GetPollResults(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "poll address is required",
		})
	}

	ctx := context.Background()

	// Get the poll to check state
	poll, err := h.db.GetPollByAddress(ctx, address)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve poll",
		})
	}

	if poll == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "poll not found",
		})
	}

	// Check if results exist
	result, err := h.db.GetResultByPoll(ctx, address)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve results",
		})
	}

	if result == nil {
		// If poll is not tallied yet, return vote counts so far
		if poll.State != "tallied" {
			voteCount, err := h.db.GetVoteCount(ctx, address, false)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "failed to get vote count",
				})
			}

			return c.JSON(fiber.Map{
				"status":      "pending",
				"poll_state":  poll.State,
				"total_votes": voteCount,
				"message":     "poll results not yet tallied",
			})
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "results not found",
		})
	}

	return c.JSON(result)
}

// GetVoteCount retrieves vote statistics for a poll
// GET /api/polls/:address/stats
func (h *PollHandler) GetVoteCount(c *fiber.Ctx) error {
	address := c.Params("address")
	if address == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "poll address is required",
		})
	}

	ctx := context.Background()

	totalVotes, err := h.db.GetVoteCount(ctx, address, false)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get vote count",
		})
	}

	revealedVotes, err := h.db.GetVoteCount(ctx, address, true)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to get revealed vote count",
		})
	}

	return c.JSON(fiber.Map{
		"poll_address":   address,
		"total_votes":    totalVotes,
		"revealed_votes": revealedVotes,
		"pending_reveals": totalVotes - revealedVotes,
	})
}
