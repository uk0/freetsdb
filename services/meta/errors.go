package meta

import (
	"errors"
	"fmt"
)

var (
	// ErrStoreOpen is returned when opening an already open store.
	ErrStoreOpen = errors.New("store already open")

	// ErrStoreClosed is returned when closing an already closed store.
	ErrStoreClosed = errors.New("raft store already closed")

	// ErrTooManyPeers is returned when more than 3 peers are used.
	ErrTooManyPeers = errors.New("too many peers")
)

var (
	// ErrNodeExists is returned when creating an already existing node.
	ErrNodeExists = errors.New("node already exists")

	// ErrNodeNotFound is returned when mutating a node that doesn't exist.
	ErrNodeNotFound = errors.New("node not found")

	// ErrNodesRequired is returned when at least one node is required for an operation.
	// This occurs when creating a shard group.
	ErrNodesRequired = errors.New("at least one node required")

	// ErrNodeIDRequired is returned when using a zero node id.
	ErrNodeIDRequired = errors.New("node id must be greater than 0")

	// ErrNodeUnableToDropFinalNode is returned if the node being dropped is the last
	// node in the cluster
	ErrNodeUnableToDropFinalNode = errors.New("unable to drop the final node in a cluster")

	ErrNodeUnableToDropNode = errors.New("unable to drop the node in a cluster")
)

var (
	// ErrDatabaseExists is returned when creating an already existing database.
	ErrDatabaseExists = errors.New("database already exists")

	// ErrDatabaseNotExists is returned when operating on a not existing database.
	ErrDatabaseNotExists = errors.New("database does not exist")

	// ErrDatabaseNameRequired is returned when creating a database without a name.
	ErrDatabaseNameRequired = errors.New("database name required")
)

var (
	// ErrRetentionPolicyExists is returned when creating an already existing policy.
	ErrRetentionPolicyExists = errors.New("retention policy already exists")

	// ErrRetentionPolicyDefault is returned when attempting a prohibited operation
	// on a default retention policy.
	ErrRetentionPolicyDefault = errors.New("retention policy is default")

	// ErrRetentionPolicyNameRequired is returned when creating a policy without a name.
	ErrRetentionPolicyNameRequired = errors.New("retention policy name required")

	// ErrRetentionPolicyNameExists is returned when renaming a policy to
	// the same name as another existing policy.
	ErrRetentionPolicyNameExists = errors.New("retention policy name already exists")

	// ErrRetentionPolicyDurationTooLow is returned when updating a retention
	// policy that has a duration lower than the allowed minimum.
	ErrRetentionPolicyDurationTooLow = errors.New(fmt.Sprintf("retention policy duration must be at least %s",
		MinRetentionPolicyDuration))

	// ErrRetentionPolicyConflict is returned when creating a retention policy conflicts
	// with an existing policy.
	ErrRetentionPolicyConflict = errors.New("retention policy conflicts with an existing policy")

	// ErrReplicationFactorTooLow is returned when the replication factor is not in an
	// acceptable range.
	ErrReplicationFactorTooLow = errors.New("replication factor must be greater than 0")
)

var (
	// ErrShardGroupExists is returned when creating an already existing shard group.
	ErrShardGroupExists = errors.New("shard group already exists")

	// ErrShardGroupNotFound is returned when mutating a shard group that doesn't exist.
	ErrShardGroupNotFound = errors.New("shard group not found")

	// ErrShardNotReplicated is returned if the node requested to be dropped has
	// the last copy of a shard present and the force keyword was not used
	ErrShardNotReplicated = errors.New("shard not replicated")
)

var (
	// ErrContinuousQueryExists is returned when creating an already existing continuous query.
	ErrContinuousQueryExists = errors.New("continuous query already exists")

	// ErrContinuousQueryNotFound is returned when removing a continuous query that doesn't exist.
	ErrContinuousQueryNotFound = errors.New("continuous query not found")
)

var (
	// ErrSubscriptionExists is returned when creating an already existing subscription.
	ErrSubscriptionExists = errors.New("subscription already exists")

	// ErrSubscriptionNotFound is returned when removing a subscription that doesn't exist.
	ErrSubscriptionNotFound = errors.New("subscription not found")
)

var (
	// ErrUserExists is returned when creating an already existing user.
	ErrUserExists = errors.New("user already exists")

	// ErrUserNotFound is returned when mutating a user that doesn't exist.
	ErrUserNotFound = errors.New("user not found")

	// ErrUsernameRequired is returned when creating a user without a username.
	ErrUsernameRequired = errors.New("username required")

	// ErrAuthenticate is returned when authentication fails.
	ErrAuthenticate = errors.New("authentication failed")
)
