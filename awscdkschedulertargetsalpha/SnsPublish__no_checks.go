//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SnsPublish) validateAddTargetActionToRoleParameters(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (s *jsiiProxy_SnsPublish) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewSnsPublishParameters(topic awssns.ITopic, props *ScheduleTargetBaseProps) error {
	return nil
}
