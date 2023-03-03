// Code generated by "stringer -linecomment -type=TaskStatus -trimprefix=Task -output=task_string.go"; DO NOT EDIT.

package client

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TaskStatusUnspecified-0]
	_ = x[TaskPending-1]
	_ = x[TaskStarted-2]
	_ = x[TaskSuccess-3]
	_ = x[TaskFailure-4]
	_ = x[TaskRetry-5]
	_ = x[TaskRevoked-6]
}

const _TaskStatus_name = "StatusUnspecifiedPendingStartedSuccessFailureRetryRevoked"

var _TaskStatus_index = [...]uint8{0, 17, 24, 31, 38, 45, 50, 57}

func (i TaskStatus) String() string {
	if i < 0 || i >= TaskStatus(len(_TaskStatus_index)-1) {
		return "TaskStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TaskStatus_name[_TaskStatus_index[i]:_TaskStatus_index[i+1]]
}
